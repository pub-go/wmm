package model

import (
	"net/http"

	"code.gopub.tech/logs/pkg/arg"
)

type ErrorCode struct {
	HttpCode  int    `json:"http_code,omitempty"`
	ErrorCode int64  `json:"error_code,omitempty"`
	Message   string `json:"message,omitempty"`
	Cause     error  `json:"cause,omitempty"`
}

var _ error = (*ErrorCode)(nil)

func (e ErrorCode) Error() string {
	return arg.JSON(e).String()
}

func (e ErrorCode) WithHttpCode(httpCode int) ErrorCode {
	e.HttpCode = httpCode
	return e
}
func (e ErrorCode) WithErrorCode(code int64) ErrorCode {
	e.ErrorCode = code
	return e
}
func (e ErrorCode) WithMessage(msg string) ErrorCode {
	e.Message = msg
	return e
}
func (e ErrorCode) WithCause(cause error) ErrorCode {
	e.Cause = cause
	return e
}

var (
	ErrInvalidParams  = ErrorCode{HttpCode: http.StatusBadRequest, ErrorCode: 1000, Message: "参数错误"}
	ErrInvalidAccount = ErrorCode{HttpCode: http.StatusBadRequest, ErrorCode: 1001, Message: "账号或密码错误"}
)
