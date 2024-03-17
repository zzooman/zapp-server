package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/zzooman/zapp-server/db/sqlc"
)
type Server struct {
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{
		store: store, 
		router: gin.Default(),
	}	

	server.router.POST("/users", server.createUser)
	server.router.DELETE("/users/:id", server.deleteUser)
	server.router.POST("/accounts", server.createAccount)

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) map[string]any {
	return map[string]any{"error": err.Error()}
}