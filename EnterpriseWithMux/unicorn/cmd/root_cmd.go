package cmd

import (
	"EnterpriseWeb/EnterpriseWithMux/unicorn/pkg/database"
	"EnterpriseWeb/EnterpriseWithMux/unicorn/pkg/router"
	"log"

	"github.com/spf13/cobra"
)

const PROJECT = "unicorn"

func init() {
	rootCMD.AddCommand(migrateCMD)
}

var rootCMD = &cobra.Command{
	Use: PROJECT,
	Run: func(cmd *cobra.Command, args []string) {
		database.EngineInit()
		defer database.Engine.Close()
		router.CollectionRouters()
	},
}

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		log.Panicln("Run Root CMD...")
	}
}
