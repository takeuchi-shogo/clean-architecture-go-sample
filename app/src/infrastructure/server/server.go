package server

import (
	"github.com/gin-gonic/gin"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/config"
	db "github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/database"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/route"
)

type Server struct {
	DB      *db.DB
	Gin     *gin.Engine
	Port    string
	Routing *route.Routing
}

func NewServer(c *config.Config, db *db.DB, routing *route.Routing) *Server {
	s := &Server{
		DB:      db,
		Gin:     gin.Default(),
		Port:    c.Server.Port,
		Routing: routing,
	}

	return s
}

func (s *Server) Run(port string) {
	s.Gin.Run(port)
}
