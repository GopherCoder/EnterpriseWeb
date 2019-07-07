package company

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Register(r *mux.Router) *mux.Router {
	r.HandleFunc("/company", getAllCompaniesHandler).Methods(http.MethodPost)
	r.HandleFunc("/company/{company_id}", getOneCompanyHandler).Methods(http.MethodGet)
	r.HandleFunc("/company/{company_id}", patchOneCompanyHandler).Methods(http.MethodPatch)
	r.HandleFunc("/company/{company_id}", deleteOneCompanyHandler).Methods(http.MethodDelete)
	return r
}
