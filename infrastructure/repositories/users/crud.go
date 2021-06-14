package user

import "github.com/nataneb32/template-api/entities"

func (it *repo) Create(user *entities.User) error {
	return it.db.Create(user).Error
}

func (it *repo) Delete(user *entities.User) error {
	return it.db.Create(user).Error
}

func (it *repo) Update(user *entities.User) error {
	return it.db.Create(user).Error
}
