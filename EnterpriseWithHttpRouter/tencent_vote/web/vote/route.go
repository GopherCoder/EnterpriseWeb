package vote

import "net/http"

type Vote struct {
}

var Default = Vote{}

func (v Vote) GetVote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}
}

func (v Vote) CreateVote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
}

func (v Vote) PatchVote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		return
	}
}

func (v Vote) DeleteVote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		return
	}
}

func (v Vote) GetAllVotes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
	}
}
