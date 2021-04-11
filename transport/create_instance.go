package transport

import (
	"context"

	"github.com/cownetwork/instance-service/endpoint"
	instanceapiv1 "github.com/cownetwork/mooapis-go/cow/instance/v1"
)

func (s *grpcServer) Create(ctx context.Context, req *instanceapiv1.CreateInstanceRequest) (*instanceapiv1.CreateInstanceResponse, error) {
	_, resp, err := s.createHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err // TODO: log
	}
	return resp.(*instanceapiv1.CreateInstanceResponse), nil
}

func decodeCreateInstanceRequest(_ context.Context, req interface{}) (interface{}, error) {
	create := req.(*instanceapiv1.CreateInstanceRequest)
	return &endpoint.CreateInstanceRequest{
		Manifest: create.Manifest,
	}, nil
}

func encodeCreateInstanceReponse(_ context.Context, resp interface{}) (interface{}, error) {
	create := resp.(endpoint.CreateInstanceResponse)
	return &instanceapiv1.CreateInstanceResponse{
		Instance: instanceToProto(create.Instance),
	}, nil
}
