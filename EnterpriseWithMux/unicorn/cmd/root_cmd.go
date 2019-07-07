package cmd

import (
	"EnterpriseWeb/EnterpriseWithBeego/unicorn/pkg/database"
	"log"

	"github.com/spf13/cobra"
)

const PROJECT = "unicorn"

var rootCMD = &cobra.Command{
	Use: PROJECT,
	Run: func(cmd *cobra.Command, args []string) {
		database.EngineInit()
		defer database.Engine.Close()

	},
}

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		log.Panicln("Run Root CMD...")
	}
}
