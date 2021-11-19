package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"questionsAnswers/env"
)

func Postgres(dbc *env.Database) (*sql.DB, error) {
	return sql.Open(
		dbc.Driver,
		fmt.Sprintf(
			`host=%s port=%s user=%s dbname=%s sslmode=disable password=%s`,
			dbc.Host,
			"5432",
			dbc.User,
			dbc.Name,
			dbc.Password),
	)
}
