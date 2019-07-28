package cmd

import (
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/database"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/model"
	"log"

	"github.com/spf13/cobra"
)

var SyncCMD = &cobra.Command{
	Use:     "sync",
	Aliases: []string{"s", "S", "-s", "-S"},
	PreRun: func(cmd *cobra.Command, args []string) {
		database.MySQLInit()
	},
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		defer database.Engine.Close()
	},
}

func Run() {
	tables := []interface{}{
		new(model.Admin),
		new(model.Address),
		new(model.Lottery),
		new(model.Admin2Lottery),
		new(model.AdminTakePart),
		new(model.WinnerLottery),
		new(model.Level),
	}
	for _, i := range tables {
		err := database.Engine.Sync2(i)
		if err != nil {
			log.Fatal("Sync Table Fail")
			return
		}
	}
}
