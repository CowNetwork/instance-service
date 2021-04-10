package transport

import (
	"context"

	instanceapiv1 "github.com/cownetwork/mooapis-go/cow/instance/v1"
)

func (s *grpcServer) Create(ctx context.Context, req *instanceapiv1.CreateInstanceRequest) (*instanceapiv1.CreateInstanceResponse, error) {
	return nil, nil
}
