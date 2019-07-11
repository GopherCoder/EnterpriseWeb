package cmd

import (
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/pkg/database"
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/pkg/log"
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/pkg/router"
	"log"

	"github.com/spf13/cobra"
)

var RootCMD = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		database.EngineInit()
		defer database.Engine.Close()
		log_tencent_votes.LoggerInit()
		router.CollectionOfRouter()
	},
}

func Execute() {
	if err := RootCMD.Execute(); err != nil {
		log.Panic("ROOT CMD RUN FAIL")
	}
}
