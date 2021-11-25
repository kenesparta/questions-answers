package route

import (
	"github.com/gorilla/mux"
	"github.com/kenesparta/questions-answers/shared/storage"
	"net/http"
)

func NewUserHandler(s *storage.PostgresRepository, r *mux.Router) {
	u := New(s.User)
	r.HandleFunc("/user", u.GetAll).Methods(http.MethodGet)
}
