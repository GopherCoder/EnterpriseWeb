package router

import (
	"EnterpriseWeb/EnterpriseWithMux/unicorn/pkg/middleware"
	"EnterpriseWeb/EnterpriseWithMux/unicorn/web/company"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func CollectionRouters() {
	r := mux.NewRouter()
	r.Use(middleware.LoggerMiddle)

	r.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		var results = make(map[string]interface{})
		results["code"] = http.StatusOK
		results["data"] = "pong"
		err := json.NewEncoder(writer).Encode(results)
		if err != nil {
			panic("CONNECT FAIL")
		}
	})

	s := r.PathPrefix("/v1/api").Subrouter()
	company.Register(s)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
