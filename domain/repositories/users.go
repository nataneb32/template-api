package repositories

import "github.com/nataneb32/template-api/entities"

type User interface {
	Create(user *entities.User) error
}
