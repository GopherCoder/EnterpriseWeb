package cmd

import (
	"EnterpriseWeb/EnterpriseWithMux/unicorn/pkg/database"

	"github.com/spf13/cobra"
)

var dataCMD = &cobra.Command{
	Use: "data",
	Run: func(cmd *cobra.Command, args []string) {
		database.EngineInit()
		defer database.Engine.Close()

	},
}

type UnicornCompanyRequest struct {
	Request   string `json:"request"`
	ParseFunc func([]byte) interface{}
}

func Run() {}
