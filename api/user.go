package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/zzooman/zapp-server/db/sqlc"
	"github.com/zzooman/zapp-server/utils"
)

// Create User
type Location string
const (
	LocationSeoul    Location = "Seoul"
	LocationBusan    Location = "Busan"
	LocationIncheon  Location = "Incheon"
	LocationDaegu    Location = "Daegu"
	LocationDaejeon  Location = "Daejeon"
	LocationGwangju  Location = "Gwangju"
	LocationUlsan    Location = "Ulsan"
	LocationSejong   Location = "Sejong"
	LocationGyeonggi Location = "Gyeonggi"
	// Add more location options as needed
)
type createUserRequest struct {
	Username string   `json:"username" binding:"required,alphanum,min=4,max=32"`
	Password string   `json:"password" binding:"required,min=6,max=32"`
	Email    string   `json:"email" binding:"required,email,max=64"`
	Phone    string   `json:"phone" binding:"omitempty,phone"`
	Location Location `json:"location" binding:"required"`
}
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}	
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {    
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	payload := db.CreateUserParams{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,		
		Location: string(req.Location),		
	}
	if req.Phone != "" {payload.Phone = pgtype.Text{String: req.Phone, Valid: true}}

	user, err := server.store.CreateUser(ctx, payload)
	if err != nil {
		fmt.Println(`CreateUser Error:`, err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// Get User
type getUserRequest struct {
	Username string `uri:"username" binding:"required"`
}
func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)

}


// Delete User
type deleteUserRequest struct {
	Username string `uri:"username" binding:"required"`
}
func (server *Server) deleteUser(ctx *gin.Context) {
	var req deleteUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Find the user with the given ID
	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {		
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Delete the user
	err = server.store.DeleteUser(ctx, user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
