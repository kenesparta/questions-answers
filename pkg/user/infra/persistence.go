package infra

import "database/sql"

type UserPersistence struct {
	tableName string
	db        *sql.DB
}

func NewUserPersistence(tableName string, db *sql.DB) *UserPersistence {
	return &UserPersistence{
		tableName: tableName,
		db:        db,
	}
}
