package repositories

import (
	"github.com/nataneb32/template-api/domain/repositories"
	users "github.com/nataneb32/template-api/infrastructure/repositories/users"
	"gorm.io/gorm"
)

type Repos struct {
	Users repositories.User
}

func NewRepos(db *gorm.DB) *Repos {
	u := users.New(db)
	return &Repos{u}
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&users.User{})
}
