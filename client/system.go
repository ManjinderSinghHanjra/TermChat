package main

import (
	"bufio"
	"log"
	"os"
	"sync"

	"../common"
)

// ----------------------------------------------------------------------------
// System : This class provides methods to init, start and stop the entire system and its modules
type System struct {
	rpcClient *RpcClient

	protocol string
	address  string
	port     string
}

// ----------------------------------------------------------------------------
func (this *System) Init() {

	// read config
	this.ReadConfig()

	// allocate objects
	if err := this.allocateObjects(); err != nil {
		panic(err)
	}

}

// ----------------------------------------------------------------------------
func (this *System) Start() {
	this.rpcClient.Start()
}

// ----------------------------------------------------------------------------
func (this *System) Stop() {
	// close all connections
	// shutdown rpc client
	this.rpcClient.Stop()

}

// ----------------------------------------------------------------------------
func (this *System) GetRPCClient() *RpcClient {
	return this.rpcClient
}

// ----------------------------------------------------------------------------
// TODO: This flow can be improved a lot
func (this *System) PromptUserInput() {
	reader := bufio.NewReader(os.Stdin)

	for userCommand := ""; userCommand != "quit"; {
		DisplayHelp()
		userCommand, _ = reader.ReadString('\n')
		HandleInput(userCommand)
	}
}

// ----------------------------------------------------------------------------
func (this *System) ReadConfig() {

	this.protocol = os.Getenv(common.Protocol)
	if this.protocol == "" {
		log.Fatalln("Protocol param is missing. Check the env file.")
		panic("not found")
	}

	this.address = os.Getenv(common.Address)
	if this.address == "" {
		log.Fatalln("Address param is missing. Check the env file.")
		panic("not found")
	}

	// conversion isn't needed
	this.port = os.Getenv(common.Port)

	if this.port == "" {
		log.Fatalln("Port param is missing. Check the env file.")
		panic("not found")
	}

	log.Printf("Protocol: %s, Address: %s, Port: %s", this.protocol, this.address, this.port)
}

// ----------------------------------------------------------------------------
func (this *System) allocateObjects() error {

	this.rpcClient = &RpcClient{}
	address := this.address + ":" + this.port
	this.rpcClient.Init(this.protocol, address)
	return nil
}

// SINGLETON OBJECT
var (
	SystemSingleton *System
	once            sync.Once
)

func GetSystem() *System {
	once.Do(func() {
		SystemSingleton = &System{}
	})

	return SystemSingleton
}
