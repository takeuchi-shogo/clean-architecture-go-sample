package route

import (
	"github.com/gin-gonic/gin"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/controllers/product"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/config"
	db "github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/database"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/middleware"
)

func setUserAuthRoutes(rg *gin.RouterGroup, db *db.DB, c *config.Config, jwt *middleware.JwtAuth) {

	userAuthController := product.NewUserAuthController(product.UserAuthControllerProvider{
		DB:  db,
		Jwt: jwt,
	})
	auth := rg.Group("/authorization")
	{
		auth.POST("/", func(c *gin.Context) {
			userAuthController.Post(c)
		})
	}
}
