package aim

import (
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/database"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/error"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/middleware"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/web/model"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/web/result"
	"net/http"
	"strconv"

	"qiniupkg.com/x/log.v7"

	"github.com/labstack/echo"
)

func getAllAimsHandler(c echo.Context) error {
	return make_result.DefaultNilResponseWithJson(c)
}

func getOneAimHandler(c echo.Context) error {
	return make_result.DefaultNilResponseWithJson(c)
}

func createAimHandler(c echo.Context) error {
	var param createParam
	if err := c.Bind(&param); err != nil {
		bindErr := error_target_notes.ParamErrorTarget
		bindErr.Report = err.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, bindErr)
	}

	if err := make_result.DefaultErrorResponseWithJson(c, param); err != nil {
		return make_result.ResponseWithJson(c, http.StatusBadRequest, err)
	}
	log.Println("Param: ", param)
	admin := middleware.CurrentAdmin(c)
	log.Println("Admin: ", admin)
	var target model.Target
	target = model.Target{
		AdminId: admin.Id,
		Title:   param.Data.Title,
		Level:   model.NORMAL,
	}
	tx := database.Engine.NewSession()
	defer tx.Close()
	tx.Begin()
	if _, dbError := tx.InsertOne(&target); dbError != nil {
		tx.Rollback()
		dbErr := error_target_notes.InsertErrorTarget
		dbErr.Report = dbError.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)
	}

	tx.Commit()
	return make_result.ResponseWithJson(c, http.StatusOK, target.Serializer())

}

func patchAimHandler(c echo.Context) error {
	AimId := c.Param("aim_id")
	aimId, _ := strconv.Atoi(AimId)
	otherTarget := middleware.CurrentOtherTarget(c)
	if otherTarget.Id == int64(aimId) {
		return make_result.ResponseWithJson(c, http.StatusBadRequest, error_target_notes.RightErrorTarget)
	}

	var param patchParam
	if err := c.Bind(&param); err != nil {
		bindErr := error_target_notes.ParamErrorTarget
		bindErr.Report = err.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, bindErr)
	}

	if err := make_result.DefaultErrorResponseWithJson(c, param); err != nil {
		return make_result.ResponseWithJson(c, http.StatusBadRequest, err)
	}

	tx := database.Engine.NewSession()
	defer tx.Close()
	tx.Begin()
	var aim model.Target
	if _, dbError := tx.ID(AimId).Get(&aim); dbError != nil {
		dbErr := error_target_notes.RecordErrorTarget
		dbErr.Report = dbError.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)
	}

	if param.Data.Level != "" {
		aim.Level, _ = strconv.Atoi(param.Data.Level)
		if _, dbError := tx.ID(AimId).Cols("level").Update(&aim); dbError != nil {
			tx.Rollback()
			dbErr := error_target_notes.UpdateErrorTarget
			dbErr.Report = dbError.Error()
			return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)
		}
	}

	if param.Data.Description != "" {
		aim.Description = param.Data.Description
		if _, dbError := tx.ID(AimId).Cols("description").Update(&aim); dbError != nil {
			tx.Rollback()
			dbErr := error_target_notes.UpdateErrorTarget
			dbErr.Report = dbError.Error()
			return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)
		}
	}

	if param.Data.Status != "" {
		aim.Status, _ = strconv.Atoi(param.Data.Status)
		if _, dbError := tx.ID(AimId).Cols("status").Update(&aim); dbError != nil {
			tx.Rollback()
			dbErr := error_target_notes.UpdateErrorTarget
			dbErr.Report = dbError.Error()
			return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)
		}
	}
	tx.Commit()
	return make_result.ResponseWithJson(c, http.StatusOK, aim.Serializer())
}

func deleteAimHandler(c echo.Context) error {
	return make_result.DefaultNilResponseWithJson(c)
}
