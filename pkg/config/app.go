package config

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "satya:Bookstore123/simplerest?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	db=d
}

func Getdb() *gorm.DB{
	return db
}
