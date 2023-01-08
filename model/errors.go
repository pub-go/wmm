package model

import (
	"net/http"

	"code.gopub.tech/logs/pkg/arg"
	"github.com/youthlin/t"
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

func ErrInvalidParams(t *t.Translations) ErrorCode {
	return ErrorCode{HttpCode: http.StatusBadRequest, ErrorCode: 1000, Message: t.T("Invalid params")}
}

func ErrInvalidAccount(t *t.Translations) ErrorCode {
	return ErrorCode{HttpCode: http.StatusBadRequest, ErrorCode: 1000, Message: t.T("invalid account")}
}
