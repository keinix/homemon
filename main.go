package main

import (
	"github.com/gin-gonic/gin"
	"homemon/config"
	"homemon/scanner"
)

func main() {
	localConfig := config.LocalConfig{}
	if err := localConfig.SetFromFile("config/config.yaml"); err != nil {
		panic(err)
	}
	go scanner.ScanNetwork(&localConfig.Scan)
	startWebServer(&localConfig.Server)
}

func startWebServer(c *config.ServerConfig) {
	r := gin.Default()
	r.GET("/raw", RawHandler)
	r.GET("/now", NowHandler)
	r.GET("/activity", ActivityHandler)

	if err := r.Run(":" + c.Port); err != nil {
		panic(err)
	}
}

func RawHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "/raw endpoint hit",
	})
}

func NowHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "/now endpoint hit",
	})
}

func ActivityHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "/activity endpoint hit",
	})
}
