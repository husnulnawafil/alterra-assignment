package routes

import (
	"alterra/test/delivery/handler"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, userHandler *handler.UserHandler) {
	e.POST("/users", userHandler.CreateUserHandler())
	e.GET("/users", userHandler.GetListUsersHandler())
	e.DELETE("/users/:id", userHandler.DeleteUserHandler())
	e.PUT("/users/:id", userHandler.UpdateUserHandler())
}
