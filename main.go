package main

import (
	"alterra/test/configs"
	"alterra/test/delivery/handler"
	"alterra/test/delivery/routes"
	"alterra/test/repository"
	"alterra/test/usecase"
	"alterra/test/utility"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	config := configs.GetConfig()
	database := utility.InitDB(config)

	userRepository := repository.NewUserRepository(database)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	authRepository := repository.NewAuthRepository(database)
	authUseCase := usecase.NewAuthUseCase(authRepository)
	authHandler := handler.NewAuthHandler(authUseCase)

	e := echo.New()
	routes.RegisterPath(e, userHandler, authHandler)
	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}
