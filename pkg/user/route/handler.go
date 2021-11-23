package route

import (
	"encoding/json"
	"github.com/kenesparta/questions-answers/shared/header"
	"github.com/kenesparta/questions-answers/user/domain"
	"log"
	"net/http"
)

type User struct {
	app domain.UserRepository
}

func New(app domain.UserRepository) *User {
	return &User{
		app: app,
	}
}

func (u *User) GetAll(w http.ResponseWriter, r *http.Request) {
	header.Headers(w)
	users, err := u.app.GetAll()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if users == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No data"))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&users); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
