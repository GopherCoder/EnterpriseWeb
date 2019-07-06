package aim

import "github.com/labstack/echo"

func Register(g *echo.Group) {
	g.GET("/aims", getAllAimsHandler)
	g.GET("/aim/:aim_id", getOneAimHandler)
	g.POST("/aim", createAimHandler)
	g.PATCH("/aim/:aim_id", patchAimHandler)
	g.DELETE("/aim/:aim_id", deleteAimHandler)
	g.PATCH("/aim_things/:aim_id/:task_id", patchAimThingsHandler)
	g.POST("/aim_task/:aim_id", createTaskHandler)
}
