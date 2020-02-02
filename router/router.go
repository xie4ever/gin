package router

import "github.com/gin-gonic/gin"

var engine *gin.Engine

func init() {
	engine = gin.Default()
}

// GetEngine ...
func GetEngine() *gin.Engine {
	return engine
}
