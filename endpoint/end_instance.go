package endpoint

import (
	"context"

	instancev1 "github.com/cownetwork/instance-controller/api/v1"
	"github.com/cownetwork/instance-service/kubernetes"
	"github.com/go-kit/kit/endpoint"
)

type EndInstanceRequest struct {
	ID string
}

type EndInstanceResponse struct {
	Instance *instancev1.Instance
}

func MakeEndInstanceEndpoint(svc kubernetes.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		del := req.(EndInstanceRequest)
		instance, err := svc.DeleteInstance(ctx, del.ID)
		if err != nil {
			return nil, err
		}
		return EndInstanceResponse{
			Instance: instance,
		}, nil
	}
}
