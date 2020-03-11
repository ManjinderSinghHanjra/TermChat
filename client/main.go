package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("Welcome to TermChat Client")

	// TODO: Add implementation
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	GetSystem().Init()
	GetSystem().Start()
	go GetSystem().PromptUserInput()

	defer GetSystem().Stop()

	for {
		// TODO: find some other way to handle this
		time.Sleep(10)
	}
}
