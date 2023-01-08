package infrastructure

import (
	"github.com/takeuchi-shogo/clean-architecture-golang/lib"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/config"
	db "github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/database"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/route"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/server"
)

// type Infrastructure struct {
// 	config    config.Config
// 	datastore db.DB
// 	route     route.Routing
// 	server    server.Server
// }

// main.go の記述をシンプルにするために
func NewInfrastructure(lib lib.Library) {

	config := config.NewConfig(lib.Env)

	db := db.NewDB(config)

	route := route.NewRouting(config, db, lib.Handler)

	server := server.NewServer(config, lib.Handler, db, route)

	server.Run(server.Port)
}
