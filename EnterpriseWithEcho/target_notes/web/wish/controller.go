package wish

import (
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/database"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/error"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/log"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/web/make"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/web/model"
	"net/http"

	"github.com/labstack/echo"
)

func getWish(c echo.Context) error {
	return make_result.ResponseWithJson(c, http.StatusOK, nil)
}

func getAllWish(c echo.Context) error {
	return make_result.ResponseWithJson(c, http.StatusOK, nil)

}

func postWish(c echo.Context) error {

	var param postWishPara

	if err := c.Bind(&param); err != nil {
		var paramErr error_target_notes.ErrorTarget
		paramErr = error_target_notes.ParamErrorTarget
		paramErr.Report = err.Error()
		log_target_notes.Logger.Fatalf("bind param fail")
		return paramErr
	}

	if validErr := param.Valid(); validErr != nil {
		var validErr error_target_notes.ErrorTarget
		validErr = error_target_notes.ValidErrorTarget
		validErr.Report = validErr.Error()
		log_target_notes.Logger.Fatalf("param invalid")
		return validErr
	}

	var wish model.Wish
	wish = model.Wish{
		Title: param.Title,
	}

	if _, dbError := database.Engine.InsertOne(&wish); dbError != nil {
		var err error_target_notes.ErrorTarget
		err = error_target_notes.InsertErrorTarget
		err.Report = dbError.Error()
		return err
	}

	return make_result.ResponseWithJson(c, http.StatusOK, wish.Serializer())

}

func deleteWish(c echo.Context) error {
	return make_result.ResponseWithJson(c, http.StatusOK, nil)

}

func patchWish(c echo.Context) error {
	return make_result.ResponseWithJson(c, http.StatusOK, nil)

}
