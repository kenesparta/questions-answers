package domain

type Question struct {
	Id      string `json:"id,omitempty"`
	Content string `json:"content"`
	UserId  string `json:"user_id,omitempty"`
}

type QuestionRepository interface {
	GetByUser(userId string) ([]Question, error)
	Get(string) (*Question, error)
	GetAll() ([]Question, error)
	Save(Question) error
	Update(content, Id string) error
	Delete(string) error
}
