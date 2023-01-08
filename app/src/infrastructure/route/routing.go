package route

import (
	"github.com/takeuchi-shogo/clean-architecture-golang/lib"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/config"
	db "github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/database"
)

type Routing struct {
	DB      *db.DB
	Handler lib.RequestHandler
}

func NewRouting(c *config.Config, db *db.DB, handler lib.RequestHandler) *Routing {
	r := &Routing{
		DB:      db,
		Handler: handler,
	}

	r.setRouting(c)

	return r
}

func (r *Routing) setRouting(c *config.Config) {

	v1 := r.Handler.Gin.Group("/v1.0")
	{
		setAccountRoutes(v1)
		setUserRoutes(v1, c)
	}
}
