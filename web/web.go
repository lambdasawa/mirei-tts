package web

import (
	"mirei-tts/config"
	"mirei-tts/web/middleware"
	"mirei-tts/web/router"
	"mirei-tts/web/validator"

	"github.com/labstack/echo/v4"
)

func Start() error {
	e := echo.New()
	validator.Set(e)
	middleware.Set(e)
	router.Set(e)
	return e.Start(config.GetAddress())
}
