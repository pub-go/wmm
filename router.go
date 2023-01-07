package main

import (
	"code.gopub.tech/wmm/settings/hander"
	"github.com/gin-gonic/gin"
)

func setRouter(r *gin.Engine) {
	api := r.Group("/api")
	api.POST("/login", Use(hander.Login))
}
