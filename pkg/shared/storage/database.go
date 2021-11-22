package storage

import (
	"database/sql"
	"fmt"
	"github.com/kenesparta/questions-answers/config"
	pgQuestion "github.com/kenesparta/questions-answers/question/infra"
	pgUser "github.com/kenesparta/questions-answers/user/infra"
	_ "github.com/lib/pq"
	"log"
)

type PostgresRepository struct {
	Question *pgQuestion.QuestionPersistence
	User     *pgUser.UserPersistence
}

func NewPostgresRepository(dbc *config.Database) (*PostgresRepository, error) {
	db, err := sql.Open(
		dbc.Driver,
		fmt.Sprintf(
			`host=%s port=%s user=%s dbname=%s sslmode=disable password=%s`,
			dbc.Host,
			"5432",
			dbc.User,
			dbc.Name,
			dbc.Password),
	)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &PostgresRepository{
		Question: pgQuestion.NewQuestionPersistence("question", db),
		User:     pgUser.NewUserPersistence("user_qa", db),
	}, err
}
