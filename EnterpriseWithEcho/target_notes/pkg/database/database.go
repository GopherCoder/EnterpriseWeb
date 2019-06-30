package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

var Engine *xorm.Engine

func EngineInit() {
	engine, err := xorm.NewEngine("mysql", "root:admin123@/targetNotes?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic("CAN NOT OPEN DATABASE")
	}
	Engine = engine
	Engine.ShowSQL(true)
	Engine.Logger().SetLevel(core.LOG_DEBUG)
	Engine.Logger()
	Engine.SetTableMapper(core.GonicMapper{})
	Engine.SetColumnMapper(core.GonicMapper{})
}
