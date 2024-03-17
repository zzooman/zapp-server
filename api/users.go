package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/zzooman/zapp-server/db/sqlc"
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
	Username string   `json:"username" binding:"required"`
	Password string   `json:"password" binding:"required"`
	Email    string   `json:"email" binding:"required"`
	Phone    string   `json:"phone"`
	Location Location `json:"location" binding:"required"`
}
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}	
	payload := db.CreateUserParams{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,		
		Location: string(req.Location),		
	}
	if req.Phone != "" {payload.Phone = pgtype.Text{String: req.Phone, Valid: true}}

	user, err := server.store.CreateUser(ctx, payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}


// Delete User
type deleteUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}
func (server *Server) deleteUser(ctx *gin.Context) {
	var req deleteUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Find the user with the given ID
	user, err := server.store.GetUser(ctx, req.ID)
	if err != nil {		
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Delete the user
	err = server.store.DeleteUser(ctx, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
