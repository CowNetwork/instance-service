package transport

import (
	"context"

	"github.com/cownetwork/instance-service/endpoint"
	"github.com/cownetwork/mooapis-go/cow/instance/v1"
	instanceapiv1 "github.com/cownetwork/mooapis-go/cow/instance/v1"
)

func (s *grpcServer) Get(ctx context.Context, req *instanceapiv1.GetInstanceRequest) (*instanceapiv1.GetInstanceResponse, error) {
	_, resp, err := s.create.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err // TODO: log err
	}
	return resp.(*instance.GetInstanceResponse), nil
}

func decodeGetInstanceRequest(_ context.Context, req interface{}) (interface{}, error) {
	get := req.(*instanceapiv1.GetInstanceRequest)
	return &endpoint.GetInstanceRequest{
		ID: get.Id,
	}, nil
}

func encodeGetInstanceResponse(_ context.Context, resp interface{}) (interface{}, error) {
	get := resp.(endpoint.GetInstanceResponse)
	return &instanceapiv1.GetInstanceResponse{
		Instance: instanceToProto(get.Instance),
	}, nil
}
