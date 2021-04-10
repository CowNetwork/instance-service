package endpoint

import (
	"context"

	instancev1 "github.com/cownetwork/instance-controller/api/v1"
	"github.com/cownetwork/instance-service/kubernetes"
	"github.com/go-kit/kit/endpoint"
)

type GetInstanceRequest struct {
	ID string
}

type GetInstanceResponse struct {
	Instance *instancev1.Instance
}

func MakeGetInstanceEndpoint(svc kubernetes.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		get := req.(GetInstanceRequest)
		instance, err := svc.GetInstance(get.ID)
		if err != nil {
			return nil, err
		}
		return GetInstanceResponse{
			Instance: instance,
		}, nil
	}
}
