package main

import (
	"fmt"
	"log"
	"os"
)

type Environment struct {}

func (Environment) GetPort() string {
	port := os.Getenv("API_PORT")
	if port == "" {
		log.Fatalf("Port is not set")
	}
	return fmt.Sprintf(":%s", port)
}