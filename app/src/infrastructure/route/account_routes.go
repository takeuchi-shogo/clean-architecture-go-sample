package route

import (
	"github.com/gin-gonic/gin"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/controllers/product"
)

func addAccountRoute(rg *gin.RouterGroup) {
	accountsController := product.NewAccountsController()

	accounts := rg.Group("/accounts")

	accounts.POST("/:id", func(c *gin.Context) {
		accountsController.Post(c)
	})
	// accounts.PATCH("/", func(c *gin.Context) {
	// 	accountsController.Patch(c)
	// })
}
