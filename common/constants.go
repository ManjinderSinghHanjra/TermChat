package common

import "time"

const (
	// DefaultRPCServerAddress : Comment
	DefaultRPCServerAddress = "127.0.0.1"

	// DefaultRPCPort : Comment
	DefaultRPCPort = "9876"
)

const (
	// Protocol : Comment
	Protocol = "protocol"

	// Address : Comment
	Address = "address"

	// Port : Comment
	Port = "port"

	// DBConnectionString : Comment
	DBConnectionString = "db_connection_string"
)

const (
	UserInteractionService = "user-interaction-service"
)

const (
	LIGHT_RPC_TIMEOUT  = 5 * time.Second
	MEDIUM_RPC_TIMEOUT = 15 * time.Second
	HEAVY_RPC_TIMEOUT  = 30 * time.Second
)
