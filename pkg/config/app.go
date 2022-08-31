package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

// Helper Method to Connect with MySQL Database
func Connect(){
	d, error := gorm.Open("mysql", "raghadasfour:password/simplerest?charset=utf8&parseTime=True&loc=Local")
	if error != nil{
		panic(error)
	}
	db = d
}

func GetDB()*gorm.DB{
	return db
}