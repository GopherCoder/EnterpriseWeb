package router

import (
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/pkg/middleware"
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/web/vote"
	"context"
	"encoding/json"
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

	var voteRouter vote.Vote
	http.HandleFunc("/vote", middleware.Logger(voteRouter.CreateVote))

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
