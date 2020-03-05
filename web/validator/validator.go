package validator

import (
	"mirei-tts/web/server"

	"github.com/go-playground/validator"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func Set(s *server.Server) {
	s.Echo.Validator = &CustomValidator{validator: validator.New()}

	s.Log.Info("validator intialized", nil)
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
