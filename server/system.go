package main

import "sync"

// ----------------------------------------------------------------------------
// System : This class provides methods to init, start and stop the entire system and its modules
type System struct {
	rpcServer *RpcServer

	protocol           string
	address            string
	port               uint16
	dbConnectionString string
}

// ----------------------------------------------------------------------------
func (this *System) Init() {

	// read config
	if err := this.ReadConfig(); err != nil {
		panic(err)
	}

	// allocate objects
	if err := this.allocateObjects(); err != nil {
		panic(err)
	}

}

// ----------------------------------------------------------------------------
func (this *System) Start() {

}

// ----------------------------------------------------------------------------
func (this *System) Stop() {
	// close all connections
	// shutdown rpc server

}

// ----------------------------------------------------------------------------
func (this *System) ReadConfig() error {

	// TODO: Read these values from env or a file
	//

	// set default values as of now
	this.protocol = "tcp"
	this.address = "localhost"
	this.port = 9876

	this.dbConnectionString = "TODO"

	return nil
}

// ----------------------------------------------------------------------------
func (this *System) GetDBConnectionString() string {
	return ""
}

// ----------------------------------------------------------------------------
func (this *System) allocateObjects() error {

	this.rpcServer = &RpcServer{}
	// Beware, conversion might panic
	this.rpcServer.Init(this.protocol, this.address+":"+string(this.port))
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
