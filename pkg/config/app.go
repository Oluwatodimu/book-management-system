package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dbConn *gorm.DB

func Connect() {
	var err error
	dbConn, err = gorm.Open("mysql", "root:Zawarudo12!@@tcp(127.0.0.1:3306)/book-inventory")
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return dbConn
}
