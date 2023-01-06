package main

import (
	"github.com/takeuchi-shogo/clean-architecture-golang/lib"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure"
)

func main() {
	library := lib.NewLibrary()
	infrastructure.NewInfrastructure(library)
}
