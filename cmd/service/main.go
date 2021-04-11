package main

import (
	"context"
	"log"

	instancev1 "github.com/cownetwork/instance-controller/api/v1"
	"github.com/cownetwork/instance-service/kubernetes"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	// TODO: read config from environment variables

	log.Println("HELLO!!")

	/*
		listener, err := net.Listen("tcp", ":1337")
		if err != nil {
			// TODO: log
			os.Exit(1)
		}*/

	client, err := kubernetes.CreateClient()
	if err != nil {
		log.Fatal(err)
	}

	kubesvc := kubernetes.NewService(client)

	i := &instancev1.Instance{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Instance",
			APIVersion: "instance.cow.network/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test",
			Namespace: "default",
		},
		Spec: instancev1.InstanceSpec{
			Template: v1.PodSpec{
				Containers: []v1.Container{
					{
						Name:  "testtesttest",
						Image: "paulbouwer/hello-kubernetes:1.9",
						Ports: []v1.ContainerPort{
							{
								ContainerPort: 8080,
							},
						},
					},
				},
			},
		},
	}

	err = kubesvc.CreateInstance(context.Background(), i)
	if err != nil {
		log.Fatal(err)
	}

	//grpcServer := transport.New(kubesvc)

	/*
		if err := grpcServer.Serve(listener); err != nil {
			// TODO: log err
			os.Exit(1)
		}*/
}
