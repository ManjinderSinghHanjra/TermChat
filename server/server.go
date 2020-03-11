package main

import (
	"log"
	"net"

	TermChatProtos "../common/generatedCode"
	"./services"
	"google.golang.org/grpc"
)

type RpcServer struct {
	serverOptions []grpc.ServerOption
	server        *grpc.Server
	listener      *net.Listener
}

func (this *RpcServer) Init(protocol string, address string) {
	this.serverOptions = []grpc.ServerOption{}
	this.server = grpc.NewServer(this.serverOptions...)

	// Review & TODO: port conversion might panic
	listener, err := net.Listen(protocol, address)

	if err != nil {
		log.Fatalf("Failed to listen with protocol: %s , and address: %s", protocol, address)
	}

	this.listener = &listener
}

func (this *RpcServer) registerService(serviceName string) {
	switch serviceName {
	case "user-interaction-service":
		TermChatProtos.RegisterUserInteractionServer(this.server, services.GetUserInteractionImplementation())
	default:
		log.Fatalf("Unable to register service with name: ", serviceName)
	}
}

func (this *RpcServer) Start() {
	this.server.Serve(*this.listener)
}

func (this *RpcServer) Stop() {

}
