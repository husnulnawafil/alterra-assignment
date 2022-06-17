package handler

import (
	"alterra/test/entities"
	"alterra/test/usecase"

	"github.com/labstack/echo"
)

type UserHandler struct {
	UserUseCase usecase.UserUseCaseInterface
}

func NewUserHandler(userUseCase usecase.UserUseCaseInterface) *UserHandler {
	return &UserHandler{
		UserUseCase: userUseCase,
	}
}

func (userHandler *UserHandler) CreateUserHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var newUser entities.User
		if err := ctx.Bind(newUser); err != nil {
			return
		}
	}
}
