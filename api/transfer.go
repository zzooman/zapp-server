package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/zzooman/zapp-server/db/sqlc"
	"github.com/zzooman/zapp-server/token"
)

// Create Transfer
type createTransferRequest struct {
	FromAccountID int64 `json:"from_account_id" binding:"required,min=1"`
    ToAccountID   int64 `json:"to_account_id" binding:"required,min=1"`
    Amount        int64 `json:"amount" binding:"required,min=1"`
}
func (server *Server) createTransfer(ctx *gin.Context) {
	var req createTransferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fromAccount, err := server.validAccount(ctx, req.FromAccountID)
	username := ctx.MustGet(AUTH_TOKEN).(*token.Payload).Username
	if fromAccount.Owner != username {
		err := fmt.Errorf("account does not belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	toAccount, err := server.validAccount(ctx, req.ToAccountID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return	
	}
	if fromAccount.Currency != toAccount.Currency {
		err := fmt.Errorf("currency mismatch: %s %s", fromAccount.Currency, toAccount.Currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if fromAccount.Balance < req.Amount {
		err := fmt.Errorf("insufficient balance: %d", fromAccount.Balance)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := server.store.TransferTx(ctx, db.CreateTransferParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,				
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// Get Transfer
type getTransferRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}
func (server *Server) getTransfer(ctx *gin.Context) {
	var req getTransferRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	transfer, err := server.store.GetTransfer(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, transfer)
}

// validate account
func (server *Server) validAccount(ctx *gin.Context, accountId int64) (account db.Account, err error) {
	account, err = server.store.GetAccount(ctx, accountId)
	if err != nil {
		return account, err
	}
	return account, nil
}	