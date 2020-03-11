package main

import (
	"context"
	"log"
	"net"

	"../common"
	TermChatProtos "../common/generatedCode"
	"google.golang.org/grpc"
)

type RpcClient struct {
	clientOptions  []grpc.DialOption
	client         *grpc.ClientConn
	listener       *net.Listener
	dialingAddress string
}

func (this *RpcClient) Init(protocol string, address string) {
	this.dialingAddress = protocol + "://" + address
	this.clientOptions = []grpc.DialOption{}
	this.clientOptions = append(this.clientOptions, grpc.WithInsecure())
}

func (this *RpcClient) Start() {
	client, err := grpc.Dial(this.dialingAddress, this.clientOptions...)
	if err != nil {
		// TODO: Change this later
		log.Fatalf("Failed to dial: %v", err)
	}

	this.client = client
}

func (this *RpcClient) Stop() {
	this.client.Close()
}

func (this *RpcClient) getUserInteractionClient() TermChatProtos.UserInteractionClient {
	return TermChatProtos.NewUserInteractionClient(this.client)
}

func (this *RpcClient) GetFriends() {
	ctx, cancel := context.WithTimeout(context.Background(), common.LIGHT_RPC_TIMEOUT)
	defer cancel()

	req := &TermChatProtos.GetFriendsRequest{}
	res, err := this.getUserInteractionClient().GetFriends(ctx, req)

	log.Printf("Response: %v, Error: %v", res, err)

	// TODO: Handle response
}

func (this *RpcClient) AddFriend() {
	ctx, cancel := context.WithTimeout(context.Background(), common.LIGHT_RPC_TIMEOUT)
	defer cancel()

	req := &TermChatProtos.AddFriendRequest{}
	res, err := this.getUserInteractionClient().AddFriend(ctx, req)

	log.Printf("Response: %v, Error: %v", res, err)

	// TODO: Handle response
}

func (this *RpcClient) DeleteFriend() {
	ctx, cancel := context.WithTimeout(context.Background(), common.LIGHT_RPC_TIMEOUT)
	defer cancel()

	req := &TermChatProtos.RemoveFriendRequest{}
	res, err := this.getUserInteractionClient().RemoveFriend(ctx, req)

	log.Printf("Response: %v, Error: %v", res, err)

	// TODO: Handle response
}

func (this *RpcClient) ListUsers() {
	ctx, cancel := context.WithTimeout(context.Background(), common.LIGHT_RPC_TIMEOUT)
	defer cancel()

	req := &TermChatProtos.ListUsersRequest{}
	res, err := this.getUserInteractionClient().ListUsers(ctx, req)

	log.Printf("Response: %v, Error: %v", res, err)

	// TODO: Handle response
}
