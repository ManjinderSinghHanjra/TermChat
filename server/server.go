package main

import (
	"log"
	"net"

	"../common"
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
	log.Println("Initialising rpc server ...")
	this.serverOptions = []grpc.ServerOption{}
	this.server = grpc.NewServer(this.serverOptions...)
	listener, err := net.Listen(protocol, address)

	if err != nil {
		log.Fatalf("Failed to listen with protocol: %s , and address: %s", protocol, address)
	}

	this.listener = &listener
}

func (this *RpcServer) registerService(serviceName string) {
	log.Printf("Registering rpc service : %s", serviceName)
	switch serviceName {
	case common.UserInteractionService:
		TermChatProtos.RegisterUserInteractionServer(this.server, services.GetUserInteractionImplementation())
	default:
		log.Fatalf("Unable to register service with name: ", serviceName)
	}
}

func (this *RpcServer) Start() {
	log.Println("Starting rpc server ...")
	this.server.Serve(*this.listener)
}

func (this *RpcServer) Stop() {
	log.Println("Stopping rpc server ...")
	this.server.Stop()
}
