package cmd

import (
	"EnterpriseWeb/EnterpriseWithMux/unicorn/pkg/database"
	"EnterpriseWeb/EnterpriseWithMux/unicorn/web/model"

	"github.com/spf13/cobra"
)

var migrateCMD = &cobra.Command{
	Use: "migrate",
	Run: func(cmd *cobra.Command, args []string) {
		database.EngineInit()
		defer database.Engine.Close()
		for _, i := range models() {
			database.Engine.AutoMigrate(i)
		}

	},
}

var models = func() []interface{} {
	return []interface{}{
		new(model.Company),
		new(model.Category),
		new(model.Country),
	}
}
