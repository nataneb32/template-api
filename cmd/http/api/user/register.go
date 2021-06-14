package user

import (
	"github.com/labstack/echo/v4"
	"github.com/nataneb32/template-api/entities"
)

func (it *UserHandler) RegisterHandler(c echo.Context) error {
	var r entities.User

	err := c.Bind(&r)
	if err != nil {
		return err
	}

	err = it.service.Register(&r)
	if err != nil {
		return err
	}

	c.JSON(200, r.HidePassword())

	return nil
}
