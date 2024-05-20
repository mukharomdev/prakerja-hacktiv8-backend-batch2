package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"uk-hacktiv8-prakerja-mukharom/config"
	"uk-hacktiv8-prakerja-mukharom/routes"
)

type Server struct {
	host   string
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		host:   config.GetConfig().ServerHost,
		port:   config.GetConfig().ServerPort,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigUserRoutes(s.server)

	log.Printf("Server running at port: %v", s.port)
	log.Fatal(router.Run(":" + s.port))
}
