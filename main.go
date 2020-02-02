package main

import "github.com/gin-gonic/gin"

var engine *gin.Engine

func init() {
	engine = gin.Default()
}

func main()  {
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	engine.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
