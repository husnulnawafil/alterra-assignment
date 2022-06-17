package routes

import (
	"alterra/test/delivery/handler"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, userHandler *handler.UserHandler) {
	e.POST("/user", userHandler.CreateUserHandler())
}
