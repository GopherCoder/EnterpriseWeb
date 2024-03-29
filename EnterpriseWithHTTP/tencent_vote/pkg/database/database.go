package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	PROJECT = "tencentVotes_"
)

var Engine *gorm.DB

func EngineInit() {
	db, err := gorm.Open("mysql", "root:admin123@/tencent_votes?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		log.Panic("CONNECT DATABASE ERROR")
	}
	db.LogMode(true)
	Engine = db
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return PROJECT + defaultTableName
	}

}
