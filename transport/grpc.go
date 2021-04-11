package transport

import (
	"github.com/cownetwork/instance-service/endpoint"
	"github.com/cownetwork/instance-service/kubernetes"
	instanceapiv1 "github.com/cownetwork/mooapis-go/cow/instance/v1"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

type grpcServer struct {
	createHandler grpctransport.Handler
	deleteHandler grpctransport.Handler
	getHandler    grpctransport.Handler
}

func New(kubesvc kubernetes.Service) *grpc.Server {
	grpcServer := &grpcServer{
		getHandler: grpctransport.NewServer(
			endpoint.MakeGetInstanceEndpoint(kubesvc),
			decodeGetInstanceRequest,
			encodeGetInstanceResponse,
		),
		deleteHandler: grpctransport.NewServer(
			endpoint.MakeDeleteInstanceEndpoint(kubesvc),
			decodeDeleteInstanceRequest,
			encodeDeleteInstanceResponse,
		),
		createHandler: grpctransport.NewServer(
			endpoint.MakeCreateInstanceEndpoint(kubesvc),
			decodeCreateInstanceRequest,
			encodeCreateInstanceReponse,
		),
	}
	handle := grpc.NewServer()
	instanceapiv1.RegisterInstanceServiceServer(handle, grpcServer)
	return handle
}
