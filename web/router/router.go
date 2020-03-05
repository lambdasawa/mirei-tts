package router

import (
	"mirei-tts/web/endpoint/speech"

	"github.com/labstack/echo/v4"
)

func Set(e *echo.Echo) {
	e.File("/", "public/index.html")
	e.GET("/speech", speech.Get)
}
