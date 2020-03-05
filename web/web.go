package web

import (
	"mirei-tts/config"
	"mirei-tts/web/errorhandler"
	"mirei-tts/web/log"
	"mirei-tts/web/middleware"
	"mirei-tts/web/router"
	"mirei-tts/web/server"
	"mirei-tts/web/validator"

	"github.com/labstack/echo/v4"
)

func Start() error {
	server := &server.Server{
		Echo: echo.New(),
		Log:  log.New(),
	}

	server.Echo.HideBanner = true
	server.Echo.HidePort = true
	errorhandler.Set(server)
	validator.Set(server)
	middleware.Set(server)
	router.Set(server)

	server.Log.Info("server initialized", nil)

	return server.Echo.Start(config.GetAddress())
}
