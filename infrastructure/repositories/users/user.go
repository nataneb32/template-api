package user

import (
	"github.com/nataneb32/template-api/domain/repositories"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) repositories.User {
	return &repo{db}
}
