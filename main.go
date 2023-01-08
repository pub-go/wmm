package main

import (
	"code.gopub.tech/wmm/settings"
	"github.com/gin-gonic/gin"
)

func MustInit() {
	if err := settings.Init(); err != nil {
		panic(err)
	}
}

func main() {
	MustInit()
	r := gin.Default()
	setRouter(r)
	r.Run()
}
