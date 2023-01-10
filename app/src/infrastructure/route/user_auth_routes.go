package route

import (
	"github.com/gin-gonic/gin"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/controllers/product"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/config"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/middleware"
)

func setUserAuthRoutes(rg *gin.RouterGroup, c *config.Config, jwt *middleware.JwtAuth) {

	userAuthController := product.NewUserAuthController()
	auth := rg.Group("/authorization")
	{
		auth.POST("/", func(c *gin.Context) {
			userAuthController.Post(c)
		})
	}
}
