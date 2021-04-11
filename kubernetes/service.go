package kubernetes

import (
	"context"
	"fmt"
	"log"

	instancev1 "github.com/cownetwork/instance-controller/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/dynamic"
)

var gvr = instancev1.GroupVersion.WithResource("instances")

type Service interface {
	CreateInstance(ctx context.Context, instance *instancev1.Instance) error
	GetInstance(id string) (*instancev1.Instance, error)
	DeleteInstance(id string) error
}

func NewService(client dynamic.Interface) Service {
	return &kube{
		client: client,
	}
}

type kube struct {
	client dynamic.Interface
}

func (k *kube) CreateInstance(ctx context.Context, instance *instancev1.Instance) error {
	obj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(instance)
	if err != nil {
		return err
	}

	u := &unstructured.Unstructured{}
	u.SetUnstructuredContent(obj)

	if _, err := k.client.Resource(gvr).
		Namespace(instance.Namespace).
		Create(ctx, u, metav1.CreateOptions{}); err != nil {
		return err
	}

	watcher, err := k.client.Resource(gvr).Watch(ctx, metav1.ListOptions{
		FieldSelector: fmt.Sprintf("metadata.name=%s", instance.Name),
	})
	if err != nil {
		return err
	}

	defer watcher.Stop()

	for event := range watcher.ResultChan() {
		if event.Type != watch.Modified {
			continue
		}
		if watched, ok := event.Object.(*unstructured.Unstructured); ok {
			var instanceObj instancev1.Instance
			if err := runtime.DefaultUnstructuredConverter.
				FromUnstructured(watched.UnstructuredContent(), &instanceObj); err != nil {
				return err
			}

			// If an IP has been assigned to the Instance we know that the actual
			// server has been created
			if len(instanceObj.Status.IP) > 0 {
				log.Printf("Instance has been scheduled with IP %s\n", instanceObj.Status.IP)
				break
			}
		}
	}

	return nil
}

func (k *kube) GetInstance(id string) (*instancev1.Instance, error) {
	return nil, nil
}

func (k *kube) DeleteInstance(id string) error {
	return nil
}
