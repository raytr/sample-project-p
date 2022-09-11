package handler

import (
	"github.com/labstack/echo/v4"
)

func BuildRouter(e *echo.Echo, userService userHandler) {
	e.GET("/users", userService.GetUsers)
	e.POST("/signup", userService.SignUp)
}
