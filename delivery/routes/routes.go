package routes

import (
	"alterra/test/delivery/handler"
	"alterra/test/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, userHandler *handler.UserHandler, authHandler *handler.AuthHandler) {
	// AUTHENTICATION ENDPOINTS
	e.POST("/login", authHandler.LoginHandler())

	// USER ENDPOINTS
	e.POST("/users", userHandler.CreateUserHandler(), middlewares.JWTMiddleware())
	e.GET("/users", userHandler.GetListUsersHandler(), middlewares.JWTMiddleware())
	e.GET("/users/:id", userHandler.GetUserByIdHandler(), middlewares.JWTMiddleware())
	e.DELETE("/users/:id", userHandler.DeleteUserHandler(), middlewares.JWTMiddleware())
	e.PUT("/users/:id", userHandler.UpdateUserHandler(), middlewares.JWTMiddleware())
}
