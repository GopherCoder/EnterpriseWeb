package other

import "github.com/labstack/echo"

func Register(e echo.Group) {
	e.GET("/other", getOtherTargetHandler)
	e.POST("/other/mission", createOneMissionHandler)
	e.PATCH("/other/mission/:mission_id", patchOneMissionHandler)
	e.GET("/other/missions", getAllMissionsHandler)
	e.DELETE("/other/mission/:mission_id", deleteOneMissionHandle)
	e.POST("/other/missions_order", orderMissionHandler)
}
