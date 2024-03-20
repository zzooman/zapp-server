package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/zzooman/zapp-server/db/sqlc"
	"github.com/zzooman/zapp-server/token"
)
type Server struct {
	store db.Store
	tokenMaker token.Maker
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{
		store: store, 
		tokenMaker: token.NewPasetoMaker(),
		router: gin.Default(),
	}			
	
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setUpRouter(server.router)
	return server
}

func (server *Server) setUpRouter(router *gin.Engine) {
	router.POST("/login", server.loginUser)
	
	router.POST("/user", server.createUser)
	router.GET("/user/:id", server.getUser)
	router.DELETE("/user/:id", server.deleteUser)	

	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.DELETE("/account/:id", server.deleteAccount)

	router.POST("/transfer", server.createTransfer)
	router.GET("/transfer/:id", server.getTransfer)
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) map[string]any {
	return map[string]any{"error": err.Error()}
}