package routes

import (
	"alterra/test/delivery/handler"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, userHandler *handler.UserHandler, authHandler *handler.AuthHandler) {
	// AUTHENTICATION ENDPOINTS
	e.POST("/login", authHandler.LoginHandler())

	// USER ENDPOINTS
	e.POST("/users", userHandler.CreateUserHandler())
	e.GET("/users", userHandler.GetListUsersHandler())
	e.GET("/users/:id", userHandler.GetUserByIdHandler())
	e.DELETE("/users/:id", userHandler.DeleteUserHandler())
	e.PUT("/users/:id", userHandler.UpdateUserHandler())
}
