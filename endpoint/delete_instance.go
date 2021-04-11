package endpoint

import (
	"context"

	instancev1 "github.com/cownetwork/instance-controller/api/v1"
	"github.com/cownetwork/instance-service/kubernetes"
	"github.com/go-kit/kit/endpoint"
)

type DeleteInstanceRequest struct {
	ID string
}

type DeleteInstanceResponse struct {
	Instance *instancev1.Instance
}

func MakeDeleteInstanceEndpoint(svc kubernetes.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		del := req.(DeleteInstanceRequest)
		instance, err := svc.DeleteInstance(ctx, del.ID)
		if err != nil {
			return nil, err
		}
		return DeleteInstanceResponse{
			Instance: instance,
		}, nil
	}
}
