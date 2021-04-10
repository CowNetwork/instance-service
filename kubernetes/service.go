package kubernetes

import (
	instancev1 "github.com/cownetwork/instance-controller/api/v1"
)

type Service interface {
	CreateInstance(instance *instancev1.Instance) error
	GetInstance(id string) (*instancev1.Instance, error)
	DeleteInstance(id string) error
}

func NewService() Service {
	return &kube{}
}

type kube struct {
}

func (k *kube) CreateInstance(instance *instancev1.Instance) error {
	return nil
}

func (k *kube) GetInstance(id string) (*instancev1.Instance, error) {
	return nil, nil
}

func (k *kube) DeleteInstance(id string) error {
	return nil
}
