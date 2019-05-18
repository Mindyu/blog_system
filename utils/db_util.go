package utils

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/gommon/log"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:mindyu55@tcp(188.131.213.13:3306)/blog_system?charset=utf8&parseTime=True&loc=Local")
	//db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/blog_system?charset=utf8&parseTime=True&loc=Local")
	if err == nil {
		return db, err
	}
	log.Errorf("init db fail! err: %s", err.Error())
	return nil, err
}
