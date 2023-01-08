package main

import (
	"errors"
	"fmt"
	"net/http"

	"code.gopub.tech/logs"
	"code.gopub.tech/wmm/model"
	"github.com/gin-gonic/gin"
	"github.com/youthlin/t"
)

func Use(h func(*gin.Context) (any, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		logs.Info(c, t.T("request url=%v header=%v", c.Request.URL, c.Request.Header))
		result, err := h(c)
		var (
			httpCode = http.StatusOK
			response model.Response
		)
		defer func() {
			logs.Info(c, t.T("url = %v, http code = %v, response = %v", c.Request.URL, httpCode, response))
		}()
		if err != nil {
			var (
				errorCode = int64(http.StatusInternalServerError)
				msg       = fmt.Sprintf("系统错误: %v", err)
			)
			var e model.ErrorCode
			if errors.As(err, &e) {
				httpCode = e.HttpCode
				errorCode = e.ErrorCode
				if e.Cause != nil {
					msg = fmt.Sprintf("%s: %v", e.Message, e.Cause)
				} else {
					msg = e.Message
				}
			}
			response = model.Response{
				Code:     errorCode,
				Messsage: msg,
				Data:     nil,
			}
			c.JSON(httpCode, response)
			return
		}
		response = model.Response{
			Code:     0,
			Messsage: "ok",
			Data:     result,
		}
		c.JSON(httpCode, response)
	}
}
