package handler

import (
	"github.com/labstack/echo/v4"
)

func BuildRouter(e *echo.Echo, userService userHandler) {
	e.POST("/signup", userService.SignUp)
	e.POST("/login", userService.SignUp)
	e.GET("/users", userService.GetUsers)
}
