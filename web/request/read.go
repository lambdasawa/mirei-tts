package request

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type (
	ReadError struct {
		Err error
	}
)

func (e ReadError) Error() string {
	return e.Err.Error()
}

func Read(c echo.Context, req interface{}) error {
	if err := c.Bind(req); err != nil {
		return ReadError{
			Err: fmt.Errorf("bind request: %v", err),
		}
	}

	if err := c.Validate(req); err != nil {
		return ReadError{
			Err: fmt.Errorf("validate request: %v", err),
		}
	}

	return nil
}
