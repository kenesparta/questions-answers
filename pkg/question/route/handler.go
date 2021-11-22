package route

import (
	"encoding/json"
	"fmt"
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

	if question == nil {
		w.WriteHeader(http.StatusNotFound)
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

	if questions == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&questions); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (q *Question) GetByUser(w http.ResponseWriter, r *http.Request) {
	infra.Headers(w)
	userId := mux.Vars(r)["userId"]
	fmt.Println(userId)
	if _, err := uuid.Parse(userId); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	questions, err := q.app.GetByUser(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if questions == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&questions); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (q *Question) Save(w http.ResponseWriter, r *http.Request) {

}