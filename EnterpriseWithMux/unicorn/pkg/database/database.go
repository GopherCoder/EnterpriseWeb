package database

import (
	"time"

	"github.com/jinzhu/gorm"
)

var Engine *gorm.DB

func EngineInit() {
	engine, _ := gorm.Open("mysql", "root:admin123@/unicorn?charset=utf8&parseTime=true&loc=Local")
	engine.LogMode(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "unicorn_"
	}
	engine.DB().SetMaxOpenConns(3)
	engine.DB().SetConnMaxLifetime(time.Hour)
	engine.DB().SetMaxIdleConns(3)
	Engine = engine
}
