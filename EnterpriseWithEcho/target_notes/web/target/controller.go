package target

import (
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/database"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/error"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/middleware"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/web/model"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/web/result"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/labstack/echo"
)

func getOtherTargetHandler(c echo.Context) error {
	otherTarget := middleware.CurrentOtherTarget(c)

	log.Println("taskIds: ", otherTarget.TaskIds)
	return make_result.ResponseWithJson(c, http.StatusOK, otherTarget.SerializerWithTaskTitle())
}

func getOneTaskHandler(c echo.Context) error {
	taskId := c.Param("task_id")
	var task model.Task
	if _, dbError := database.Engine.ID(taskId).Get(&task); dbError != nil {
		dbErr := error_target_notes.RecordErrorTarget
		dbErr.Report = dbError.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)
	}
	return make_result.ResponseWithJson(c, http.StatusOK, task.Serializer())
}

func createOneTaskHandler(c echo.Context) error {

	var param createTaskParam
	var err error

	if err := c.Bind(&param); err != nil {
		bindErr := error_target_notes.ParamErrorTarget
		bindErr.Report = err.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, bindErr)
	}

	if err = make_result.DefaultErrorResponseWithJson(c, param); err != nil {
		log.Print("err: ", err)
		return err
	}

	log.Println("Param", param)

	otherTargetId := middleware.CurrentOtherTargetId(c)
	tx := database.Engine.NewSession()
	defer tx.Close()
	tx.Begin()

	var task model.Task
	task = model.Task{
		TargetId: otherTargetId,
		Title:    param.Title,
	}

	if _, dbError := tx.InsertOne(&task); dbError != nil {
		tx.Rollback()
		return make_result.DefaultErrorDataBaseWithJson(c, error_target_notes.InsertErrorTarget, dbError)
	}
	otherTarget := middleware.CurrentOtherTarget(c)
	otherTarget.TaskIds = append(otherTarget.TaskIds, task.Id)

	if _, dbError := tx.ID(otherTargetId).Cols("task_ids").Update(&otherTarget); dbError != nil {
		tx.Rollback()
		return make_result.DefaultErrorDataBaseWithJson(c, error_target_notes.InsertErrorTarget, dbError)
	}
	tx.Commit()
	return make_result.ResponseWithJson(c, http.StatusOK, otherTarget.Serializer())
}

func patchOneTaskHandler(c echo.Context) error {
	var param patchTaskParam
	if err := c.Bind(&param); err != nil {
		paramErr := error_target_notes.ParamErrorTarget
		paramErr.Report = err.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, paramErr)
	}
	if err := make_result.DefaultErrorResponseWithJson(c, param); err != nil {
		return err
	}

	log.Println("Param: ", param)

	tx := database.Engine.NewSession()
	defer tx.Close()
	tx.Begin()

	taskId := c.Param("task_id")

	var task model.Task
	if _, dbError := tx.ID(taskId).Get(&task); dbError != nil {
		return make_result.DefaultErrorDataBaseWithJson(c, error_target_notes.RecordErrorTarget, dbError)
	}
	if param.Data.TargetId != 0 {
		if ok, dbError := tx.ID(param.Data.TargetId).Get(&model.Target{}); !ok && dbError != nil {
			return make_result.DefaultErrorDataBaseWithJson(c, error_target_notes.RecordErrorTarget, dbError)
		}
		task.TargetId = int64(param.Data.TargetId)
		if _, dbError := tx.ID(taskId).Cols("target_id").Update(&task); dbError != nil {
			tx.Rollback()
			return make_result.DefaultErrorDataBaseWithJson(c, error_target_notes.InsertErrorTarget, dbError)
		}
	}

	if param.Data.Title != "" {
		task.Title = param.Data.Title
	}
	if param.Data.Description != "" {
		task.Description = param.Data.Description
	}

	task.Status = param.Data.Status
	if _, dbError := tx.ID(taskId).Cols("title", "description", "status").Update(&task); dbError != nil {
		tx.Rollback()
		return make_result.DefaultErrorDataBaseWithJson(c, error_target_notes.InsertErrorTarget, dbError)
	}
	tx.Commit()
	return make_result.ResponseWithJson(c, http.StatusOK, task.Serializer())
}

