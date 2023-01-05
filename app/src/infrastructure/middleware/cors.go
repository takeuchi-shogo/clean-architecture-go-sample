package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/config"
)

type Cors struct {
	Config gin.HandlerFunc
}

func NewCors(cfg *config.Config) *Cors {
	c := &Cors{}

	c.Config = cors.New(cors.Config{
		AllowOrigins: cfg.Cors.AllowOringins,
	})

	return c
}
