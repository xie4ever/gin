package main

import (
	_ "gin/controller"
	"gin/router"
)

func main() {
	engine := router.GetEngine()
	engine.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
