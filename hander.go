package main

import (
	"errors"
	"fmt"
	"net/http"

	"code.gopub.tech/logs"
	"code.gopub.tech/wmm/model"
	"github.com/gin-gonic/gin"
)

func Use(hander func(*gin.Context) (any, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		logs.Info(c, "request url=%v header=%v", c.Request.URL, c.Request.Header)
		response, err := hander(c)
		if err != nil {
			var (
				httpCode  = http.StatusInternalServerError
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
			c.JSON(httpCode, model.Response[any]{
				Code:     errorCode,
				Messsage: msg,
				Data:     nil,
			})
			return
		}
		c.JSON(http.StatusOK, model.Response[any]{
			Code:     0,
			Messsage: "ok",
			Data:     response,
		})
	}
}
