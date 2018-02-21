package main

import (
	"github.com/Ranksai/HouseholdAccountBook/src/handler"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	userGroup := e.Group("/user")
	handler.InitUserHandler(userGroup)
	accountGroup := e.Group("/account")
	handler.InitAccountGroupHandler(accountGroup)

	e.Start(":8080")
}
