package endpoint

import (
	"context"
	"fmt"

	instancev1 "github.com/cownetwork/instance-controller/api/v1"
	"github.com/cownetwork/instance-service/kubernetes"
	"github.com/cownetwork/instance-service/template"
	"github.com/go-kit/kit/endpoint"
)

type CreateInstanceRequest struct {
	TemplateName         string
	TemplateInstanceName string
	Instance             *instancev1.Instance
}

type CreateInstanceResponse struct {
	Instance *instancev1.Instance
}

func MakeCreateInstanceEndpoint(kubesvc kubernetes.Service, templatesvc template.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		const op = "endpoint/MakeCreateInstanceEndpoint"
		create := req.(CreateInstanceRequest)

		var instance *instancev1.Instance

		if create.Instance != nil {
			created, err := kubesvc.CreateInstance(ctx, create.Instance)
			if err != nil {
				return nil, fmt.Errorf("%s: %v", op, err)
			}
			instance = created
		} else if create.TemplateName != "" {
			templated, ok := templatesvc.Get(create.TemplateName)
			if !ok {
				return nil, fmt.Errorf("%s: no template found", op)
			}
			templated.Name = create.TemplateInstanceName
			created, err := kubesvc.CreateInstance(ctx, templated)
			if err != nil {
				return nil, fmt.Errorf("%s: %v", op, err)
			}
			instance = created
		}

		return CreateInstanceResponse{
			Instance: instance,
		}, nil
	}
}
