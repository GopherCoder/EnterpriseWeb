package admin

import (
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/database"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/error"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/web/model"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/web/result"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func registerHandler(c echo.Context) error {

	var param registerParam
	if err := c.Bind(&param); err != nil {
		paramErr := error_target_notes.ParamErrorTarget
		paramErr.Report = err.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, paramErr)
	}
	log.Print("Param ", param)
	if err := param.Valid(); err != nil {
		validErr := error_target_notes.ValidErrorTarget
		validErr.Report = err.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, validErr)
	}

	log.Println("valid", param)

	password, _ := generateFromPassword(param.Password)

	log.Println("Password", string(password))
	var admin model.Admin
	admin = model.Admin{
		AccountName: param.AccountName,
		Password:    string(password),
		Token:       generateToken(20),
	}
	tx := database.Engine.NewSession()
	defer tx.Close()
	tx.Begin()
	if _, dbError := tx.InsertOne(&admin); dbError != nil {
		tx.Rollback()
		dbErr := error_target_notes.InsertErrorTarget
		dbErr.Report = dbError.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, dbError)
	}
	log.Println("Insert", "Insert successful")
	tx.Commit()
	return make_result.ResponseWithJson(c, http.StatusOK, admin.Serializer())
}

func loginHandler(c echo.Context) error {
	var param registerParam
	if err := c.Bind(&param); err != nil {
		bindError := error_target_notes.ParamErrorTarget
		bindError.Report = err.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, bindError)
	}
	if err := param.Valid(); err != nil {
		invalidError := error_target_notes.ValidErrorTarget
		invalidError.Report = err.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, invalidError)
	}

	var admin model.Admin
	if _, dbError := database.Engine.Where("account_name = ?", param.AccountName).Get(&admin); dbError != nil {
		dbErr := error_target_notes.RecordErrorTarget
		dbErr.Report = dbError.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)
	}

	if ok := compareHashAndPassword([]byte(admin.Password), []byte(param.Password)); !ok {
		passwordErr := error_target_notes.PassWordErrorTarget
		passwordErr.Report = fmt.Sprintf("password is not correct")
		return make_result.ResponseWithJson(c, http.StatusBadRequest, passwordErr)
	}

	return make_result.ResponseWithJson(c, http.StatusOK, admin.Serializer())
}

func logoutHandler(c echo.Context) error {
	var param registerParam
	if err := c.Bind(&param); err != nil {
		bindError := error_target_notes.ParamErrorTarget
		bindError.Report = err.Error()
		return bindError
	}
	if err := param.Valid(); err != nil {
		invalidError := error_target_notes.ValidErrorTarget
		invalidError.Report = err.Error()
		return invalidError
	}

	var admin model.Admin
	if _, dbError := database.Engine.Where("account_name = ?", param.AccountName).Get(&admin); dbError != nil {
		dbErr := error_target_notes.RecordErrorTarget
		dbErr.Report = dbError.Error()
		return dbError
	}

	return c.Redirect(http.StatusOK, "/v1/api/login")
}
