package server

import (
	"github.com/gin-gonic/gin"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/config"
	db "github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/database"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/middleware"
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
		Gin:     middleware.NewRequestHandler().Gin,
		Port:    c.Server.Port,
		Routing: routing,
	}

	s.setCors(middleware.NewCors(c))

	return s
}

func (s *Server) setCors(cors *middleware.Cors) {
	s.Gin.Use(cors.Config)
}

func (s *Server) Run(port string) {
	s.Gin.Run(port)
}
