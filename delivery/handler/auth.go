package handler

import (
	"alterra/test/delivery/helper"
	"alterra/test/entities"
	"alterra/test/usecase"
	"errors"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthUseCase usecase.AuthUseCaseInterface
}

func NewAuthHandler(authUseCase usecase.AuthUseCaseInterface) *AuthHandler {
	return &AuthHandler{
		AuthUseCase: authUseCase,
	}
}

func (authHandler *AuthHandler) LoginHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var user entities.User

		if errBind := ctx.Bind(&user); errBind != nil {
			return ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Invalid request"))
		}
		token, errLogin := authHandler.AuthUseCase.Login(user.Email, user.Password)
		if errLogin != nil {
			ctx.JSON(http.StatusBadRequest, helper.ResponseFailed(fmt.Sprintf("%v", errLogin)))
		}
		extract, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte("rahasia"), nil
		})
		if err != nil {
			return errors.New("error to extract token")

		}
		if !extract.Valid {
			return errors.New("Invalid")
		}
		claims := extract.Claims.(jwt.MapClaims)
		role := claims["role"].(string)
		responseToken := map[string]interface{}{
			"token": token,
			"role":  role,
		}
		return ctx.JSON(http.StatusOK, helper.ResponseSuccess("Successfully logged in", responseToken))
	}
}
