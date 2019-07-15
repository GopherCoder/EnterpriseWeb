package router

import (
	"EnterpriseWeb/EnterpriseWithHTTP/tencent_vote/pkg/middleware"
	"EnterpriseWeb/EnterpriseWithHTTP/tencent_vote/web/admin"
	"EnterpriseWeb/EnterpriseWithHTTP/tencent_vote/web/vote"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func CollectionOfRouter() {

	http.HandleFunc("/ping", middleware.Logger(func(writer http.ResponseWriter, request *http.Request) {
		var result = make(map[string]interface{})
		result["code"] = 200
		result["data"] = "pong"
		err := json.NewEncoder(writer).Encode(&result)
		if err != nil {
			log.Panic("CONNECT TO SERVER FAIL")
		}
	}))
	v1 := fmt.Sprintf("/v1/api")

	var voteRouter vote.Vote
	http.HandleFunc(fmt.Sprintf(v1+"/vote"), middleware.Auth(middleware.Logger(voteRouter.Vote)))
	http.HandleFunc(fmt.Sprintf(v1+"/votes"), middleware.Auth(middleware.Logger(voteRouter.GetAllVotes)))

	var adminRouter admin.Admin
	http.HandleFunc(fmt.Sprintf(v1+"/register"), middleware.Logger(adminRouter.Register))
	http.HandleFunc(fmt.Sprintf(v1+"/login"), middleware.Logger(adminRouter.Login))
	http.HandleFunc(fmt.Sprintf(v1+"/logout"), middleware.Logger(adminRouter.Logout))

	//服务启动
	go func() {
		if err := http.ListenAndServe(":7201", nil); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	_, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	log.Println("shutting down")
	os.Exit(0)
}
