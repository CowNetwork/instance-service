package transport

import (
	"context"

	instanceapiv1 "github.com/cownetwork/mooapis-go/cow/instance/v1"
)

func (s *grpcServer) Delete(ctx context.Context, req *instanceapiv1.DeleteInstanceRequest) (*instanceapiv1.DeleteInstanceResponse, error) {
	return nil, nil
}
