package infra

import (
	"database/sql"
	"fmt"
	"github.com/kenesparta/questions-answers/entities/question/domain"
	"log"
)

type QuestionPersistence struct {
	tableName string
	db        *sql.DB
	domain.QuestionRepository
}

func NewQuestionPersistence(tableName string, db *sql.DB) *QuestionPersistence {
	return &QuestionPersistence{
		tableName: tableName,
		db:        db,
	}
}

func (qp *QuestionPersistence) Get(id string) (*domain.Question, error) {
	var (
		q   domain.Question
		err = qp.db.
			QueryRow(fmt.Sprintf("SELECT id, content FROM %s WHERE id=$1", qp.tableName), id).
			Scan(&q.Id, &q.Content)
	)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return &q, nil
}

func (qp *QuestionPersistence) GetAll() ([]domain.Question, error) {
	var (
		q         domain.Question
		questions []domain.Question
		rows, err = qp.db.Query(fmt.Sprintf("SELECT * FROM %s", qp.tableName))
	)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&q.Id, &q.Content, &q.UserId)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		questions = append(questions, q)
	}
	return questions, nil
}

func (qp *QuestionPersistence) GetByUser(userId string) ([]domain.Question, error) {
	var (
		q         domain.Question
		questions []domain.Question
		rows, err = qp.db.Query(
			fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", qp.tableName),
			userId,
		)
	)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&q.Id, &q.Content, &q.UserId)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		questions = append(questions, q)
	}
	return questions, nil
}

func (qp *QuestionPersistence) Save(q domain.Question) (*string, error) {
	var (
		idCreated string
		err       = qp.db.
				QueryRow(
				fmt.Sprintf("INSERT INTO %s VALUES (DEFAULT, $1, $2) RETURNING id", qp.tableName),
				q.Content,
				q.UserId,
			).
			Scan(&idCreated)
	)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return &idCreated, nil
}

func (qp *QuestionPersistence) Update(q domain.Question) error {
	var (
		idUpdated string
		err       = qp.db.
				QueryRow(fmt.Sprintf("UPDATE %s SET content=$1 WHERE id=$2 RETURNING id", qp.tableName),
				q.Content,
				q.Id,
			).
			Scan(&idUpdated)
	)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func (qp *QuestionPersistence) Delete(id string) error {
	var (
		idDeleted string
		err       = qp.db.
				QueryRow(fmt.Sprintf("DELETE FROM %s WHERE id=$1 RETURNING id", qp.tableName), id).
				Scan(&idDeleted)
	)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
