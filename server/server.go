package main

import (
	"fmt"
	"k8s-grpc-health-check/proto"
	"k8s-grpc-health-check/server/healthcheck"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type HelloWorldServiceImpl struct{}

// Implement function HelloWorld interface HelloWorldServiceServer
func (s *HelloWorldServiceImpl) HelloWorld(helloReq *proto.HelloWorldRequest, srv proto.HelloWorldService_HelloWorldServer) error {
	log.Println("Server received an rpc request with the following parameter " + helloReq.GetName())
	resp := &proto.HelloWorldResponse{
		Message: fmt.Sprintf("Hello " + helloReq.GetName()),
	}
	srv.SendMsg(resp)
	return nil
}

func main() {
	// listen for tcp connections
	serverAddress := "localhost:8888"
	listenAddr, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Println("Error while starting the listening service " + err.Error())
	}

	grpcServer := grpc.NewServer()
	// register gRPC proccessor --> HelloWorld function implementation
	proto.RegisterHelloWorldServiceServer(grpcServer, &HelloWorldServiceImpl{})

	// register gRPC health check, we will use in the Kubernetes probes
	healthService := healthcheck.NewHealthChecker()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthService)

	log.Println("Server starting to listen on " + serverAddress)
	if err = grpcServer.Serve(listenAddr); err != nil {
		log.Println("Error while starting the gRPC server:" + err.Error())
	}
}
