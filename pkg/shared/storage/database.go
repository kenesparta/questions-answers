package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"questionsAnswers/config"
	pgQuestion "questionsAnswers/question/infra"
)

type PostgresRepository struct {
	Question *pgQuestion.QuestionPersistence
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
	}, err
}
