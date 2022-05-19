package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db        *gorm.DB
	username  string
	password  string
	tablename string
)

func Connect() {
	access := fmt.Sprintf("%s:%s/%s?charset=utf8&parseTime=True&loc=Local", username, password, tablename)
	d, err := gorm.Open("mysql", access)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db

}
