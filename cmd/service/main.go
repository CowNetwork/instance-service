package main

import (
	"net"
	"os"

	"github.com/cownetwork/instance-service/kubernetes"
	"github.com/cownetwork/instance-service/transport"
)

func main() {
	// TODO: read config from environment variables

	listener, err := net.Listen("tcp", ":1337")
	if err != nil {
		// TODO: log
		os.Exit(1)
	}

	kubesvc := kubernetes.NewService()
	grpcServer := transport.New(kubesvc)

	if err := grpcServer.Serve(listener); err != nil {
		// TODO: log err
		os.Exit(1)
	}
}
