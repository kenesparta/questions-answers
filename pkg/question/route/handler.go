package route

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/kenesparta/questions-answers/question/domain"
	"github.com/kenesparta/questions-answers/user/infra"
	"log"
	"net/http"
)

type Question struct {
	app domain.QuestionRepository
}

func New(app domain.QuestionRepository) *Question {
	return &Question{
		app: app,
	}
}

func (q *Question) Get(w http.ResponseWriter, r *http.Request) {
	infra.Headers(w)
	questionId := mux.Vars(r)["id"]
	if _, err := uuid.Parse(questionId); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	question, err := q.app.Get(questionId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&question); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (q *Question) GetAll(w http.ResponseWriter, _ *http.Request) {
	infra.Headers(w)
	questions, err := q.app.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&questions); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

