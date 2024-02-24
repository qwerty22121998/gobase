package base_controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type echoValidator struct {
	validator *validator.Validate
}

func NewEchoValidator() echo.Validator {
	return &echoValidator{validator: validator.New()}
}

func (v *echoValidator) Validate(i any) error {
	if err := v.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
