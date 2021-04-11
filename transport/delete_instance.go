package transport

import (
	"context"

	"github.com/cownetwork/instance-service/endpoint"
	instanceapiv1 "github.com/cownetwork/mooapis-go/cow/instance/v1"
)

func (s *grpcServer) Delete(ctx context.Context, req *instanceapiv1.DeleteInstanceRequest) (*instanceapiv1.DeleteInstanceResponse, error) {
	_, resp, err := s.deleteHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err // TODO: log
	}
	return resp.(*instanceapiv1.DeleteInstanceResponse), nil
}

func decodeDeleteInstanceRequest(_ context.Context, req interface{}) (interface{}, error) {
	del := req.(*instanceapiv1.DeleteInstanceRequest)
	return &endpoint.DeleteInstanceRequest{
		ID: del.Id,
	}, nil
}

func encodeDeleteInstanceResponse(_ context.Context, resp interface{}) (interface{}, error) {
	del := resp.(endpoint.DeleteInstanceResponse)
	return &instanceapiv1.DeleteInstanceResponse{
		Instance: instanceToProto(del.Instance),
	}, nil

}
