package log_target_notes

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const PROJECT = "./EnterpriseWeb/EnterpriseWithEcho/target_notes/log/target_notes"

var logger *log.Logger
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
	logger = log.New(LoggerFile, projectName, log.LstdFlags)
	logger.Println("Start log.")
}
