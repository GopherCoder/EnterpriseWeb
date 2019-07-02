package wish

import "github.com/labstack/echo"

func Register(echo *echo.Group) {
	echo.GET("/wish/:wish_id", getWish)
	echo.GET("/wishes", getAllWish)
	echo.POST("/wish", postWish)
	echo.PATCH("/wish/:wish_id", patchWish)
	echo.DELETE("/wish/:wish_id", deleteWish)
}
