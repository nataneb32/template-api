package user

import (
	"github.com/nataneb32/template-api/entities"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (it *User) To() *entities.User {
	return &entities.User{
		Username: it.Username,
		Password: it.Password,
	}
}

func From(user *entities.User) *User {
	return &User{
		Username: user.Username,
		Password: user.Password,
	}
}
