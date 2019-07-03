package wish

import (
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/database"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/error"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/log"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/web/model"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/web/result"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func getWish(c echo.Context) error {
	wishId := c.Param("wish_id")
	adminId := c.Get("current_admin_id")
	var wish model.Wish
	if ok, dbErr := database.Engine.Id(wishId).Where("admin_id = ?", adminId).Get(&wish); !ok && dbErr != nil {
		err := error_target_notes.RecordErrorTarget
		err.Report = dbErr.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, err)
	}

	return make_result.ResponseWithJson(c, http.StatusOK, wish.Serializer())
}

func getAllWish(c echo.Context) error {
	adminId := c.Get("current_admin_id")
	log.Print("AdminId ", adminId)
	var (
		wishes []model.Wish
	)

	returnAll := c.QueryParam("return")
	if returnAll == "" {
		returnAll = "all_list"
	}
	search := c.QueryParam("search")

	query := database.Engine.Where("admin_id = ?", adminId)
	if search != "" {
		query = query.Where("title like ? or hope like ?", fmt.Sprintf("%%%s%%", search), fmt.Sprintf("%%%s%%", search))
	}

	if _, dbError := query.FindAndCount(&wishes); dbError != nil {
		err := error_target_notes.RecordErrorTarget
		err.Report = dbError.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, err)
	}
	if returnAll == "all_count" {
		var results = make(map[string]interface{})
		results["count"] = len(wishes)
		return make_result.ResponseWithJson(c, http.StatusOK, results)
	}
	if returnAll == "all_list" {
		var results model.WishSerializers
		for _, i := range wishes {
			results = append(results, i.Serializer())
		}
		return make_result.ResponseWithJson(c, http.StatusOK, results)
	}

	return make_result.ResponseWithJson(c, http.StatusOK, nil)

}

func postWish(c echo.Context) error {

	var param postWishParam

	adminId := c.Get("current_admin_id")

	if err := c.Bind(&param); err != nil {
		var paramErr error_target_notes.ErrorTarget
		paramErr = error_target_notes.ParamErrorTarget
		paramErr.Report = err.Error()
		log_target_notes.Logger.Fatalf("bind param fail")
		return make_result.ResponseWithJson(c, http.StatusBadRequest, paramErr)
	}

	if invalidErr := param.Valid(); invalidErr != nil {
		var validErr error_target_notes.ErrorTarget
		validErr = error_target_notes.ValidErrorTarget
		validErr.Report = validErr.Error()
		log_target_notes.Logger.Fatalf("param invalid")
		return make_result.ResponseWithJson(c, http.StatusBadRequest, validErr)
	}

	var wish model.Wish
	wish = model.Wish{
		Title:          param.Title,
		AdminId:        adminId.(int64),
		DesireLevel:    -1,
		ChallengeLevel: -1,
		TimeLevel:      -1,
	}

	if _, dbError := database.Engine.InsertOne(&wish); dbError != nil {
		var err error_target_notes.ErrorTarget
		err = error_target_notes.InsertErrorTarget
		err.Report = dbError.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, err)
	}

	return make_result.ResponseWithJson(c, http.StatusOK, wish.Serializer())

}

func deleteWish(c echo.Context) error {
	wishId := c.Param("wish_id")

	tx := database.Engine.NewSession()
	defer tx.Close()
	tx.Begin()
	var wish model.Wish
	if _, dbError := tx.ID(wishId).Get(&wish); dbError != nil {
		err := error_target_notes.RecordErrorTarget
		err.Report = dbError.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, err)
	}
	if _, dbError := tx.Delete(&wish); dbError != nil {
		tx.Rollback()
		err := error_target_notes.DeleteErrorTarget
		err.Report = dbError.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, err)
	}
	tx.Commit()
	return make_result.ResponseWithJson(c, http.StatusOK, wish.Serializer())

}

func patchWish(c echo.Context) error {
	wishId := c.Param("wish_id")
	adminId := c.Get("current_admin_id")

	var param patchWishParam

	if err := c.Bind(&param); err != nil {
		bindErr := error_target_notes.ParamErrorTarget
		bindErr.Report = err.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, bindErr)
	}

	if invalidErr := param.Valid(); invalidErr != nil {
		err := error_target_notes.ValidErrorTarget
		err.Report = invalidErr.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, err)
	}

	tx := database.Engine.NewSession()
	defer tx.Close()
	tx.Begin()

	var wish model.Wish
	if ok, dbError := tx.ID(wishId).Where("admin_id = ?", adminId).Get(&wish); !ok && dbError != nil {
		err := error_target_notes.RecordErrorTarget
		err.Report = dbError.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, err)
	}

	if param.Data.Title != "" {
		wish.Title = param.Data.Title
	}
	if param.Data.Hope != "" {
		wish.Hope = param.Data.Hope
	}
	if param.Data.TargetId != 0 {
		var target model.Target
		if ok, dbError := tx.ID(param.Data.TargetId).Get(&target); !ok && dbError != nil {
			err := error_target_notes.RecordErrorTarget
			err.Report = dbError.Error()
			return make_result.ResponseWithJson(c, http.StatusBadRequest, err)
		}
		wish.TargetId = target.Id
	}
	wish.DesireLevel = param.Data.DesireLevel
	wish.ChallengeLevel = param.Data.ChallengeLevel
	wish.TimeLevel = param.Data.TimeLevel

	if _, dbError := tx.ID(wish.Id).Update(&wish); dbError != nil {
		tx.Rollback()
		err := error_target_notes.UpdateErrorTarget
		err.Report = dbError.Error()
		return make_result.ResponseWithJson(c, http.StatusBadRequest, err)

	}

	tx.Commit()

	return make_result.ResponseWithJson(c, http.StatusOK, wish.Serializer())

}
