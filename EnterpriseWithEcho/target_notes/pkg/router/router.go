package router

import (
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/pkg/middleware"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/web/admin"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/web/aim"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/web/target"
	"EnterpriseWeb/EnterpriseWithEcho/target_notes/web/wish"
	"net/http"

	"github.com/labstack/gommon/log"

	middle "github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

const (
	VERSION = "v0.0.1"
)

func RouteCollection() {
	e := echo.New()
	e.Pre(middle.MethodOverride())

	e.Use(middle.Recover())
	e.Use(middle.CORS())
	e.Use(middle.BodyLimit("2M"))
	e.Use(middle.Logger())
	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level} ${line}")
	}
	e.Debug = true

	e.GET("/ping", func(context echo.Context) error {
		var result map[string]interface{}
		result = make(map[string]interface{})
		result["code"] = http.StatusOK
		result["data"] = "pong"
		return context.JSON(http.StatusOK, result)
	})

	groupWithOut := e.Group("/v1/api")
	admin.Register(groupWithOut)

	group := e.Group("/v1/api", middleware.Auth)

	{
		target.Register(group)
		aim.Register(group)
		wish.Register(group)
	}

	e.Logger.Fatal(e.Start(":7200"))
}
