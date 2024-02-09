package api

import (
	"github.com/qwerty22121998/gobase/pagination"
	"net/http"
)

type Meta struct {
	*pagination.Pagination
	Code    int    `json:"code"`
	Message string `json:"message"`
	Debug   any    `json:"debug,omitempty"`
}

type Response[T any] struct {
	Meta Meta `json:"meta"`
	Data *T   `json:"data,omitempty"`
}

func (r *Response[T]) WithDebug(data any) *Response[T] {
	r.Meta.Debug = data
	return r
}

func SuccessPagination[T any](data T, p *pagination.Pagination, message string) *Response[T] {
	return &Response[T]{
		Meta: Meta{
			Pagination: p,
			Code:       http.StatusOK,
			Message:    message,
		},
		Data: &data,
	}
}

func Info(message string) *Response[any] {
	return &Response[any]{
		Meta: Meta{
			Code:    http.StatusOK,
			Message: message,
		},
	}
}

func Success[T any](data T, message string) *Response[T] {
	return &Response[T]{
		Meta: Meta{
			Code:    http.StatusOK,
			Message: message,
		},
		Data: &data,
	}
}

func BadRequest(message string) *Response[any] {
	return &Response[any]{
		Meta: Meta{
			Code:    http.StatusBadRequest,
			Message: message,
		},
	}
}

func InternalError(err error) *Response[any] {
	return &Response[any]{
		Meta: Meta{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		},
	}
}
