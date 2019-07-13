package middleware

import (
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/pkg/database"
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/web/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		Message := fmt.Sprintf("[HTTP_ROUTER_LOGGER]: HOST: %s | URL: %s | TIMES: %s | METHOD: %s", request.Host, request.RequestURI, time.Now().Format("2006-01-02 15:04:05"), request.Method)
		log.Printf(Red(Message))
		next.ServeHTTP(writer, request)
	}
}

func Red(message string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", message)
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		authToken := request.Header.Get("Authorization")
		list := strings.Split(authToken, " ")
		if len(list) != 2 {
			json.NewEncoder(writer).Encode(fmt.Sprintf("Please Add Authorization"))
			return
		}
		bearer := list[1]
		var admin model.Admin
		if dbError := database.Engine.Where("token = ?", bearer).First(&admin).Error; dbError != nil {
			json.NewEncoder(writer).Encode(dbError)
			return
		}
		request.Header.Add("current_admin_id", strconv.Itoa(int(admin.ID)))
		next.ServeHTTP(writer, request)
	}
}
