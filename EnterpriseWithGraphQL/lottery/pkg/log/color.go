package log_for_lottery

import "fmt"

func Red(message string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", message)
}

func RedM(message ...interface{}) string {
	return fmt.Sprintf("\x1b[31m%v\x1b[0m", message...)
}
