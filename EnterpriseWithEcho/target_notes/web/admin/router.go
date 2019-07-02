package admin

import "github.com/labstack/echo"

func Register(group *echo.Group) {
	group.POST("/register", registerHandler)
	group.POST("/login", loginHandler)
	group.POST("/logout", logoutHandler)
}
