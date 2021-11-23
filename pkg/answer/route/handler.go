package route

import (
	"encoding/json"
	"fmt"
	"github.com/kenesparta/questions-answers/answer/domain"
	"github.com/kenesparta/questions-answers/shared/header"
	"log"
	"net/http"
)

type Answer struct {
	app domain.AnswerRepository
}

func New(app domain.AnswerRepository) *Answer {
	return &Answer{
		app: app,
	}
}

func (a *Answer) GetAll(w http.ResponseWriter, _ *http.Request) {
	header.Headers(w)
	answers, err := a.app.GetAll()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if answers == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No data"))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&answers); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a *Answer) Save(w http.ResponseWriter, r *http.Request) {
	header.Headers(w)
	var answer domain.Answer
	if err := json.NewDecoder(r.Body).Decode(&answer); nil != err {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	id, err := a.app.Save(answer)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`{"id":"%s"}`, *id)))
}