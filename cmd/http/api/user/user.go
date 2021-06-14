package user

import (
	"github.com/labstack/echo/v4"
	"github.com/nataneb32/template-api/domain/aplication/user"
)

type UserHandler struct {
	service user.User
}

func New(service user.User) *UserHandler {
	return &UserHandler{
		service,
	}
}

func (it *UserHandler) Register(r *echo.Echo) {
	r.POST("/app/register", it.RegisterHandler)
}
