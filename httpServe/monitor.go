package httpServe

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Install(ctx *gin.Context) {
	// 读取文件，保存至本地
	file, _ := ctx.FormFile("file")
	dst := "./" + file.Filename
	ctx.SaveUploadedFile(file, dst)

	// 获取插件客户端，并执行 install
	cli := GetCli(ctx)
	cli.Install(dst)

	// 之后可以删除 file
	os.Remove(dst)

	ctx.JSON(http.StatusOK, gin.H{
		"action": "install",
		"result": "success",
	})
}

// 更新一般意味着手动修改了代码
// 或者安装后插件未生效的情况
// 仅传递服务名到插件服务器
func Update(ctx *gin.Context) {
	srv_name := ctx.PostForm("srv_name")

	// 获取插件客户端，并执行 install
	cli := GetCli(ctx)
	cli.Update(srv_name)

	ctx.JSON(http.StatusOK, gin.H{
		"action": "update",
		"result": "success",
	})
}

// 卸载
func UnInstall(ctx *gin.Context) {
	srv_name := ctx.PostForm("srv_name")

	// 获取插件客户端，并执行 install
	cli := GetCli(ctx)
	cli.UnInstall(srv_name)

	ctx.JSON(http.StatusOK, gin.H{
		"action": "uninstall",
		"result": "success",
	})
}

func Banned(ctx *gin.Context) {
	srv_name := ctx.PostForm("srv_name")

	// 获取插件客户端，并执行 install
	cli := GetCli(ctx)
	cli.Banned(srv_name)

	ctx.JSON(http.StatusOK, gin.H{
		"action": "banned",
		"result": "success",
	})
}
