package cmd

import (
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/database"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/web/model"
	"log"

	"github.com/spf13/cobra"
)

var syncCMD = &cobra.Command{
	Use: "sync2",
	Run: func(cmd *cobra.Command, args []string) {
		database.EngineInit()
		defer database.Engine.Close()
		sync()
	},
}

var sync = func() {
	for _, i := range models() {
		err := database.Engine.Sync2(i)
		if err != nil {
			log.Println(err)
		}
	}
}

func models() []interface{} {
	return []interface{}{
		new(model.Target),
		new(model.Admin),
	}
}
