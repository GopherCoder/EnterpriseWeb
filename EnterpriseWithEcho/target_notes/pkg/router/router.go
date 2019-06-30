package router

import (
	"net/http"

	"github.com/labstack/echo"
)

const (
	VERSION = "v0.0.1"
)

func RouteCollection() {
	e := echo.New()
	e.GET("/ping", func(context echo.Context) error {
		var result map[string]interface{}
		result = make(map[string]interface{})
		result["code"] = http.StatusOK
		result["data"] = "pong"
		return context.JSON(http.StatusOK, result)
	})
	e.Logger.Fatal(e.Start(":7200"))
}
