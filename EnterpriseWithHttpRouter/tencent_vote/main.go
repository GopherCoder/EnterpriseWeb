package main

import (
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/cmd"
	"log"
)

func main() {
	log.Print("Start Web Server...")
	cmd.Execute()
}
