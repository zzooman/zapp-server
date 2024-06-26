package api

import (
	"github.com/gin-contrib/cors"
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
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	router.POST("/login", server.loginUser)	
	router.POST("/user", server.createUser)		
	router.GET("/posts", server.getPosts)	
	router.GET("/posts/search", server.searchPosts)
	router.GET("/post/:id", server.getPost)
	router.GET("/search/hot", server.hotSearchTexts)
	
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.GET("/me", server.me)
	authRoutes.GET("/user/:id", server.getUser)
	authRoutes.GET("/posts/liked", server.getPostsILiked)
	authRoutes.GET("/posts/sold", server.getPostsISold)
	authRoutes.GET("/posts/bought", server.getPostsIBought)
	authRoutes.PUT("/user/:id", server.updateUser)
	authRoutes.DELETE("/user/:id", server.deleteUser)
	authRoutes.POST("/post", server.createPost)	
	authRoutes.POST("/post/:id/like", server.createLike)
	authRoutes.DELETE("/post/:id/unlike", server.deleteLike)
	authRoutes.POST("/room", server.enterChatRoom)
	authRoutes.GET("/ws/:room_id", server.connectWS)
	authRoutes.GET("/messages/:room_id", server.getMessages)
	authRoutes.GET("/rooms/all", server.getAllRooms)
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) map[string]any {
	return map[string]any{"error": err.Error()}
}