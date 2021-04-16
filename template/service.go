package template

import (
	"fmt"
	"os"

	instancev1 "github.com/cownetwork/instance-controller/api/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
)

type Service interface {
	Get(name string) (*instancev1.Instance, bool)
}

type templateService struct {
	manifests map[string]*instancev1.Instance
}

func NewService(path string) (Service, error) {
	const op = "template/NewService"
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)

	}

	mapping := make(map[string]*instancev1.Instance)

	for _, entry := range entries {
		f, err := os.Open(path + "/" + entry.Name())
		if err != nil {
			return nil, fmt.Errorf("%s: %v", op, err)
		}
		decoder := yaml.NewYAMLToJSONDecoder(f)
		instance := &instancev1.Instance{}
		if err := decoder.Decode(&instance); err != nil {
			return nil, fmt.Errorf("%s: %v", op, err)
		}
		mapping[entry.Name()] = instance
	}

	return &templateService{
		manifests: mapping,
	}, nil
}

func (m *templateService) Get(name string) (*instancev1.Instance, bool) {
	if instance, ok := m.manifests[name]; ok {
		return instance, true
	}
	return nil, false
}
