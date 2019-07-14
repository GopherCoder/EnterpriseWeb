package vote

import (
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/pkg/database"
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/pkg/error"
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/pkg/middleware"
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/web/make_response"
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/web/model"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Vote struct {
}

var Default = Vote{}

func (v Vote) GetVote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		make_response.Result(w, r, http.StatusBadRequest, error_tencent_votes.ErrorMethod)
		return
	}
	id := r.URL.Query().Get("vote_id")
	var vote model.Vote
	if dbError := database.Engine.Where("id = ?", id).First(&vote).Error; dbError != nil {
		make_response.Result(w, r, http.StatusBadRequest, dbError.Error())
		return
	}
	make_response.Result(w, r, http.StatusOK, vote.Serializer(database.Engine))
}

func (v Vote) CreateVote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		make_response.Result(w, r, http.StatusBadRequest, error_tencent_votes.ErrorMethod)
		return
	}
	var param CreateParam
	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		make_response.Result(w, r, http.StatusBadRequest, err.Error())
		return
	}
	log.Println("Param: ", param)
	if err := param.Valid(); err != nil {
		make_response.Result(w, r, http.StatusBadRequest, err.Error())
		return
	}

	log.Println("Param: ", param)
	var (
		deadline time.Time
		err      error
	)
	if deadline, err = param.toTime(); err != nil {
		make_response.Result(w, r, http.StatusBadRequest, err.Error())
		return
	}

	admin := middleware.CurrentAdmin()
	log.Println("Admin: ", admin)
	tx := database.Engine.Begin()
	//defer tx.Close()
	var vote model.Vote
	vote = model.Vote{
		Title:       param.Title,
		AdminId:     admin.ID,
		Description: param.Description,
		DeadLine:    deadline,
		IsSingle:    param.IsSingle,
		IsAnonymous: param.IsAnonymous,
	}
	if dbError := tx.Save(&vote).Error; dbError != nil {
		tx.Rollback()
		make_response.Result(w, r, http.StatusBadRequest, dbError.Error())
		return
	}

	var choices = make([]model.Choice, len(param.Choice))
	for index, i := range param.Choice {
		choices[index].VoteId = vote.ID
		choices[index].Title = i
	}
	log.Println("Choice", choices)
	//for _, i := range choices {
	//	if dbError := tx.Save(&i).Error; dbError != nil {
	//		tx.Rollback()
	//		make_response.Result(w, r, http.StatusBadRequest, dbError.Error())
	//		return
	//	}
	//}

	if dbError := tx.Model(&vote).Association("Choice").Append(choices).Error; dbError != nil {
		tx.Rollback()
		make_response.Result(w, r, http.StatusBadRequest, dbError.Error())
		return
	}
	tx.Commit()
	make_response.Result(w, r, http.StatusOK, vote.Serializer(database.Engine))
}

func (v Vote) PatchVote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		make_response.Result(w, r, http.StatusBadRequest, error_tencent_votes.ErrorMethod)
		return
	}
	var param PatchParam

}

func (v Vote) DeleteVote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		make_response.Result(w, r, http.StatusBadRequest, error_tencent_votes.ErrorMethod)
		return
	}

	id := r.URL.Query().Get("vote_id")
	tx := database.Engine.Begin()
	defer tx.Close()

	var vote model.Vote
	if dbError := tx.Where("id = ?", id).First(&vote).Error; dbError != nil {
		tx.Rollback()
		make_response.Result(w, r, http.StatusBadRequest, dbError.Error())
		return
	}

	if dbError := tx.Model(&vote).Association("Choice").Clear().Error; dbError != nil {
		tx.Rollback()
		make_response.Result(w, r, http.StatusBadRequest, dbError.Error())
		return
	}
	if dbError := tx.Delete(&vote).Error; dbError != nil {
		tx.Rollback()
		make_response.Result(w, r, http.StatusBadRequest, dbError.Error())
		return
	}

	make_response.Result(w, r, http.StatusBadRequest, vote.Serializer(tx))
	tx.Commit()
	return
}

func (v Vote) GetAllVotes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
	}
	admin := middleware.CurrentAdmin()
	var votes []model.Vote
	if dbError := database.Engine.Where("admin_id = ?", admin.ID).Find(&votes).Error; dbError != nil {
		make_response.Result(w, r, http.StatusBadRequest, dbError.Error())
		return
	}
	var results []model.VoteSerializer
	for _, i := range votes {
		results = append(results, i.Serializer(database.Engine))
	}
	make_response.Result(w, r, http.StatusOK, results)
}

func (v Vote) Vote(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPatch {
		v.PatchVote(w, r)
		return
	}
	if r.Method == http.MethodDelete {
		v.DeleteVote(w, r)
		return
	}
	if r.Method == http.MethodGet {
		v.GetVote(w, r)
		return
	}
	if r.Method == http.MethodPost {
		v.CreateVote(w, r)
		return
	}
}
