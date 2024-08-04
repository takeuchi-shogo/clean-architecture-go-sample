package route

import (
	"github.com/gin-gonic/gin"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/controllers/product"
	db "github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/database"
)

func setPostRoutes(rg *gin.RouterGroup, db *db.DB) {

	postsController := product.NewPostsController(product.PostsControllerProvider{
		DB: db,
	})

	posts := rg.Group("/posts")
	{
		posts.POST("/", func(c *gin.Context) {
			postsController.Post(c)
		})
		posts.GET("/:id", func(c *gin.Context) {
			postsController.Get(c)
		})
	}
}
