package api

import (
	"github.com/labstack/echo/v4"
	"github.com/nataneb32/template-api/cmd/http/api/user"
	userApp "github.com/nataneb32/template-api/domain/aplication/user"
)

type Api struct {
	handlers []Handlers
}

type Handlers interface {
	Register(g *echo.Echo)
}

func New(
	userApp userApp.User,
) *Api {
	return &Api{
		[]Handlers{
			user.New(userApp),
		},
	}
}

func (it *Api) Run() {
	r := echo.New()

	for _, h := range it.handlers {
		h.Register(r)
	}

	r.Start("127.0.0.1:8080")
}
