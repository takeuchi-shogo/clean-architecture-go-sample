package route

import (
	"github.com/gin-gonic/gin"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/config"
	db "github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/database"
)

type Routing struct {
	DB   *db.DB
	Gin  *gin.Engine
	Port string
}

func NewRouting(c *config.Config, db *db.DB) *Routing {
	r := &Routing{
		DB:  db,
		Gin: gin.Default(),
	}

	r.setRouting()

	return r
}

func (r *Routing) setRouting() {

	v1 := r.Gin.Group("/v1/api")
	{
		addUserRoutes(v1)
	}
}
