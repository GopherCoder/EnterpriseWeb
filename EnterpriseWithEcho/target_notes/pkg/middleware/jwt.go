package middleware

import (
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/database"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/web/model"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		token := context.Request().Header.Get("Authorization")
		tokenList := strings.Split(token, " ")
		if len(tokenList) != 2 {
			var results = make(map[string]interface{})
			results["code"] = http.StatusBadRequest
			results["data"] = "Token: Bearer xx"
			return context.JSON(http.StatusBadRequest, results)
		}
		var admin model.Admin
		database.Engine.Where("token = ?", tokenList[1]).Get(&admin)
		context.Set("current_admin", admin)
		context.Set("current_admin_id", admin.Id)
		context.Set("current_account_name", admin.AccountName)
		return next(context)
	}
}
