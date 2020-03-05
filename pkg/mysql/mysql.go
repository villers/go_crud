package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go_crud/pkg/container"
	"log"
	"os"
)

func Configure(app *container.Container) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "common_databases"
	dbHost := "127.0.0.1:3306"

	db, err := gorm.Open(dbDriver, fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName))

	if err != nil {
		panic(err.Error())
	}

	if app.Debug {
		db.LogMode(true)
		db.SetLogger(log.New(os.Stdout, "\r\n", 0))
	}

	app.DB = db
}
