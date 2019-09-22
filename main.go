package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/raw", RawHandler)
	r.GET("/now", NowHandler)
	r.GET("/activity", ActivityHandler)

	if err := r.Run(); err != nil {
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
