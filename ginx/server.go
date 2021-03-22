package ginx

import (
	"github.com/gin-gonic/gin"
)

func Default() *gin.Engine {
	server := gin.New()

	server.Use(LogHandlerFuncWithIgnorePath())
	server.Use(gin.Recovery())

	return server
}
