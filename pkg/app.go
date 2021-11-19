package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	dbc "questionsAnswers/db"
	"questionsAnswers/env"
)

type app struct {
	apiConfig env.Api
	dbConfig  env.Database
	router    *mux.Router
}

func (a *app) Initialize() {
	a.apiConfig.Set()
	a.dbConfig.Set()
	a.router = mux.NewRouter()
	pg, err := dbc.Postgres(&a.dbConfig)
	if err != nil {
		return
	}
	fmt.Println(pg)
}

func (a *app) Run() {
	if err := http.ListenAndServe(
		fmt.Sprintf(":%s", a.apiConfig.Port),
		a.router); err != nil {
		log.Fatalln(err)
	}
}
