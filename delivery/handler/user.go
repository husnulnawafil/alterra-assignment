package handler

import (
	"alterra/test/delivery/helper"
	"alterra/test/delivery/middlewares"
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
		// check login status
		_, errToken := middlewares.ExtractUserID(ctx)
		if errToken != nil {
			return ctx.JSON(http.StatusUnauthorized, helper.ResponseFailed("Unauthorized"))
		}
		// check role
		role, errRole := middlewares.ExtractRole(ctx)
		if errRole != nil {
			return ctx.JSON(http.StatusUnauthorized, helper.ResponseFailed("Unauthorized"))
		}
		if role != "admin" {
			return ctx.JSON(http.StatusUnauthorized, helper.ResponseFailed("Unauthorized"))
		}
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
		// check login status
		_, errToken := middlewares.ExtractUserID(ctx)
		if errToken != nil {
			return ctx.JSON(http.StatusUnauthorized, helper.ResponseFailed("Unauthorized"))
		}
		search := ctx.QueryParam("search")
		listUsers, err := userHandler.UserUseCase.GetListUsers(search)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get list of users"))
		}
		return ctx.JSON(http.StatusOK, helper.ResponseSuccess("Successfully get list of users", listUsers))
	}
}

func (userHandler *UserHandler) DeleteUserHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// check login status
		_, errToken := middlewares.ExtractUserID(ctx)
		if errToken != nil {
			return ctx.JSON(http.StatusUnauthorized, helper.ResponseFailed("Unauthorized"))
		}
		// check role
		role, errRole := middlewares.ExtractRole(ctx)
		if errRole != nil {
			return ctx.JSON(http.StatusUnauthorized, helper.ResponseFailed("Unauthorized"))
		}
		if role != "admin" {
			return ctx.JSON(http.StatusUnauthorized, helper.ResponseFailed("Unauthorized"))
		}
		userID, _ := strconv.Atoi(ctx.Param("id"))
		if err := userHandler.UserUseCase.DeleteUser(userID); err != nil {
			return ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to delete user"))
		}
		return ctx.JSON(http.StatusOK, helper.ResponseSuccessWithoutData(fmt.Sprintf("Successfully delete user with id %v", userID)))
	}
}

func (userHandler *UserHandler) UpdateUserHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// check login status
		_, errToken := middlewares.ExtractUserID(ctx)
		if errToken != nil {
			return ctx.JSON(http.StatusUnauthorized, helper.ResponseFailed("Unauthorized"))
		}
		// check role
		role, errRole := middlewares.ExtractRole(ctx)
		if errRole != nil {
			return ctx.JSON(http.StatusUnauthorized, helper.ResponseFailed("Unauthorized"))
		}
		if role != "admin" {
			return ctx.JSON(http.StatusUnauthorized, helper.ResponseFailed("Unauthorized"))
		}
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

func (userHandler *UserHandler) GetUserByIdHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// check login status
		_, errToken := middlewares.ExtractUserID(ctx)
		if errToken != nil {
			return ctx.JSON(http.StatusUnauthorized, helper.ResponseFailed("Unauthorized"))
		}
		userID, _ := strconv.Atoi(ctx.Param("id"))
		user, err := userHandler.UserUseCase.GetUserById(userID)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get user"))
		}
		if user.Email == "" && user.Name == "" {
			return ctx.JSON(http.StatusNotFound, helper.ResponseFailed("User not found"))
		}
		return ctx.JSON(http.StatusOK, helper.ResponseSuccess(fmt.Sprintf("Successfully get user with id %v", userID), user))
	}
}
