package database

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Engine *gorm.DB
var err error

func EngineInit() {
	engine, err := gorm.Open("mysql", "root:admin123@/unicorn?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		log.Panic("CONNECT MYSQL FAIL ", err)
		return
	}
	engine.LogMode(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "unicorn_"
	}
	gorm.DefaultTableNameHandler(engine, "unicorn_")
	engine.DB().SetMaxOpenConns(3)
	engine.DB().SetConnMaxLifetime(time.Hour)
	engine.DB().SetMaxIdleConns(3)
	Engine = engine
}
