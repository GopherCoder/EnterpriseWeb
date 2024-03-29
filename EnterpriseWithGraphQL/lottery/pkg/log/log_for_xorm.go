package log_for_lottery

import (
	"fmt"
	"log"

	"xorm.io/core"
)

type SimpleLogger struct {
	DEBUG   *log.Logger
	ERR     *log.Logger
	INFO    *log.Logger
	WARN    *log.Logger
	level   core.LogLevel
	showSQL bool
}

// Error implement core.ILogger
func (s *SimpleLogger) Error(v ...interface{}) {
	if s.level <= core.LOG_ERR {
		s.ERR.Output(2, Red(fmt.Sprint(v...)))
	}
	return
}

// Errorf implement core.ILogger
func (s *SimpleLogger) Errorf(format string, v ...interface{}) {
	if s.level <= core.LOG_ERR {
		s.ERR.Output(2, Red(fmt.Sprintf(format, v...)))
	}
	return
}

// Debug implement core.ILogger
func (s *SimpleLogger) Debug(v ...interface{}) {
	if s.level <= core.LOG_DEBUG {
		s.DEBUG.Output(2, Red(fmt.Sprint(v...)))
	}
	return
}

// Debugf implement core.ILogger
func (s *SimpleLogger) Debugf(format string, v ...interface{}) {
	if s.level <= core.LOG_DEBUG {
		s.DEBUG.Output(2, Red(fmt.Sprintf(format, v...)))
	}
	return
}

// Info implement core.ILogger
func (s *SimpleLogger) Info(v ...interface{}) {
	if s.level <= core.LOG_INFO {
		s.INFO.Output(2, Red(fmt.Sprint(v...)))
	}
	return
}

// Infof implement core.ILogger
func (s *SimpleLogger) Infof(format string, v ...interface{}) {
	if s.level <= core.LOG_INFO {
		s.INFO.Output(2, Red(fmt.Sprintf(format, v...)))
	}
	return
}

// Warn implement core.ILogger
func (s *SimpleLogger) Warn(v ...interface{}) {
	if s.level <= core.LOG_WARNING {
		s.WARN.Output(2, Red(fmt.Sprint(v...)))
	}
	return
}

// Warnf implement core.ILogger
func (s *SimpleLogger) Warnf(format string, v ...interface{}) {
	if s.level <= core.LOG_WARNING {
		s.WARN.Output(2, Red(fmt.Sprintf(format, v...)))
	}
	return
}

// Level implement core.ILogger
func (s *SimpleLogger) Level() core.LogLevel {
	return s.level
}

// SetLevel implement core.ILogger
func (s *SimpleLogger) SetLevel(l core.LogLevel) {
	s.level = l
	return
}

// ShowSQL implement core.ILogger
func (s *SimpleLogger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		s.showSQL = true
		return
	}
	s.showSQL = show[0]
}

// IsShowSQL implement core.ILogger
func (s *SimpleLogger) IsShowSQL() bool {
	return s.showSQL
}
