package main

import (
	"github.com/nataneb32/template-api/cmd/api/api"
	"github.com/nataneb32/template-api/domain/aplication/user"
	"github.com/nataneb32/template-api/infrastructure/repositories"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=password dbname=template port=5432"
	db, _ := gorm.Open(postgres.Open(dsn))
	repositories.AutoMigrate(db)

	repos := repositories.NewRepos(db)

	api := api.New(user.New(repos.Users))

	api.Run()
}
