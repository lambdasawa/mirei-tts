package errorhandler

import (
	"errors"
	"fmt"
	"log"
	"mirei-tts/web/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	ErrorRes struct {
		Message string `json:"message"`
	}
)

func Set(e *echo.Echo) {
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		handleHTTPError(c, e, err)
	}
}

func handleHTTPError(c echo.Context, e *echo.Echo, err error) {
	{
		if errors.As(err, &request.ReadError{}) {
			sendErrorJSON(c, http.StatusBadRequest, "Failed to read request.")
			return
		}
	}

	e.DefaultHTTPErrorHandler(err, c)

}

func sendErrorJSON(c echo.Context, status int, msg string) {
	if err := c.JSON(status, ErrorRes{Message: msg}); err != nil {
		log.Println(fmt.Errorf("send error json: %v", err))
	}
}
