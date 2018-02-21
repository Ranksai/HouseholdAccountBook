package handler

import (
	"net/http"

	"github.com/Ranksai/HouseholdAccountBook/src/model/entity"
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
	id := c.Param("account_id")
	account := new(entity.Account)
	xormEngine := xorm.NewXormEngine()

	has, err := xormEngine.ID(id).Get(account)
	if err != nil {
		log.Fatal(err)
		return c.NoContent(http.StatusInternalServerError)
	} else if !has {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, account)
}

func AccountList(c echo.Context) error {

}

func AccountItemAccountTypeSave(c echo.Context) error {
	return nil
}
