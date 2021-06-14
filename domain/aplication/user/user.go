package user

import (
	"github.com/nataneb32/template-api/domain/repositories"
)

type User interface {
	Register
}

type user struct {
	repo repositories.User
}

func New(repo repositories.User) User {
	return &user{repo}
}
