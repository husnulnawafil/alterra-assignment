package handler

import (
	"alterra/test/delivery/helper"
	"alterra/test/entities"
	"alterra/test/usecase"
	"fmt"
	"net/http"
	"strconv"

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

func (userHandler *UserHandler) GetListUsersHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		listUsers, err := userHandler.UserUseCase.GetListUsers()
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get list of users"))
		}
		return ctx.JSON(http.StatusOK, helper.ResponseSuccess("Successfully get list of users", listUsers))
	}
}

func (userHandler *UserHandler) DeleteUserHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userID, _ := strconv.Atoi(ctx.Param("id"))
		if err := userHandler.UserUseCase.DeleteUser(userID); err != nil {
			return ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to delete user"))
		}
		return ctx.JSON(http.StatusOK, helper.ResponseSuccessWithoutData(fmt.Sprintf("Successfully delete user with id %v", userID)))
	}
}

func (userHandler *UserHandler) UpdateUserHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var updateUser entities.User
		if errBind := ctx.Bind(&updateUser); errBind != nil {
			return ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Invalid request"))
		}
		userID, _ := strconv.Atoi(ctx.Param("id"))

		if err := userHandler.UserUseCase.UpdateUser(updateUser, userID); err != nil {
			return ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to update user"))
		}
		return ctx.JSON(http.StatusOK, helper.ResponseSuccessWithoutData(fmt.Sprintf("Successfully update user with id %v", userID)))
	}
}
