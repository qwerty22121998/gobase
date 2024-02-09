package errors

import (
	"errors"
	"fmt"
	"net/http"
)

type Wrapper struct {
	BaseError error
	Message   string
	HttpCode  int
}

func (w *Wrapper) Error() string {
	if w.Message != "" {
		return w.Message
	}
	if w.BaseError != nil {
		return w.BaseError.Error()
	}
	return ""
}

func As(err error, target any) bool {
	return errors.As(err, target)
}

func New(template string, args ...any) error {
	return errors.New(fmt.Sprintf(template, args...))
}

func Wrap(code int, err error) error {
	return NewError(code, "", err)
}

func NewError(code int, message string, base error) error {
	return &Wrapper{
		BaseError: base,
		Message:   message,
		HttpCode:  code,
	}
}

func BadRequest(template string, args ...any) error {
	return NewError(http.StatusBadRequest, fmt.Sprintf(template, args...), nil)
}

func InternalError(err error) error {
	return NewError(http.StatusInternalServerError, "", err)
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}
