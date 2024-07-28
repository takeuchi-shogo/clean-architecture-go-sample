package main

import (
	"github.com/takeuchi-shogo/clean-architecture-golang/lib"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/cmd"
)

func main() {
	library := lib.NewLibrary()

	cmd.RunServer(library)
}
