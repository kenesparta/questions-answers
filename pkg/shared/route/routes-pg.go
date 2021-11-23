package route

import (
	"github.com/gorilla/mux"
	apiAnswer "github.com/kenesparta/questions-answers/answer/route"
	apiQuestion "github.com/kenesparta/questions-answers/question/route"
	"github.com/kenesparta/questions-answers/shared/storage"
	apiUser "github.com/kenesparta/questions-answers/user/route"
)

func InitPostgresRoutes(s *storage.PostgresRepository, r *mux.Router) {
	apiQuestion.NewQuestionHandler(s, r)
	apiUser.NewUserHandler(s, r)
	apiAnswer.NewAnswerHandler(s, r)
}
