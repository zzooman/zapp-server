package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/zzooman/zapp-server/db/sqlc"
	"github.com/zzooman/zapp-server/token"
	"github.com/zzooman/zapp-server/utils"
)

// Create User
func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:          user.Username,
		Email:             user.Email,
		Phone:             user.Phone,		
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

type createUserRequest struct {
	Username string   `json:"username" binding:"required,alphanum,min=4,max=32"`
	Password string   `json:"password" binding:"required,min=6,max=32"`
	Email    string   `json:"email" binding:"required,email,max=64"`
	Phone    string   `json:"phone" binding:"omitempty"`	
}
type userResponse struct {
	Username          string             `json:"username"`	
	Email             string             `json:"email"`
	Phone             pgtype.Text        `json:"phone"`	
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

// Update User
type updateUserRequest struct {
	Username string `uri:"username" binding:"required"`
	Email    string `json:"email" binding:"omitempty,email,max=64"`
	Phone    string `json:"phone" binding:"omitempty"`
	Profile  string `json:"profile" binding:"omitempty"`
}
func (server *Server) updateUser(ctx *gin.Context) {
	var req updateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := ctx.MustGet(AUTH_TOKEN).(*token.Payload)
	user, err := server.store.GetUser(ctx, payload.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Update the user with the provided values
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Phone != "" {
		user.Phone = pgtype.Text{String: req.Phone, Valid: true}
	}
	if req.Profile != "" {
		user.Profile = pgtype.Text{String: req.Profile, Valid: true}
	}

	// Save the updated user
	updatedUser, err := server.store.UpdateUser(ctx, db.UpdateUserParams{
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
		Profile:  user.Profile,		
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newUserResponse(updatedUser))
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
	Username string `json:"username" binding:"required,alphanum,min=4,max=32"`
	Password string `json:"password" binding:"required,min=6,max=32"`
}
type loginUserResponse struct {
	AuthToken string      `json:"auth_token"`
	User      userResponse `json:"user"`
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

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("invalid credentials")))
		return
	}

	// 토큰 만료 시간을 12시간으로 설정
	authToken, err := server.tokenMaker.CreateToken(user.Username, time.Hour*12)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		AuthToken: authToken,
		User:      newUserResponse(user),
	}

	// Set the cookie with the access token, 만료 시간도 12시간으로 설정
	ctx.SetCookie("auth_token", authToken, int((time.Hour * 12).Seconds()), "/", "localhost", false, false)
	ctx.JSON(http.StatusOK, rsp)
}


func (server *Server) me(ctx *gin.Context) {		
	payload := ctx.MustGet(AUTH_TOKEN).(*token.Payload)	
	user, err := server.store.GetUser(ctx, payload.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

