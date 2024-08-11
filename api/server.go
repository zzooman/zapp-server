package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	
	server.setUpRouter(server.router)
	return server
}

func (server *Server) setUpRouter(router *gin.Engine) {	
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOrigins = []string{"http://localhost:8000"}
	router.Use(cors.New(config))

	router.POST("/login", server.loginUser)	
	router.POST("/user", server.createUser)	
	router.GET("/feeds", server.getFeeds)	
	router.GET("/feed/:id", server.getFeed)
	router.GET("/products", server.getProducts)		
	router.GET("/product/:id", server.getProduct)

	router.GET("/feeds/search", server.searchFeeds)
	router.GET("/products/search", server.searchProducts)
	router.GET("/search/hot", server.hotSearchTexts)
	
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.GET("/me", server.me)
	authRoutes.GET("/user/:id", server.getUser)
	authRoutes.GET("/products/liked", server.getProductsILiked)
	authRoutes.GET("/products/sold", server.getProductsISold)
	authRoutes.GET("/products/bought", server.getProductsIBought)
	authRoutes.PUT("/user/:id", server.updateUser)
	authRoutes.DELETE("/user/:id", server.deleteUser)
	authRoutes.POST("/feed", server.createFeed)
	authRoutes.POST("/product", server.createProduct)	
	authRoutes.POST("/product/:id/like", server.createLike)
	authRoutes.DELETE("/product/:id/unlike", server.deleteLike)
	
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