package main

import (
	"context"
	"io"
	"k8s-grpc-health-check/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	// Connect to gRPC Server
	serverAddress := "localhost:8888"
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(100*time.Millisecond))
	if err != nil {
		log.Println("Could not connect to gRPC server " + serverAddress + ":" + err.Error())
		return
	}
	// Close gRPC Server connection in the end of this method
	defer conn.Close()
	helloWorldClient := proto.NewHelloWorldServiceClient(conn)
	helloWorldRequest := &proto.HelloWorldRequest{
		Name: "John",
	}
	// Call HelloWorld Service in the gRPC Server
	stream, err := helloWorldClient.HelloWorld(context.Background(), helloWorldRequest)

	// Read and print the gRPC Server Stream (Response)
	for {
		streamData, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error getting the response from server: " + err.Error())
		}
		log.Println(streamData)
	}
	log.Println("End calling HelloWorld on the server")
}
