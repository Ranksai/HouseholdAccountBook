package main

import (
	"github.com/Ranksai/HouseholdAccountBook/src/handler"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	userGroup := e.Group("/v1/user")
	handler.InitUserHandler(userGroup)

	e.Start(":8080")
}
