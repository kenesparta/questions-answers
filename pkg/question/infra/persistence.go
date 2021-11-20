package infra

import (
	"database/sql"
)

type QuestionPersistence struct {
	tableName string
	db        *sql.DB
}

func NewQuestionPersistence(tableName string, db *sql.DB) *QuestionPersistence {
	return &QuestionPersistence{
		tableName: tableName,
		db:        db,
	}
}