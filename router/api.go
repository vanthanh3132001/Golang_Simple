package router

import (
	"github.com/example/handler"
	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter() {
	user := api.Echo.Group("/user")
	user.POST("/sign-up", api.UserHandler.HandleSignUp)
	//notification.DELETE("/delete/:id", api.NotificationHandler.HandleDeleteNotification)

}
