package handler

import (
	"github.com/labstack/echo"

	"net/http"

	"strconv"

	"github.com/Ranksai/HouseholdAccountBook/src/model/row"
	"github.com/Ranksai/HouseholdAccountBook/src/xorm"
	"github.com/labstack/gommon/log"
)

func InitItemGroupHandler(e *echo.Group) {
	e.POST("/get/:id", ItemGet)
	e.POST("/list", ItemList)
	e.POST("/save", ItemSave)
	e.POST("/save/account/:account_id", ItemSaveAccount)
	e.POST("/delete/:id", ItemDelete)
}

func ItemGet(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err)
		return c.NoContent(http.StatusBadRequest)
	}
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

type itemParam struct {
	name        string `form:"name" query:"name"`
	description string `form:"description" query:"description"`
	amount      int    `form:"amount" query:"amount"`
}

func ItemSave(c echo.Context) error {
	itemParam := new(itemParam)
	if err := c.Bind(itemParam); err != nil {
		log.Error(err)
		return c.NoContent(http.StatusBadRequest)
	}

	itemRow := new(row.Item)
	itemRow.Name = c.FormValue("name")
	itemRow.Description = c.FormValue("description")
	itemRow.Amount, _ = strconv.Atoi(c.FormValue("amount"))

	xormEngine := xorm.NewXormEngine()
	insertNum, err := xormEngine.Insert(itemRow)
	if err != nil {
		log.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	} else if insertNum == 0 {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, itemRow)
}

func ItemSaveAccount(c echo.Context) error {
	accountIdStr := c.Param("account_id")
	accountId, err := strconv.Atoi(accountIdStr)
	if err != nil {
		log.Error(err)
		return c.NoContent(http.StatusBadRequest)
	}
	itemParam := new(itemParam)
	if err := c.Bind(itemParam); err != nil {
		log.Error(err)
		return c.NoContent(http.StatusBadRequest)
	}
	// TODO: itemId
	accountSaveRow := new(row.AccountItem)
	accountSaveRow.AccountId = accountId

	xormEngine := xorm.NewXormEngine()

	insertNum, err := xormEngine.Insert(accountSaveRow)
	if err != nil {
		log.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	} else if insertNum == 0 {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, accountSaveRow)
}
func ItemDelete(c echo.Context) error {
	return nil
}
