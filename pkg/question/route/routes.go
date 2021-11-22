package route

import (
	"github.com/gorilla/mux"
	"github.com/kenesparta/questions-answers/shared/storage"
	"net/http"
)

func NewQuestionHandler(s *storage.PostgresRepository, r *mux.Router) {
	q := New(s.Question)
	r.HandleFunc("/question", q.Get).Queries(
		"id", "{id}",
	).Methods(http.MethodGet)
}
