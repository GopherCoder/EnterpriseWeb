package vote

import (
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/pkg/database"
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/pkg/error"
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/web/make_response"
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/web/model"
	"encoding/json"
	"log"
	"net/http"
)

type Vote struct {
}

var Default = Vote{}

func (v Vote) GetVote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		make_response.Result(w, r, http.StatusBadRequest, error_tencent_votes.ErrorMethod)
		return
	}
}

func (v Vote) CreateVote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		make_response.Result(w, r, http.StatusBadRequest, error_tencent_votes.ErrorMethod)
		return
	}
	var param CreateParam
	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		make_response.Result(w, r, http.StatusBadRequest, err)
		return
	}

	if err := param.Valid(); err != nil {
		make_response.Result(w, r, http.StatusBadRequest, err)
		return
	}

	log.Println("Param: ", param)
	tx := database.Engine.Begin()
	defer tx.Close()
	var vote model.Vote
	vote = model.Vote{
		Title: param.Title,
	}
}

func (v Vote) PatchVote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		make_response.Result(w, r, http.StatusBadRequest, error_tencent_votes.ErrorMethod)
		return
	}
}

func (v Vote) DeleteVote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		make_response.Result(w, r, http.StatusBadRequest, error_tencent_votes.ErrorMethod)

		return
	}
}

func (v Vote) GetAllVotes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
	}
}
