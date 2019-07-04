package target

import "github.com/labstack/echo"

func Register(e *echo.Group) {
	e.GET("/other", getOtherTargetHandler)
	e.GET("/other/task/:task_id", getOneTaskHandler)
	e.POST("/other/task", createOneTaskHandler)
	e.PATCH("/other/task/:task_id", patchOneTaskHandler)
	e.DELETE("/other/task/:task_id", deleteOneTaskHandler)
	e.POST("/other/tasks_order", orderTaskHandler)
	e.PATCH("/other/task_thing/:task_id", patchOneThingHandler)
}
