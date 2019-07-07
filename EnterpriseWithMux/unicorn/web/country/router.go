package country

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Register(r *mux.Router) *mux.Router {
	r.HandleFunc("/countries", getAllCountriesHandler).Methods(http.MethodGet)
	r.HandleFunc("/country/{country_id}", getOneCountryHandler).Methods(http.MethodGet)
	r.HandleFunc("/country", postOneCountryHandler).Methods(http.MethodPost)
	r.HandleFunc("/country/{country_id}", deleteOneCountryHandler).Methods(http.MethodDelete)
	return r
}
