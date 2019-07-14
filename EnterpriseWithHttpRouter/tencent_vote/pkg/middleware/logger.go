package middleware

import (
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/pkg/database"
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/web/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

var Store map[string]interface{}

func init() {
	Store = make(map[string]interface{})
}

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
		var results = make(map[string]interface{})
		authToken := request.Header.Get("Authorization")
		list := strings.Split(authToken, " ")
		if len(list) != 2 {
			writer.Header().Set("Content-type", "application/json")
			results["code"] = http.StatusBadRequest
			results["error"] = "Please Add Authorization"
			json.NewEncoder(writer).Encode(results)
			return
		}
		bearer := list[1]
		var admin model.Admin
		if dbError := database.Engine.Where("token = ?", bearer).First(&admin).Error; dbError != nil {
			writer.Header().Set("Content-type", "application/json")
			results["code"] = http.StatusBadRequest
			results["error"] = dbError.Error()
			json.NewEncoder(writer).Encode(results)
			return
		}
		Store["current_admin_id"] = admin.ID
		Store["current_admin"] = admin
		next.ServeHTTP(writer, request)
	}
}

func CurrentAdmin() model.Admin {
	return Store["current_admin"].(model.Admin)
}

func CurrentAdminId() uint {
	return Store["current_admin_id"].(uint)
}
