package make_result

import (
	"net/http"

	"github.com/labstack/echo"
)

// 统一的格式输出
func ResponseWithJson(c echo.Context, statusOk int, data interface{}) error {
	var response = make(map[string]interface{})
	response["code"] = statusOk
	if statusOk == http.StatusOK {
		response["data"] = data
		return c.JSON(http.StatusOK, response)
	} else {
		response["error"] = data
		return c.JSON(statusOk, response)
	}
}
