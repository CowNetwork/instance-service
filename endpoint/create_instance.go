package endpoint

import (
	"context"

	instancev1 "github.com/cownetwork/instance-controller/api/v1"
	"github.com/cownetwork/instance-service/kubernetes"
	"github.com/go-kit/kit/endpoint"
)

type CreateInstanceRequest struct {
	Manifest string
}

type CreateInstanceResponse struct {
	Instance *instancev1.Instance
}

func MakeCreateInstanceEndpoint(svc kubernetes.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		create := req.(CreateInstanceRequest)
		instance, err := svc.CreateInstance(ctx, create.Instance)
		if err != nil {
			return nil, err
		}
		return CreateInstanceResponse{
			Instance: instance,
		}, nil
	}
}
