package route

import (
	"github.com/takeuchi-shogo/clean-architecture-golang/lib"
	db "github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/database"
)

type Routing struct {
	DB      *db.DB
	Handler lib.RequestHandler
}

func NewRouting(db *db.DB, handler lib.RequestHandler) *Routing {
	r := &Routing{
		DB:      db,
		Handler: handler,
	}

	r.setRouting()

	return r
}

func (r *Routing) setRouting() {

	v1 := r.Handler.Gin.Group("/v1/api")
	{
		addUserRoutes(v1)
	}
}
