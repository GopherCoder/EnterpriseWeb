package database

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/go-xorm/xorm"
	"xorm.io/core"

	_ "github.com/go-sql-driver/mysql"
)

var Engine *xorm.Engine

var (
	db       = "lottery"
	user     = "root"
	password = "admin123"
)

func MySQLInit() {
	dataSourceName := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=true&loc=Local", user, password, db)
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		log.Fatalln("CONNECT DATABASE FAIL :", err.Error())
		return
	}
	//logger := log_for_lottery.SimpleLogger{}
	//engine.SetLogger(&logger)
	Engine = engine

	l := CustomerLogger(os.Stdout)
	l.SetLevel(core.LOG_INFO)
	Engine.SetLogger(l)
	Engine.ShowSQL(true)
	Engine.SetMapper(core.GonicMapper{})
	Engine.SetTableMapper(core.GonicMapper{})
}

func CustomerLogger(out io.Writer) *xorm.SimpleLogger {
	prefix := xorm.DEFAULT_LOG_PREFIX
	flag := xorm.DEFAULT_LOG_FLAG
	return &xorm.SimpleLogger{
		DEBUG: log.New(out, White(fmt.Sprintf("%s [debug] ", prefix)), flag),
		ERR:   log.New(out, Red(fmt.Sprintf("%s [error] ", prefix)), flag),
		INFO:  log.New(out, Blue(fmt.Sprintf("%s [info] ", prefix)), flag),
		WARN:  log.New(out, Red(fmt.Sprintf("%s [warn]  ", prefix)), flag),
	}
}
func Bold(message string) string {
	return fmt.Sprintf("\x1b[1m%s\x1b[21m", message)
}

// Black returns a black string
func Black(message string) string {
	return fmt.Sprintf("\x1b[30m%s\x1b[0m", message)
}

// White returns a white string
func White(message string) string {
	return fmt.Sprintf("\x1b[37m%s\x1b[0m", message)
}

// Cyan returns a cyan string
func Cyan(message string) string {
	return fmt.Sprintf("\x1b[36m%s\x1b[0m", message)
}

// Blue returns a blue string
func Blue(message string) string {
	return fmt.Sprintf("\x1b[34m%s\x1b[0m", message)
}

// Red returns a red string
func Red(message string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", message)
}

// Green returns a green string
func Green(message string) string {
	return fmt.Sprintf("\x1b[32m%s\x1b[0m", message)
}

// Yellow returns a yellow string
func Yellow(message string) string {
	return fmt.Sprintf("\x1b[33m%s\x1b[0m", message)
}

// Gray returns a gray string
func Gray(message string) string {
	return fmt.Sprintf("\x1b[37m%s\x1b[0m", message)
}

// Magenta returns a magenta string
func Magenta(message string) string {
	return fmt.Sprintf("\x1b[35m%s\x1b[0m", message)
}
