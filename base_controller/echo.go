package base_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/qwerty22121998/gobase/api"
	"github.com/qwerty22121998/gobase/errors"
	"net/http"
)

type IEchoController interface {
	Register(g *echo.Group)
}

type EchoController struct {
}

var customErrorHandler = make(map[int]func(c echo.Context, err error) error)

func RegisterErrorHandler(code int, fn func(c echo.Context, err error) error) {
	customErrorHandler[code] = fn
}

func init() {
	RegisterErrorHandler(http.StatusBadRequest, func(c echo.Context, err error) error {
		return c.JSON(http.StatusBadRequest, api.BadRequest(err.Error()))
	})
}

func ServeError(c echo.Context, err error) error {
	var wrapper *errors.Wrapper
	ok := errors.As(err, &wrapper)
	if ok {
		handler, ok := customErrorHandler[wrapper.HttpCode]
		if ok {
			return handler(c, wrapper)
		}
		return c.JSON(http.StatusInternalServerError, api.InternalError(wrapper))
	}
	return c.JSON(http.StatusInternalServerError, api.InternalError(err))
}

func Handle(c echo.Context, data *api.Response, err error) error {
	if err != nil {
		return ServeError(c, err)
	}
	return c.JSON(data.Meta.Code, data)
}

type Validator interface {
	Validate() error
}

func GenericServe[DTO any](c echo.Context, handler func(c echo.Context, req DTO) (*api.Response, error)) error {
	var req DTO
	if err := c.Bind(&req); err != nil {
		return errors.Wrap(http.StatusBadRequest, err)
	}
	if validator, ok := interface{}(req).(Validator); ok {
		if err := validator.Validate(); err != nil {
			return errors.Wrap(http.StatusBadRequest, err)
		}
	}
	resp, err := handler(c, req)
	return Handle(c, resp, err)
}
