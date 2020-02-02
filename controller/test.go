package controller

import (
	"gin/router"
	"github.com/gin-gonic/gin"
)

func init() {
	engine := router.GetEngine()
	engine.GET("/ping", testMethod)
}

func testMethod(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
