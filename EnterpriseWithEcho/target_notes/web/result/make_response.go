package make_result

import (
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/error"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/web/param"
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

func DefaultNilResponseWithJson(c echo.Context) error {
	return ResponseWithJson(c, http.StatusOK, nil)
}

func DefaultErrorResponseWithJson(c echo.Context, form param.ValidWithParam) error {
	if invalidError := form.Valid(); invalidError != nil {
		validError := error_target_notes.ValidErrorTarget
		validError.Report = invalidError.Error()
		return ResponseWithJson(c, http.StatusBadRequest, validError)
	}
	return nil
}

func DefaultErrorDataBaseWithJson(c echo.Context, err error_target_notes.ErrorTarget, report error) error {
	err.Report = report.Error()
	return ResponseWithJson(c, http.StatusBadRequest, err)
}
