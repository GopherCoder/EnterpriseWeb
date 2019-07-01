package router

import (
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/web/wish"
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

	group := e.Group("/v1/api")
	wish.Register(group)

	e.Logger.Fatal(e.Start(":7200"))
}

func Middleware(h http.Header) {
}
