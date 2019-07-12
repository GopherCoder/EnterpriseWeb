package cmd

import (
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/pkg/database"
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/pkg/log"
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/web/model"

	"github.com/spf13/cobra"
)

var syncCMD = &cobra.Command{
	Use:     "sync",
	Aliases: []string{"s", "S", "SYNC", "sync"},
	PreRun: func(cmd *cobra.Command, args []string) {
		database.EngineInit()
		log_tencent_votes.LoggerInit()
		log_tencent_votes.LoggerSysOut.Println("Start Sync Database")
		log_tencent_votes.LoggerSysOut.Println("SYNC DATABASE TABLE")
	},
	Run: func(cmd *cobra.Command, args []string) {
		log_tencent_votes.LoggerSysOut.Println("Start Sync Database")
		Run()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		log_tencent_votes.LoggerSysOut.Println("SYNC DATABASE TABLES DONE")
	},
}

func Models() []interface{} {
	return []interface{}{
		&model.Vote{},
		&model.Choice{},
	}
}
func Run() {
	for _, i := range Models() {
		database.Engine.AutoMigrate(i)
	}
}
