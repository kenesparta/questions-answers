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
	Save(Question) (*string, error)
	Update(Question) error
	Delete(string) error
}
