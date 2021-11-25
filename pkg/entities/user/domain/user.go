package domain

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserRepository interface {
	GetAll() ([]User, error)
}
