package target

import "github.com/labstack/echo"

func Register(e *echo.Group) {
	e.GET("/other", getOtherTargetHandler)
	e.POST("/other/task", createOneTaskHandler)
	e.PATCH("/other/task/:task_id", patchOneTaskHandler)
	e.DELETE("/other/task/:task_id", deleteOneTaskHandle)
	e.POST("/other/tasks_order", orderTaskHandler)
}
