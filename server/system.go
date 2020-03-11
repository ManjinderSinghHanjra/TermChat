package main

import (
	"log"
	"os"
	"sync"

	"../common"
)

// ----------------------------------------------------------------------------
// System : This class provides methods to init, start and stop the entire system and its modules
type System struct {
	rpcServer *RpcServer

	protocol           string
	address            string
	port               string
	dbConnectionString string
}

// ----------------------------------------------------------------------------
func (this *System) Init() {
	log.Println("Initialising system ...")
	// read config
	this.ReadConfig()

	// allocate objects
	if err := this.allocateObjects(); err != nil {
		panic(err)
	}
	log.Println("System initialized successfully.")

}

// ----------------------------------------------------------------------------
func (this *System) Start() {
	log.Println("Starting system ...")

	log.Println("System started successfully")
}

// ----------------------------------------------------------------------------
func (this *System) Stop() {
	log.Println("Stopping system ...")
	// close all connections
	// shutdown rpc server
	this.rpcServer.Stop()
	log.Println("System stopped successfully")
}

// ----------------------------------------------------------------------------
func (this *System) ReadConfig() {

	log.Println("Reading config from .env file ...")

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

	this.dbConnectionString = os.Getenv(common.DBConnectionString)
	if this.dbConnectionString == "" {
		log.Fatalln("DbConnectionString param is missing. Check the env file.")
		panic("not found")
	}

	log.Printf("Protocol: %s, Address: %s, Port: %s", this.protocol, this.address, this.port)

	log.Println("Config reading has been completed.")
}

// ----------------------------------------------------------------------------
func (this *System) GetDBConnectionString() string {
	return this.dbConnectionString
}

// ----------------------------------------------------------------------------
func (this *System) allocateObjects() error {

	log.Println("Allocating system objects ...")

	this.rpcServer = &RpcServer{}
	// Beware, conversion might panic
	address := this.address + ":" + this.port
	this.rpcServer.Init(this.protocol, address)
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
