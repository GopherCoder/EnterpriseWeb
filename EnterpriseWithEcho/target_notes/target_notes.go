package main

import (
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/cmd"
	"log"
)

var Env string

func main() {
	log.Println("Start Execute Web Server...")
	cmd.Execute()

}
