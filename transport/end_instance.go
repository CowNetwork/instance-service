package transport

import (
	"context"

	"github.com/cownetwork/instance-service/endpoint"
	instanceapiv1 "github.com/cownetwork/mooapis-go/cow/instance/v1"
)

func (s *grpcServer) End(ctx context.Context, req *instanceapiv1.EndInstanceRequest) (*instanceapiv1.EndInstanceResponse, error) {
	_, resp, err := s.endHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err // TODO: log
	}
	return resp.(*instanceapiv1.EndInstanceResponse), nil
}

func decodeEndInstanceRequest(_ context.Context, req interface{}) (interface{}, error) {
	del := req.(*instanceapiv1.EndInstanceRequest)
	return &endpoint.EndInstanceRequest{
		ID: del.Id,
	}, nil
}

func encodeEndInstanceResponse(_ context.Context, resp interface{}) (interface{}, error) {
	del := resp.(endpoint.EndInstanceResponse)
	return &instanceapiv1.DeleteInstanceResponse{
		Instance: instanceToProto(del.Instance),
	}, nil

}
