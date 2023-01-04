package main

import (
	"fmt"

	"github.com/takeuchi-shogo/clean-architecture-golang/lib"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/config"
	db "github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/database"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/route"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/server"
)

func main() {

	env := lib.NewEnv()
	fmt.Println(env)

	config := config.NewConfig(env)
	db := db.NewDB(config)

	route := route.NewRouting(config, db)

	server := server.NewServer(config, db, route)

	server.Run(server.Port)
}
