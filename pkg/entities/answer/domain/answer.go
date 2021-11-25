package domain

type Answer struct {
	Id         string `json:"id"`
	Content    string `json:"content"`
	UserId     string `json:"user_id"`
	QuestionId string `json:"question_id"`
}

type AnswerRepository interface {
	Save(Answer) (*string, error)
	GetAll() ([]Answer, error)
}