func deleteOneTaskHandler(c echo.Context) error {

	otherTarget := middleware.CurrentOtherTarget(c)
	taskId := c.Param("task_id")

	var ok bool
	for _, i := range otherTarget.TaskIds {
		if strconv.Itoa(int(i)) == taskId {
			ok = true
		}
	}
	if !ok {
		return make_result.ResponseWithJson(c, http.StatusBadRequest, fmt.Sprintf("task id not in target"))
	}
	var task model.Task
	tx := database.Engine.NewSession()
	defer tx.Close()
	tx.Begin()
	if _, dbError := tx.ID(taskId).Get(&task); dbError != nil {
		return make_result.DefaultErrorDataBaseWithJson(c, error_target_notes.RecordErrorTarget, dbError)
	}

	if _, dbError := tx.ID(taskId).Delete(&task); dbError != nil {
		tx.Rollback()
		return make_result.DefaultErrorDataBaseWithJson(c, error_target_notes.DeleteErrorTarget, dbError)
	}

	if len(task.ThingIds) != 0 {
		for _, i := range task.ThingIds {
			tx.ID(i).Delete(new(model.Things))
		}
	}

	tx.Commit()
	return make_result.ResponseWithJson(c, http.StatusOK, task.Serializer())
}

func orderTaskHandler(c echo.Context) error {
	var params patchOrderParam

	if err := c.Bind(&params); err != nil {
		paramErr := error_target_notes.ParamErrorTarget
		paramErr.Report = err.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, paramErr)
	}

	if err := make_result.DefaultErrorResponseWithJson(c, params); err != nil {
		return make_result.ResponseWithJson(c, http.StatusBadRequest, err)
	}

	otherTarget := middleware.CurrentOtherTarget(c)
	taskIds := otherTarget.TaskIds
	var taskIdsInt []int
	for _, i := range taskIds {
		taskIdsInt = append(taskIdsInt, int(i))
	}

	sort.Ints(taskIdsInt)

	tx := database.Engine.NewSession()
	defer tx.Close()
	tx.Begin()

	for index, i := range params.Data.TaskIds {
		if _, dbError := tx.Table(new(model.Task)).ID(i).Update(map[string]interface{}{"order_level": taskIdsInt[index]}); dbError != nil {
			tx.Rollback()
			dbErr := error_target_notes.UpdateErrorTarget
			dbErr.Report = dbError.Error()
			return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)
		}
	}
	var tasks []model.Task
	tx.In("id", params.Data.TaskIds).OrderBy("order_level").Desc("order_level").Find(&tasks)
	tx.Commit()

	var result []model.TaskSerializer
	for _, i := range tasks {
		result = append(result, i.Serializer())
	}
	return make_result.ResponseWithJson(c, http.StatusOK, result)
}

func patchOneThingHandler(c echo.Context) error {
	taskId := c.Param("task_id")

	var param patchThingParam
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

	var task model.Task
	if _, dbError := tx.ID(taskId).Get(&task); dbError != nil {
		dbErr := error_target_notes.RecordErrorTarget
		dbErr.Report = dbError.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)
	}

	for _, i := range param.Data {
		if i.Id == 0 {
			// create : 插入 task
			var thing model.Things
			thing.Description = i.Description
			if _, dbError := tx.InsertOne(&thing); dbError != nil {
				tx.Rollback()
				dbErr := error_target_notes.InsertErrorTarget
				dbErr.Report = dbError.Error()
				return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)
			}
			task.ThingIds = append(task.ThingIds, thing.Id)
			if _, dbError := tx.ID(taskId).Cols("thing_ids").Update(&task); dbError != nil {
				tx.Rollback()
				dbErr := error_target_notes.UpdateErrorTarget
				dbErr.Report = dbError.Error()
				return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)

			}
		} else {
			// update:
			thing := new(model.Things)
			thing.Description = i.Description
			if _, dbError := tx.ID(i.Id).Cols("description").Update(&thing); dbError != nil {
				tx.Rollback()
				dbErr := error_target_notes.UpdateErrorTarget
				dbErr.Report = dbError.Error()
				return make_result.ResponseWithJson(c, http.StatusBadRequest, dbErr)
			}
		}
	}
	tx.Commit()
	return make_result.ResponseWithJson(c, http.StatusOK, task.Serializer())

}
