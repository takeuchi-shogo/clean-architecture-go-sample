package route

import (
	"github.com/gin-gonic/gin"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/controllers/product"
)

func addUserRoutes(rg *gin.RouterGroup) {

	usersController := product.NewUsersController()

	users := rg.Group("/users")

	users.GET("/:id", func(c *gin.Context) {
		usersController.Get(c)
	})
}
