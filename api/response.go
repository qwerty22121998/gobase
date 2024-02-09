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

type Response struct {
	Meta Meta `json:"meta"`
	Data any  `json:"data,omitempty"`
}

func (r *Response) WithDebug(data any) *Response {
	r.Meta.Debug = data
	return r
}

func SuccessPagination(data any, p *pagination.Pagination, message string) *Response {
	return &Response{
		Meta: Meta{
			Pagination: p,
			Code:       http.StatusOK,
			Message:    message,
		},
		Data: &data,
	}
}

func Info(message string) *Response {
	return &Response{
		Meta: Meta{
			Code:    http.StatusOK,
			Message: message,
		},
	}
}

func Success[T any](data T, message string) *Response {
	return &Response{
		Meta: Meta{
			Code:    http.StatusOK,
			Message: message,
		},
		Data: &data,
	}
}

func BadRequest(message string) *Response {
	return &Response{
		Meta: Meta{
			Code:    http.StatusBadRequest,
			Message: message,
		},
	}
}

func InternalError(err error) *Response {
	return &Response{
		Meta: Meta{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		},
	}
}
