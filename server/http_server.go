package server

import (
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	Router *gin.Engine
}

func Init() (*HttpServer, error) {
	server := &HttpServer{}
	router := gin.Default()
	server.Router = router
	setUpRoutes(router)
	return server, nil
}

func (s *HttpServer) Start() error {
	s.Router.Run(":8080")
	return nil
}
