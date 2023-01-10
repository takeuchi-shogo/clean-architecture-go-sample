package lib

import "github.com/gin-gonic/gin"

type RequestHandler struct {
	Gin *gin.Engine
}

func NewRequestHandler() RequestHandler {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	return RequestHandler{
		Gin: engine,
	}
}
