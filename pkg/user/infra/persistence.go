package infra

import (
	"database/sql"
	"fmt"
	"github.com/kenesparta/questions-answers/user/domain"
	"log"
)

type UserPersistence struct {
	tableName string
	db        *sql.DB
	domain.UserRepository
}

func NewUserPersistence(tableName string, db *sql.DB) *UserPersistence {
	return &UserPersistence{
		tableName: tableName,
		db:        db,
	}
}

func (up *UserPersistence) GetAll() ([]domain.User, error) {
	var (
		u         domain.User
		users     []domain.User
		rows, err = up.db.Query(fmt.Sprintf("SELECT * FROM %s", up.tableName))
	)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&u.Id, &u.Name)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
