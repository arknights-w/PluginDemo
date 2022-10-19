package httpServe

import (
	"plugin-demo-go/export"

	"github.com/gin-gonic/gin"
)

type Context gin.Context

func GetCli(ctx *gin.Context) *export.Client {
	c, exists := ctx.Get("cli")
	if !exists {
		print("plugin-client is not exist")
	}
	cli := c.(*export.Client)
	return cli
}
