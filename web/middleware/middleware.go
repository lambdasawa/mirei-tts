package middleware

import (
	"mirei-tts/web/server"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Set(s *server.Server) {
	s.Echo.Use(middleware.RequestID())
	s.Echo.Use(middleware.Recover())
	s.Echo.Use(accessLog(s))

	s.Log.Info("middleware initialized", nil)
}

func accessLog(s *server.Server) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			start := time.Now()
			err := next(c)
			latency := time.Since(start).String()

			if err != nil {
				c.Error(err)
			}

			method := req.Method

			url := req.URL.String()

			status := res.Status

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}

			remoteIP := c.RealIP()

			referer := req.Referer()

			ua := req.UserAgent()

			logParams := map[string]interface{}{
				"method":   method,
				"url":      url,
				"id":       id,
				"status":   status,
				"latency":  latency,
				"remoteIP": remoteIP,
				"referer":  referer,
				"ua":       ua,
				"error":    err,
				"request":  c.Get("req"),
			}
			switch {
			case status < 400:
				s.Log.Info("access log", logParams)
			case 400 <= status && status < 500:
				s.Log.Warn("warn status", logParams)
			default:
				s.Log.Error("error status", logParams)
			}

			return err
		}
	}
}
