package config

import (
	"log"
	"os"
)

type Api struct {
	Port string
	General
}

func (a *Api) Set() {
	a.Port = os.Getenv("API_PORT")
	if a.Port == "" {
		log.Fatalf("Wrong environment")
	}
}
