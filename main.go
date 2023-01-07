package main

import (
	"context"

	"code.gopub.tech/logs"
	"code.gopub.tech/wmm/settings"
	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

func MustInit() {
	logs.Info(ctx, "Hello, World")
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
