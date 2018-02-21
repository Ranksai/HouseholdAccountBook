package handler

import (
	"net/http"

	"strconv"

	"fmt"

	"github.com/Ranksai/HouseholdAccountBook/src/model/row"
	"github.com/Ranksai/HouseholdAccountBook/src/xorm"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func InitAccountGroupHandler(e *echo.Group) {
	e.POST("/get/:account_id", AccountGet)
	e.POST("/list", AccountList)
	e.POST("/:account_id/item/:item_id/account_type/:account_type_id/save", AccountItemAccountTypeSave)
}

func AccountGet(c echo.Context) error {
	idStr := c.Param("account_id")
	account := new(row.Account)
	xormEngine := xorm.NewXormEngine()

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	has, err := xormEngine.ID(id).Get(account)
	if err != nil {
		log.Fatal(err)
		return c.NoContent(http.StatusInternalServerError)
	} else if !has {
		return c.NoContent(http.StatusNotFound)
	}

	fmt.Println(account)
	return c.JSON(http.StatusOK, account)
}

func AccountList(c echo.Context) error {
	return nil
}

func AccountItemAccountTypeSave(c echo.Context) error {
	return nil
}
