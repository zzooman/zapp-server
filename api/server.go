package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/zzooman/zapp-server/db/sqlc"
)
type Server struct {
	store db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{
		store: store, 
		router: gin.Default(),
	}	
	
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.router.POST("/user", server.createUser)
	server.router.GET("/user/:id", server.getUser)
	server.router.DELETE("/user/:id", server.deleteUser)

	server.router.POST("/account", server.createAccount)
	server.router.GET("/account/:id", server.getAccount)
	server.router.DELETE("/account/:id", server.deleteAccount)

	server.router.POST("/transfer", server.createTransfer)
	server.router.GET("/transfer/:id", server.getTransfer)

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) map[string]any {
	return map[string]any{"error": err.Error()}
}