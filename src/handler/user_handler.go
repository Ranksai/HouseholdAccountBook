package handler

import (
	"fmt"

	"github.com/labstack/echo"
)

func InitUserHandler(e *echo.Group) {
	e.Any("/register", userRegister)
	e.Any("/login", userLogin)
	e.Any("/logout", userLogout)

}

func userRegister(c echo.Context) error {
	fmt.Println(c.Path())
	return nil
}

func userLogin(c echo.Context) error {
	fmt.Println(c.Path())
	return nil
}

func userLogout(c echo.Context) error {
	fmt.Println(c.Path())
	return nil
}
