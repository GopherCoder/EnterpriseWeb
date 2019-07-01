package admin

import (
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/error"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/web/model"

	"github.com/labstack/echo"
)

func registerHandler(c echo.Context) error {

	var param registerParam
	if err := c.Bind(&param); err != nil {
		paramErr := error_target_notes.ParamErrorTarget
		paramErr.Report = err.Error()
		return paramErr
	}

	if err := param.Valid(); err != nil {
		validErr := error_target_notes.ValidErrorTarget
		validErr.Report = err.Error()
		return validErr
	}

	password, _ := generateFromPassword(param.Password, 20)

	var admin model.Admin
	admin = model.Admin{
		AccountName: param.AccountName,
		Password:    string(password),
		Token:       generateToken(20),
	}

}
