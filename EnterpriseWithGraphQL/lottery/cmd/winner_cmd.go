package cmd

import (
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/database"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/model"

	"github.com/spf13/cobra"
)

var WinnerCommand = &cobra.Command{
	Use:     "winner",
	Aliases: []string{"w", "W", "-w", "-W", "winner", "Winner"},
	PreRun: func(cmd *cobra.Command, args []string) {
		database.MySQLInit()
	},
	Run: func(cmd *cobra.Command, args []string) {
		for _, i := range model.DefaultWinnerLottery {
			database.Engine.InsertOne(&i)
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		defer database.Engine.Close()
	},
}
