package aim

import (
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/database"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/error"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/middleware"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/web/model"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/web/result"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func getAllAimsHandler(c echo.Context) error {

	adminId := middleware.CurrentAdminId(c)

	var aims []model.Target

	if dbError := database.Engine.Alias("t").Where("t.admin_id = ? AND t.title <> '其他'", adminId).Find(&aims); dbError != nil {
		dbErr := error_target_notes.RecordErrorTarget
		dbErr.Report = dbError.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)
	}
	var results []model.TargetSerializerWithTaskTitle
	for _, i := range aims {
		results = append(results, i.SerializerWithTaskTitle())
	}
	return make_result.ResponseWithJson(c, http.StatusOK, results)
}

func getOneAimHandler(c echo.Context) error {
	aimId := c.Param("aim_id")
	var aim model.Target
	if _, dbError := database.Engine.ID(aimId).Get(&aim); dbError != nil {
		dbErr := error_target_notes.RecordErrorTarget
		dbErr.Report = dbError.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)
	}

	return make_result.ResponseWithJson(c, http.StatusOK, aim.SerializerWithTaskTitle())
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
	aimId := c.Param("aim_id")
	var aim model.Target
	tx := database.Engine.NewSession()
	defer tx.Close()
	tx.Begin()
	if _, dbError := tx.ID(aimId).Get(&aim); dbError != nil {
		dbErr := error_target_notes.RecordErrorTarget
		dbErr.Report = dbError.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)
	}
	if len(aim.TaskIds) != 0 {
		var things []model.Things
		if dbError := tx.In("id", aim.TaskIds).Find(&things); dbError != nil {
			dbErr := error_target_notes.RecordErrorTarget
			dbErr.Report = dbError.Error()
			return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)
		}
		if len(things) != 0 {
			for _, i := range things {
				if _, dbError := tx.ID(i.Id).Delete(&i); dbError != nil {
					tx.Rollback()
					dbErr := error_target_notes.DeleteErrorTarget
					dbErr.Report = dbError.Error()
					return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)
				}
			}
		}

	}

	if _, dbError := tx.ID(aimId).Delete(&aim); dbError != nil {
		dbErr := error_target_notes.DeleteErrorTarget
		dbErr.Report = dbError.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)
	}
	tx.Commit()
	return make_result.ResponseWithJson(c, http.StatusOK, aim.SerializerWithTaskTitle())
}

func patchAimThingsHandler(c echo.Context) error {
	aimId := c.Param("aim_id")
	taskId := c.Param("task_id")
	otherTarget := middleware.CurrentOtherTarget(c)

	var param pathThingsParam
	if err := c.Bind(&param); err != nil {
		bindErr := error_target_notes.ParamErrorTarget
		bindErr.Report = err.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, bindErr)
	}

	if len(param.Data) == 0 {
		return make_result.ResponseWithJson(c, http.StatusBadRequest, fmt.Errorf("param should not be nil"))

	}

	tx := database.Engine.NewSession()
	defer tx.Close()
	tx.Begin()

	var aim model.Target
	if _, dbError := tx.ID(aimId).Get(&aim); dbError != nil && aim.Id != otherTarget.Id {
		dbErr := error_target_notes.RecordErrorTarget
		dbErr.Report = dbError.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)
	}

	var task model.Task
	if _, dbError := tx.ID(taskId).Get(&task); dbError != nil {
		dbErr := error_target_notes.RecordErrorTarget
		dbErr.Report = dbError.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)
	}

	for _, i := range param.Data {
		if i.Id == 0 {
			// create
			var tempThings model.Things
			tempThings.Description = i.Description
			if _, dbError := tx.InsertOne(&tempThings); dbError != nil {
				tx.Rollback()
				dbErr := error_target_notes.RecordErrorTarget
				dbErr.Report = dbError.Error()
				return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)
			}
			task.ThingIds = append(task.ThingIds, tempThings.Id)
			if _, dbError := tx.ID(taskId).Cols("thing_ids").Update(&task); dbError != nil {
				tx.Rollback()
				dbErr := error_target_notes.UpdateErrorTarget
				dbErr.Report = dbError.Error()
				return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)
			}
		} else {
			// update
			var tempThings model.Things
			if _, dbError := tx.ID(i.Id).Get(&tempThings); dbError != nil {
				dbErr := error_target_notes.RecordErrorTarget
				dbErr.Report = dbError.Error()
				return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)
			}
			tempThings.Description = i.Description
			if _, dbError := tx.ID(i.Id).Cols("description").Update(tempThings); dbError != nil {
				dbErr := error_target_notes.UpdateErrorTarget
				dbErr.Report = dbError.Error()
				return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)
			}
		}
	}
	tx.Commit()
	return make_result.ResponseWithJson(c, http.StatusOK, aim.SerializerWithTaskTitle())

}

func createTaskHandler(c echo.Context) error {

	aimId := c.Param("aim_id")
	otherTarget := middleware.CurrentOtherTarget(c)
	var param createTaskParam
	if err := c.Bind(&param); err != nil {
		bindErr := error_target_notes.ParamErrorTarget
		bindErr.Report = err.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, bindErr)
	}

	log.Println("Param: ", param)
	tx := database.Engine.NewSession()
	defer tx.Close()
	tx.Begin()

	var aim model.Target
	if _, dbError := tx.ID(aimId).Get(&aim); dbError != nil && aim.Id != otherTarget.Id {
		dbErr := error_target_notes.RecordErrorTarget
		dbErr.Report = dbError.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)
	}
	if aim.Id == 0 {
		return make_result.ResponseWithJson(c, http.StatusBadRequest, fmt.Sprintf("record not found"))
	}
	log.Println("aim: ", aim.Id, aim.Title)
	var task model.Task
	task = model.Task{
		Title:    param.Title,
		TargetId: aim.Id,
		Status:   model.UNDONE,
	}
	if _, dbError := tx.InsertOne(&task); dbError != nil {
		tx.Rollback()
		dbErr := error_target_notes.InsertErrorTarget
		dbErr.Report = dbError.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)

	}

	log.Println("task: ", task.Id, task.Title)

	aim.TaskIds = append(aim.TaskIds, task.Id)
	tx.ID(aim.Id).Cols("task_ids").Update(&aim)
	tx.Commit()
	return make_result.ResponseWithJson(c, http.StatusOK, aim.SerializerWithTaskTitle())

}
