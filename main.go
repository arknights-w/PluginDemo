package main

import (
	"plugin-demo-go/export"
	"plugin-demo-go/httpServe"

	"github.com/gin-gonic/gin"
)

const (
	address = "localhost:50051"
)

var cli = export.NewClient(address)

func InjectPluginClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("cli", cli)
		c.Next()
	}
}

func main() {
	engine := gin.Default()

	engine.Use(InjectPluginClient())

	engine.POST("install", httpServe.Install)
	engine.POST("update",httpServe.Update)

	engine.POST("uninstall",httpServe.UnInstall)
	engine.POST("banned",httpServe.Banned)

	engine.POST("doSMS", httpServe.DoSMS)

	engine.Run(":8000")
}
