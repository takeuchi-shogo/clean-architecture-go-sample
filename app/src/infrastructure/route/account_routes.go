package route

import (
	"github.com/gin-gonic/gin"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/controllers/product"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/config"
	db "github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/database"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/middleware"
)

func setAccountRoutes(rg *gin.RouterGroup, db *db.DB, c *config.Config, jwt *middleware.JwtAuth) {
	accountsController := product.NewAccountsController(product.AccountsControllerProvider{
		DB:  db,
		Jwt: jwt,
	})

	accounts := rg.Group("/accounts")
	{
		accounts.GET("/", func(c *gin.Context) {
			accountsController.Get(c)
		})
		accounts.POST("/", func(c *gin.Context) {
			accountsController.Post(c)
		})
		accounts.PATCH("/:id", func(c *gin.Context) {
			accountsController.Patch(c)
		})
	}
}
