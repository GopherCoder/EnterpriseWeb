package cmd

import (
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/database"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/router"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	RootCMD.AddCommand(SyncCMD)
}

var RootCMD = &cobra.Command{
	PreRun: func(cmd *cobra.Command, args []string) {
		database.MySQLInit()
	},
	Run: func(cmd *cobra.Command, args []string) {
		router.CollectionOfRouter()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		defer database.Engine.Close()
	},
}

func Execute() {
	if err := RootCMD.Execute(); err != nil {
		log.Fatal("Execute Root Command Fail")
		return
	}
}
