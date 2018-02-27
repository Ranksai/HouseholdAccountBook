package handler

import (
	"github.com/labstack/echo"

	"net/http"

	"github.com/Ranksai/HouseholdAccountBook/src/model/row"
	"github.com/Ranksai/HouseholdAccountBook/src/xorm"
	"github.com/labstack/gommon/log"
)

func InitItemGroupHandler(e *echo.Group) {
	e.POST("/get/:id", ItemGet)
	e.POST("/list", ItemList)
	e.POST("/save/account/:account_id", ItemSaveAccount)
	e.POST("/delete/:id", ItemDelete)
}

func ItemGet(c echo.Context) error {
	id := c.Param("id")
	item := new(row.Item)
	xormEngine := xorm.NewXormEngine()

	has, err := xormEngine.ID(id).Get(item)
	if err != nil {
		log.Fatal(err)
		return c.NoContent(http.StatusInternalServerError)
	} else if !has {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, item)
}
func ItemList(c echo.Context) error {
	return nil
}
func ItemSaveAccount(c echo.Context) error {
	return nil
}
func ItemDelete(c echo.Context) error {
	return nil
}
