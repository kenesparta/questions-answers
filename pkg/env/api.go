package env

import (
	"log"
	"os"
)

type Api struct {
	Port string
}

func (a *Api) Set() {
	a.Port = os.Getenv("API_PORT")
	if a.Port == "" {
		log.Fatalf("Wrong environment")
	}
}
