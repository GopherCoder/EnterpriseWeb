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
		var target model.Target
		database.Engine.Where("admin_id = ? AND title = ?", admin.Id, "其他").Get(&target)
		context.Set("current_admin", admin)
		context.Set("current_admin_id", admin.Id)
		context.Set("current_account_name", admin.AccountName)
		context.Set("current_other_target_id", target.Id)
		context.Set("current_other_target", target)
		return next(context)
	}
}

func CurrentOtherTarget(c echo.Context) model.Target {
	return c.Get("current_other_target").(model.Target)
}
func CurrentOtherTargetId(c echo.Context) int64 {
	return c.Get("current_other_target_id").(int64)
}

func CurrentAdmin(c echo.Context) model.Admin {
	return c.Get("current_admin").(model.Admin)
}

func CurrentAdminAccount(c echo.Context) string {
	return c.Get("current_account_name").(string)
}

func CurrentAdminId(c echo.Context) int64 {
	return c.Get("current_admin_id").(int64)
}
