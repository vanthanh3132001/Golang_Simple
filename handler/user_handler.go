package handler

import (
	"github.com/example/repository"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserRepo repository.UserRepo
}

func (u UserHandler) HandleSignUp(c echo.Context) error {
	return nil
}
