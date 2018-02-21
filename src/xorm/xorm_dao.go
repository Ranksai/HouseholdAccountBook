package xorm

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/labstack/gommon/log"
)

func NewXormEngine() *xorm.Engine {
	host := os.Getenv("MYSQL_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("MYSQL_PORT")
	if port == "" {
		port = "3306"
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		log.Error(err)
	}
	user := os.Getenv("MYSQL_USER")
	if user == "" {
		user = "root"
	}
	dbname := "household_account_book"

	dsn := fmt.Sprintf(
		"%s:@tcp(%s:%s)/%s",
		user,
		host,
		port,
		dbname,
	)
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		log.Error(err)
		return nil
	}
	//defer engine.Close()

	return engine

}
