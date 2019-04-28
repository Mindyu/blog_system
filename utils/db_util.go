package utils

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:mindyu55@/blog_system?charset=utf8&parseTime=True&loc=Asia/Shanghai")
	if err == nil {
		return db, err
	}
	log.Errorf("init db fail! err: %s", err.Error())
	return nil, err
}
