package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type app struct {
	config Environment
	router *mux.Router
}

func (a *app) Initialize() {
	a.router = mux.NewRouter()
}

func (a *app) Run() {
	if err := http.ListenAndServe(a.config.GetPort(), a.router); err != nil {
		log.Fatalln(err)
	}
}
