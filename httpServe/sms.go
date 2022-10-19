package httpServe

import (
	"net/http"
	"plugin-demo-go/plugins/sms"

	"github.com/gin-gonic/gin"
)

func DoSMS(ctx *gin.Context) {
	// 获取 电话 和 信息
	phone := ctx.PostForm("phone")
	text := ctx.PostForm("text")

	// 获取插件客户端
	// cli := main.Context.(ctx)
	cli := GetCli(ctx)
	sms := cli.GetSrv("sms").(sms.SM)

	// 发送 sms
	res := sms.Send(phone, text)

	ctx.JSON(http.StatusOK, gin.H{
		"action":  "sms",
		"result":  "success",
		"message": res.Msg,
	})
}
