package api

import (
	"fmt"
	"net/http"
	"time"

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
func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:          user.Username,
		Email:             user.Email,
		Phone:             user.Phone,
		Location:          user.Location,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

type createUserRequest struct {
	Username string   `json:"username" binding:"required,alphanum,min=4,max=32"`
	Password string   `json:"password" binding:"required,min=6,max=32"`
	Email    string   `json:"email" binding:"required,email,max=64"`
	Phone    string   `json:"phone" binding:"omitempty"`
	Location Location `json:"location" binding:"required"`
}
type userResponse struct {
	Username          string             `json:"username"`	
	Email             string             `json:"email"`
	Phone             pgtype.Text        `json:"phone"`
	Location          pgtype.Text        `json:"location"`
	PasswordChangedAt pgtype.Timestamptz `json:"password_changed_at"`
	CreatedAt         pgtype.Timestamptz `json:"created_at"`
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
	}
	if req.Location != "" {payload.Location = pgtype.Text{String: string(req.Location), Valid: true}}
	if req.Phone != "" {payload.Phone = pgtype.Text{String: req.Phone, Valid: true}}

	user, err := server.store.CreateUser(ctx, payload)
	if err != nil {
		fmt.Println(`CreateUser Error:`, err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
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


// Login User
type loginUserRequest struct {
	Username string   `json:"username" binding:"required,alphanum,min=4,max=32"`
	Password string   `json:"password" binding:"required,min=6,max=32"`
}
type loginUserResponse struct {
	AccessToken string 		  `json:"access_token"`
	User 	   	userResponse  `json:"user"`
}
func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if(!utils.CheckPasswordHash(req.Password, user.Password)) {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(user.Username, time.Duration(time.Duration.Minutes(720)))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := loginUserResponse{
		AccessToken: accessToken,
		User: newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)
}