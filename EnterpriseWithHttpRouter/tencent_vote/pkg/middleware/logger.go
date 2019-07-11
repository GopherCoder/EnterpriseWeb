package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		Message := fmt.Sprintf("HOST: %s | URL: %s | TIMES: %s | METHOD: %s", request.Host, request.RequestURI, time.Now().Format("2006-01-02 15:04:05"), request.Method)
		log.Printf(Red(Message))
		next.ServeHTTP(writer, request)
	}
}

func Red(message string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", message)
}
