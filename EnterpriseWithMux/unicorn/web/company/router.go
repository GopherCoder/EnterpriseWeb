package company

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	r.HandleFunc("/companies/top", topCompaniesHandler).Methods(http.MethodGet)
	r.HandleFunc("/companies/sum", sumCompaniesHandler).Methods(http.MethodGet)
	r.HandleFunc("/company", oneCompanyHandler).Methods(http.MethodGet)
}
