package router

import (
	"EnterpriseWeb/EnterpriseWithMux/unicorn/pkg/middleware"
	"EnterpriseWeb/EnterpriseWithMux/unicorn/web/company"
	"EnterpriseWeb/EnterpriseWithMux/unicorn/web/concept"
	"EnterpriseWeb/EnterpriseWithMux/unicorn/web/country"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
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
	country.Register(s)
	concept.Register(s)
	company.Register(s)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
