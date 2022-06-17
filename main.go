package main

import (
	"alterra/test/configs"
	"fmt"
	"log"

	"github.com/labstack/echo"
)

func main() {
	config := configs.GetConfig()

	e := echo.New()

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}
