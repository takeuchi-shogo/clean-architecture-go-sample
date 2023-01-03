package main

import (
	"github.com/takeuchi-shogo/clean-architecture-golang/config"
	"github.com/takeuchi-shogo/sns-api/app/src/infrastructure"
)

func main() {

	config := config.NewConfig()
	db := infrastructure.NewDB(config)
}
