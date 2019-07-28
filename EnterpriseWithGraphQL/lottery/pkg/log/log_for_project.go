package log_for_lottery

import "log"

func Println(message ...interface{}) {
	log.Println(RedM(message...))
}

func Print(message ...interface{}) {
	log.Print(RedM(message...))
}

func Fatal(message ...interface{}) {
	log.Fatal(RedM(message...))
}

func Fatalf(format string, message ...interface{}) {
	log.Fatalf(RedM(format, message))
}
func Fatalln(message ...interface{}) {
	log.Fatalln(RedM(message...))
}
