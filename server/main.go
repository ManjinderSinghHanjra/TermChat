package main

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Welcome to TestApp")
	fmt.Println("Initiating the server")

	// TODO: Add implementation
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	GetSystem().Init()
	GetSystem().Start()
	defer GetSystem().Stop()

	for {
		// TODO: find some other way to handle this
		time.Sleep(10)
	}
}
