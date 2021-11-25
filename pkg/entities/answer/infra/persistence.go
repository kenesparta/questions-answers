package infra

import (
	"database/sql"
	"fmt"
	"github.com/kenesparta/questions-answers/entities/answer/domain"
	"log"
)

type AnswerPersistence struct {
	tableName string
	db        *sql.DB
	domain.AnswerRepository
}

func NewAnswerPersistence(tableName string, db *sql.DB) *AnswerPersistence {
	return &AnswerPersistence{
		tableName: tableName,
		db:        db,
	}
}

func (ap *AnswerPersistence) update(a domain.Answer) (*string, error) {
	var (
		idAns string
		err   = ap.db.
			QueryRow(fmt.Sprintf("UPDATE %s SET content=$1, user_id=$2 WHERE question_id=$3 RETURNING id", ap.tableName),
				a.Content,
				a.UserId,
				a.QuestionId,
			).
			Scan(&idAns)
	)
	if err != nil {
		return nil, err
	}
	return &idAns, nil
}

func (ap *AnswerPersistence) insert(a domain.Answer) (*string, error) {
	var (
		idAns string
		err   = ap.db.
			QueryRow(
				fmt.Sprintf("INSERT INTO %s VALUES (DEFAULT, $1, $2, $3) RETURNING id", ap.tableName),
				a.Content,
				a.UserId,
				a.QuestionId,
			).
			Scan(&idAns)
	)
	if err != nil {
		return nil, err
	}
	return &idAns, nil
}

func (ap *AnswerPersistence) GetAll() ([]domain.Answer, error) {
	var (
		a         domain.Answer
		answers   []domain.Answer
		rows, err = ap.db.Query(fmt.Sprintf("SELECT * FROM %s", ap.tableName))
	)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&a.Id, &a.Content, &a.UserId, &a.QuestionId)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		answers = append(answers, a)
	}

	return answers, err
}

func (ap *AnswerPersistence) Save(a domain.Answer) (*string, error) {
	// Verify if the answer exists
	var (
		ansCount int
		err      = ap.db.
			QueryRow(
				fmt.Sprintf("SELECT count(1) AS ansCount FROM %s WHERE question_id=$1", ap.tableName),
				a.QuestionId,
			).
			Scan(&ansCount)
	)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	
	if ansCount == 0 {
		return ap.insert(a)
	}

	return ap.update(a)
}
