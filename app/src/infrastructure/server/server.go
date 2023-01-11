package server

import (
	"github.com/takeuchi-shogo/clean-architecture-golang/lib"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/config"
	db "github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/database"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/middleware"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/route"
)

type Server struct {
	DB      *db.DB
	Handler lib.RequestHandler
	Port    string
	Routing *route.Routing
}

func NewServer(c *config.Config, handler lib.RequestHandler, db *db.DB, routing *route.Routing) *Server {
	s := &Server{
		DB:      db,
		Handler: handler,
		Port:    c.Server.Port,
		Routing: routing,
	}

	s.setUp(middleware.NewCors(c))

	return s
}

func (s *Server) setUp(cors *middleware.Cors) {
	s.Handler.Gin.Use(cors.Config)
}

func (s *Server) Run(port string) {
	s.Handler.Gin.Run(port)
}
