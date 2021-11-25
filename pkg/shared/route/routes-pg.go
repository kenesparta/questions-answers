package route

import (
	"github.com/gorilla/mux"
	apiAnswer "github.com/kenesparta/questions-answers/entities/answer/route"
	apiQuestion "github.com/kenesparta/questions-answers/entities/question/route"
	apiUser "github.com/kenesparta/questions-answers/entities/user/route"
	"github.com/kenesparta/questions-answers/shared/storage"
)

func InitPostgresRoutes(s *storage.PostgresRepository, r *mux.Router) {
	apiQuestion.NewQuestionHandler(s, r)
	apiUser.NewUserHandler(s, r)
	apiAnswer.NewAnswerHandler(s, r)
}
