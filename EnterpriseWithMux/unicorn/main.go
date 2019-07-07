package main

import (
	"EnterpriseWeb/EnterpriseWithBeego/unicorn/cmd"
	"log"
)

func main() {
	log.Println("Run Web Server...")
	cmd.Execute()
}
