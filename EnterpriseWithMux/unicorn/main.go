package main

import (
	"EnterpriseWeb/EnterpriseWithMux/unicorn/cmd"
	"log"
)

var ENV string

func main() {
	log.Println("Run Web Server...")
	cmd.Execute()
}
