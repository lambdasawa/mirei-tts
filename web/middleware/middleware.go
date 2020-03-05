package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Set(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}
