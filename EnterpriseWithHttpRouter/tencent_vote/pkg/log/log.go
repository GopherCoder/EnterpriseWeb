package log_tencent_votes

import (
	"log"
	"os"
)

var LoggerSysOut *log.Logger
var LoggerFile *log.Logger

func LoggerInit() {
	LoggerSysOut = log.New(os.Stdout, "SysOut Logger: ", log.LstdFlags)
	logFile, _ := os.Open("../log/log.log")
	LoggerFile = log.New(logFile, "File Logger: ", log.LstdFlags)
}
