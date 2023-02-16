package webserver

import "github.com/gin-gonic/gin"

func InitServer() {
	engine := gin.Default()
	engine.Group("v1")
}
