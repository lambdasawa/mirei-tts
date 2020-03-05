package errorhandler

import (
	"errors"
	"fmt"
	"mirei-tts/web/request"
	"mirei-tts/web/server"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	ErrorRes struct {
		Message string `json:"message"`
	}
)

func Set(s *server.Server) {
	s.Echo.HTTPErrorHandler = func(err error, c echo.Context) {
		handleHTTPError(c, s, err)
	}

	s.Log.Info("http error handler initialized", nil)
}

func handleHTTPError(c echo.Context, s *server.Server, err error) {
	{
		if errors.As(err, &request.ReadError{}) {
			sendErrorJSON(s, c, http.StatusBadRequest, "Failed to read request.")
			return
		}
	}

	s.Echo.DefaultHTTPErrorHandler(err, c)
}

func sendErrorJSON(s *server.Server, c echo.Context, status int, msg string) {
	if err := c.JSON(status, ErrorRes{Message: msg}); err != nil {
		s.Log.Warn(fmt.Sprintf("send error json: %v", err), nil)
	}
}
