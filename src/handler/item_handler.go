package handler

import (
	"github.com/labstack/echo"

	"net/http"

	"strconv"

	"github.com/Ranksai/HouseholdAccountBook/src/model/entity"
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

func ItemSave(c echo.Context) error {
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

	_, err = xormEngine.Where("name = ?", itemRow.Name).Get(itemRow)
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
	// TODO: itemId
	accountSaveRow := new(row.AccountItem)
	accountSaveRow.AccountId = accountId

	xormEngine := xorm.NewXormEngine()

	accountRow := new(row.Account)
	accountRow.Id = accountId
	//TODo: found
	_, err = xormEngine.Id(accountRow.Id).Get(accountRow)
	if err != nil {
		log.Error(err)
		return c.NoContent(http.StatusNotFound)
	}

	itemRow := new(row.Item)
	itemRow.Name = c.FormValue("name")
	itemRow.Description = c.FormValue("description")
	itemRow.Amount, _ = strconv.Atoi(c.FormValue("amount"))

	insertNum, err := xormEngine.Insert(itemRow)
	if err != nil {
		log.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	} else if insertNum == 0 {
		return c.NoContent(http.StatusBadRequest)
	}

	_, err = xormEngine.Where("name = ?", itemRow.Name).Get(itemRow)
	if err != nil {
		log.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	} else if insertNum == 0 {
		return c.NoContent(http.StatusBadRequest)
	}

	accountSaveRow.ItemId = itemRow.Id
	insertNum, err = xormEngine.Insert(accountSaveRow)
	if err != nil {
		log.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	} else if insertNum == 0 {
		return c.NoContent(http.StatusBadRequest)
	}

	items := make(row.Items, 0)

	err = xormEngine.
		Join("INNER", "account_item", "account_item.item_id = item.id").
		Where("account_id = ?", accountSaveRow.AccountId).
		Find(&items)
	if err != nil {
		log.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	account := new(entity.Account)
	account.Account = *accountRow
	account.Items = items
	for i := range items {
		account.Amount += items[i].Amount
	}
	return c.JSON(http.StatusOK, account)
}

func createAccountEntity(c echo.Context, account_id int) (*entity.Account, error) {
	items := make(row.Items, 0)
	xormEngine := xorm.NewXormEngine()

	err := xormEngine.
		Join("INNER", "account_item", "account_item.item_id = item.id").
		Where("account_id = ?", account_id).
		Find(&items)
	if err != nil {
		log.Error(err)
		return nil, c.NoContent(http.StatusInternalServerError)
	}
	accountRow := new(row.Account)
	accountRow.Id = account_id
	_, err = xormEngine.Id(accountRow.Id).Get(accountRow)
	if err != nil {
		log.Error(err)
		return nil, c.NoContent(http.StatusNotFound)
	}

	account := new(entity.Account)
	account.Account = *accountRow
	account.Items = items
	for i := range items {
		account.Amount += items[i].Amount
	}
	return account, nil
}

func ItemDelete(c echo.Context) error {
	return nil
}
