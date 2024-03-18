package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/zzooman/zapp-server/db/sqlc"
)

// Create Account
type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,currency"`
}
func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if req.Currency == "" {
		req.Currency = "KRW"
	}	

	account, err := server.store.CreateAccount(ctx, db.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  0,
		Currency: req.Currency,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

// Get Account
type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}
func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

// Delete Account
type deleteAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}
func (server *Server) deleteAccount(ctx *gin.Context) {
	var req deleteAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteAccount(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, fmt.Sprintf("account %d deleted", req.ID))
}