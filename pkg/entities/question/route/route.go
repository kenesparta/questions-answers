package route

import (
	"github.com/gorilla/mux"
	"github.com/kenesparta/questions-answers/shared/storage"
	"net/http"
)

func NewQuestionHandler(s *storage.PostgresRepository, r *mux.Router) {
	q := New(s.Question)

	r.HandleFunc("/question/{id}", q.Get).Methods(http.MethodGet)

	r.HandleFunc("/question", q.GetByUser).Queries(
		"userId", "{userId}",
	).Methods(http.MethodGet)

	r.HandleFunc("/question/{id}", q.Delete).Methods(http.MethodDelete)

	r.HandleFunc("/question", q.GetAll).Methods(http.MethodGet)

	r.HandleFunc("/question", q.Save).Methods(http.MethodPost)

	r.HandleFunc("/question", q.Update).Methods(http.MethodPut)
}
