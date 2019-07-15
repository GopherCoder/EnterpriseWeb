package main

import (
	"EnterpriseWeb/EnterpriseWithHTTP/tencent_vote/cmd"
	"log"
)

var Env string

func main() {
	log.Println("Start Web Server...")
	log.Println("Env : ", Env)
	cmd.Execute()
}
