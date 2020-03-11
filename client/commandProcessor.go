package main

import (
	"log"
	"strings"
)

// TODO: Tokenization and other features are pending
func HandleInput(inputCommand string) (message string) {
	log.Println("Input Command is: ", inputCommand)
	switch strings.ToUpper(inputCommand) {
	case "GET": // call get friends service
		// TODO: parse response
		GetSystem().GetRPCClient().GetFriends()

	case "ADD": // call add friends service
		// TODO: parse response
		GetSystem().GetRPCClient().AddFriend()

	case "DELETE": // call delete friends service
		// TODO: parse response
		GetSystem().GetRPCClient().DeleteFriend()

	case "LIST": // call list users service
		// TODO: parse response
		GetSystem().GetRPCClient().ListUsers()

	case "HELP":
		DisplayHelp()

	default:
		message = "Invalid command."
	}

	return
}

func DisplayHelp() {
	log.Println("Available commands: [Get, Add, Find, Message, Quit]. For more info, type command with --help")
}
