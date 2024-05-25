package api

// // Create Account
// type createAccountRequest struct {
// 	Currency string `json:"currency" binding:"required,currency"`
// }
// func (server *Server) createAccount(ctx *gin.Context) {
// 	var req createAccountRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	if req.Currency == "" {
// 		req.Currency = "KRW"
// 	}

// 	auth_payload := ctx.MustGet(AUTH_TOKEN).(*token.Payload)
// 	account, err := server.store.CreateAccount(ctx, db.CreateAccountParams{
// 		Owner:    auth_payload.Username,
// 		Balance:  0,
// 		Currency: req.Currency,
// 	})
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, account)
// }

// // Get Account
// type getAccountRequest struct {
// 	ID int64 `uri:"id" binding:"required,min=1"`
// }
// func (server *Server) getAccount(ctx *gin.Context) {
// 	var req getAccountRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	account, err := server.store.GetAccount(ctx, req.ID)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	if account.Owner != ctx.MustGet(AUTH_TOKEN).(*token.Payload).Username {
// 		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
// 		return

// 	}
// 	ctx.JSON(http.StatusOK, account)
// }

// // Get Accounts (list)
// type listAccountsRequest struct {
// 	Limit  int32 `uri:"limit"`
// 	Offset int32 `uri:"offset"`
// }
// func (server *Server) listAccounts(ctx *gin.Context) {
// 	var req listAccountsRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}
// 	username := ctx.MustGet(AUTH_TOKEN).(*token.Payload).Username
// 	accounts, err := server.store.GetAccounts(ctx, db.GetAccountsParams{
// 		Owner: username,
// 		Limit: req.Limit,
// 		Offset: req.Offset,
// 	})
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, accounts)
// }

// // Delete Account
// type deleteAccountRequest struct {
// 	ID int64 `uri:"id" binding:"required,min=1"`
// }
// func (server *Server) deleteAccount(ctx *gin.Context) {
// 	var req deleteAccountRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	username := ctx.MustGet(AUTH_TOKEN).(*token.Payload).Username
// 	account, err := server.store.GetAccount(ctx, req.ID)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// 	if username != account.Owner {
// 		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
// 		return
// 	}

// 	err = server.store.DeleteAccount(ctx, req.ID)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, fmt.Sprintf("account %d deleted", req.ID))
// }