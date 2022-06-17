package handler

import (
	"alterra/test/delivery/helper"
	"alterra/test/entities"
	"alterra/test/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
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
		if errBind := ctx.Bind(&newUser); errBind != nil {
			return ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Invalid request"))
		}
		sign, err := userHandler.UserUseCase.CreateUser(newUser)

		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to create a new user"))
		}

		if sign == 1 {
			return ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Please complete the registration form"))
		}

		return ctx.JSON(http.StatusCreated, helper.ResponseSuccessWithoutData("Successfully create a new user"))
	}
}
