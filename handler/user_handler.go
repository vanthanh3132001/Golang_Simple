package handler

import (
	"github.com/example/model"
	"github.com/example/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type UserHandler struct {
	UserRepo repository.UserRepo
}

func (u UserHandler) HandleSignUp(c echo.Context) error {
	req := model.SignUp{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	role := model.MEMBER.String()
	userId, err := uuid.NewUUID()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})

	}
	user := model.User{
		UserID:   userId.String(),
		FullName: req.FullName,
		Role:     role,
		Password: req.Password,
		Email:    req.Email,
	}

	user, err = u.UserRepo.SaveUser(c.Request().Context(), user)
	if err != nil {

		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})

	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xu ly thanh cong",
		Data:       user,
	})

}
