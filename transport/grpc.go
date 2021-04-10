package transport

import (
	"github.com/cownetwork/instance-service/endpoint"
	"github.com/cownetwork/instance-service/kubernetes"
	instanceapiv1 "github.com/cownetwork/mooapis-go/cow/instance/v1"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

type grpcServer struct {
	create grpctransport.Handler
}

func New(kubesvc kubernetes.Service) *grpc.Server {
	grpcServer := &grpcServer{
		create: grpctransport.NewServer(
			endpoint.MakeGetInstanceEndpoint(kubesvc),
			decodeGetInstanceRequest,
			encodeGetInstanceResponse,
		),
	}
	handle := grpc.NewServer()
	instanceapiv1.RegisterInstanceServiceServer(handle, grpcServer)
	return handle
}
