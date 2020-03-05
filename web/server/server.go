package server

import (
	"mirei-tts/web/log"

	"github.com/labstack/echo/v4"
)

type (
	Server struct {
		Echo *echo.Echo
		Log  *log.Log
	}
)
