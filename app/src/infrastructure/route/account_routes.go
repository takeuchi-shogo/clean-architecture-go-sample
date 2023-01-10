package route

import (
	"github.com/gin-gonic/gin"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/controllers/product"
)

func setAccountRoutes(rg *gin.RouterGroup) {
	accountsController := product.NewAccountsController(product.AccountsControllerProvider{})

	accounts := rg.Group("/accounts")
	{
		accounts.GET("/", func(c *gin.Context) {
			accountsController.Get(c)
		})
		accounts.POST("/", func(c *gin.Context) {
			accountsController.Post(c)
		})
		// accounts.PATCH("/", func(c *gin.Context) {
		// 	accountsController.Patch(c)
		// })
	}
}
