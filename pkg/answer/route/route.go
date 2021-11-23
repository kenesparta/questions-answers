package route

import (
	"github.com/gorilla/mux"
	"github.com/kenesparta/questions-answers/shared/storage"
	"net/http"
)

func NewAnswerHandler(s *storage.PostgresRepository, r *mux.Router) {
	a := New(s.Answer)
	r.HandleFunc("/answer", a.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/answer", a.Save).Methods(http.MethodPost)
}
