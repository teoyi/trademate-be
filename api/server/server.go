package server

import (
	"github.com/gin-gonic/gin"
	"github.com/trademate-be/internal/controller/sqlc/store"
)

type Server struct {
	store  store.Store
	router *gin.Engine
}

func NewServer(store store.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
