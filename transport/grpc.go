package transport

import (
	"github.com/cownetwork/instance-service/endpoint"
	"github.com/cownetwork/instance-service/kubernetes"
	"github.com/cownetwork/instance-service/template"
	instanceapiv1 "github.com/cownetwork/mooapis-go/cow/instance/v1"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

type grpcServer struct {
	createHandler grpctransport.Handler
	endHandler    grpctransport.Handler
	getHandler    grpctransport.Handler
}

func New(kubesvc kubernetes.Service, templatesvc template.Service) *grpc.Server {
	grpcServer := &grpcServer{
		getHandler: grpctransport.NewServer(
			endpoint.MakeGetInstanceEndpoint(kubesvc),
			decodeGetInstanceRequest,
			encodeGetInstanceResponse,
		),
		endHandler: grpctransport.NewServer(
			endpoint.MakeEndInstanceEndpoint(kubesvc),
			decodeEndInstanceRequest,
			encodeEndInstanceResponse,
		),
		createHandler: grpctransport.NewServer(
			endpoint.MakeCreateInstanceEndpoint(kubesvc, templatesvc),
			decodeCreateInstanceRequest,
			encodeCreateInstanceReponse,
		),
	}
	handle := grpc.NewServer()
	instanceapiv1.RegisterInstanceServiceServer(handle, grpcServer)
	return handle
}
