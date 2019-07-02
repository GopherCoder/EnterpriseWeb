package log_target_notes

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var GOPATH = os.Getenv("GOPATH")
var PROJECT = GOPATH + "/src" + "/EnterpriseWeb/EnterpriseWithEcho/target_notes/log/target_notes"

var Logger *log.Logger
var loggerName string
var err error
var LoggerFile *os.File

func LOGInit() {
	loggerName = fmt.Sprintf("%s.log", fmt.Sprintf(PROJECT))
	LoggerFile, err = os.OpenFile(loggerName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		panic("Can Not Open Log File In Dir Of Log.")
	}
	projectName := strings.Split(PROJECT, "/")[3] + " "
	Logger = log.New(LoggerFile, projectName, log.LstdFlags)
	Logger.Println("Start log.")
}
