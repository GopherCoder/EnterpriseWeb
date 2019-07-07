package concept

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	r.HandleFunc("/concept", createOneConceptHandler).Methods(http.MethodPost)
	r.HandleFunc("/concept/{concept_id}", getOneConceptHandler).Methods(http.MethodGet)
	r.HandleFunc("/concept/{concept_id}", deleteOneConceptHandler).Methods(http.MethodDelete)
}
