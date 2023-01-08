package main

import (
	"code.gopub.tech/wmm/hander"
	"github.com/gin-gonic/gin"
)

func setRouter(r *gin.Engine) {
	// 设置为 true: 当  gin.Context 无法处理 context.Context 相关方法时 用 c.Request.Context() 处理
	r.ContextWithFallback = true
	r.Use(hander.I18n) // 为每个请求设置翻译
	api := r.Group("/api")
	api.POST("/login", Use(hander.Login))
}
