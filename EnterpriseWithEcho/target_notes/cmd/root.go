package cmd

import (
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/database"
	log_target_notes "EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/log"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/router"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCMD.AddCommand(syncCMD)
	if err := rootCMD.Execute(); err != nil {
		panic("Cmd Execute Fail")
	}
}

var rootCMD = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		log_target_notes.LOGInit()
		database.EngineInit()
		defer database.Engine.Close()

		router.RouteCollection()
	},
}
