package route

import (
	"github.com/gorilla/mux"
	apiQuestion "questionsAnswers/question/route"
	"questionsAnswers/shared/storage"
)

func InitPostgresRoutes(s *storage.PostgresRepository, r *mux.Router) {
	apiQuestion.NewQuestionHandler(s, r)
}
