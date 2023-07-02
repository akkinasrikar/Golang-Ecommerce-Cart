package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HttpServer struct {
	Router *gin.Engine
}

func Init(db *gorm.DB) (*HttpServer, error) {
	server := &HttpServer{}
	router := gin.Default()
	server.Router = router
	setUpRoutes(router, db)
	return server, nil
}

func (s *HttpServer) Start() error {
	s.Router.Run(":8080")
	return nil
}
