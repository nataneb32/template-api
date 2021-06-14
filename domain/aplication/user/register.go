package user

import "github.com/nataneb32/template-api/entities"

type Register interface {
	Register(user *entities.User) error
}

func (it *user) Register(user *entities.User) error {
	return it.repo.Create(user)
}
