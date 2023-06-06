package route

import (
	"github.com/gin-gonic/gin"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/controllers/product"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/config"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/middleware"
)

func setUserRoutes(rg *gin.RouterGroup, c *config.Config, jwt *middleware.JwtAuth) {

	usersController := product.NewUsersController(product.UsersControllerProvider{Jwt: jwt})

	users := rg.Group("/users").Use(jwt.CheckJwtToken)
	{
		users.GET("/:id", func(c *gin.Context) {
			usersController.Get(c)
		})
	}
}
