package transport

import (
	"bytes"
	"context"
	"fmt"

	instancev1 "github.com/cownetwork/instance-controller/api/v1"
	"github.com/cownetwork/instance-service/endpoint"
	instanceapiv1 "github.com/cownetwork/mooapis-go/cow/instance/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func (s *grpcServer) Create(ctx context.Context, req *instanceapiv1.CreateInstanceRequest) (*instanceapiv1.CreateInstanceResponse, error) {
	_, resp, err := s.createHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err // TODO: log
	}
	return resp.(*instanceapiv1.CreateInstanceResponse), nil
}

func decodeCreateInstanceRequest(_ context.Context, req interface{}) (interface{}, error) {
	const op = "transport/decodeCreateInstanceRequest"
	create := req.(*instanceapiv1.CreateInstanceRequest)

	switch x := create.Instance.(type) {
	case *instanceapiv1.CreateInstanceRequest_Info:
		return endpoint.CreateInstanceRequest{
			TemplateName:         create.GetInfo().TemplateName,
			TemplateInstanceName: create.GetInfo().GetInstanceName(),
		}, nil
	case *instanceapiv1.CreateInstanceRequest_Manifest:
		jsonbytes, err := protojson.Marshal(create)
		if err != nil {
			return nil, fmt.Errorf("%s: %v", op, err)
		}
		decoder := yaml.NewYAMLToJSONDecoder(bytes.NewReader(jsonbytes))
		instance := &instancev1.Instance{}
		if err := decoder.Decode(instance); err != nil {
			return nil, fmt.Errorf("%s: %v", op, err)
		}
		return endpoint.CreateInstanceRequest{
			Instance: instance,
		}, nil
	case nil:
		return nil, fmt.Errorf("%s: Instance field must be set", op)
	default:
		return nil, fmt.Errorf("%s: unknown type for Instance field (%T)", op, x)
	}
}

func encodeCreateInstanceReponse(_ context.Context, resp interface{}) (interface{}, error) {
	create := resp.(endpoint.CreateInstanceResponse)
	return &instanceapiv1.CreateInstanceResponse{
		Instance: instanceToProto(create.Instance),
	}, nil
}
