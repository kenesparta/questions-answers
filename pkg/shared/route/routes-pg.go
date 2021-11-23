package route

import (
	"github.com/gorilla/mux"
	apiQuestion "github.com/kenesparta/questions-answers/question/route"
	"github.com/kenesparta/questions-answers/shared/storage"
	apiUser "github.com/kenesparta/questions-answers/user/route"
)

func InitPostgresRoutes(s *storage.PostgresRepository, r *mux.Router) {
	apiQuestion.NewQuestionHandler(s, r)
	apiUser.NewUserHandler(s, r)
}
