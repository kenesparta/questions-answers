package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kenesparta/questions-answers/config"
	"github.com/kenesparta/questions-answers/shared/route"
	"github.com/kenesparta/questions-answers/shared/storage"
	"log"
	"net/http"
)

type app struct {
	apiConfig config.Api
	dbConfig  config.Database
	router    *mux.Router
}

func (a *app) initPostgres() {
	pg, err := storage.NewPostgresRepository(&a.dbConfig)
	if err != nil {
		log.Fatalln(err)
	}
	route.InitPostgresRoutes(pg, a.router)
}

func (a *app) initialize() {
	a.apiConfig.Set()
	a.dbConfig.Set()
	a.router = mux.NewRouter()
	a.initPostgres()
}

func (a *app) run() {
	if err := http.ListenAndServe(
		fmt.Sprintf(":%s", a.apiConfig.Port),
		a.router); err != nil {
		log.Fatalln(err)
	}
}
