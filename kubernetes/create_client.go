package kubernetes

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

// CreateClient creates an in-cluster dynamic kubernetes client
func CreateClient() (dynamic.Interface, error) {
	conf, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	client, err := dynamic.NewForConfig(conf)
	if err != nil {
		return nil, err
	}
	return client, nil
}
