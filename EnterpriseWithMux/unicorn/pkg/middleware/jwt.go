package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func LoggerMiddle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		format := fmt.Sprintf("%s | %s | %s | %s", r.Method, r.RequestURI, r.Host, time.Now().Format("2006-01-02 15:04:05"))
		log.Println(format)
		next.ServeHTTP(w, r)
	})
}
