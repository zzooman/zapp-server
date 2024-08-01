package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/zzooman/zapp-server/db/sqlc"
)

type SearchProductsRequest struct {
	Query  string `form:"query" binding:"required"`
	Limit  int32  `form:"limit" binding:"required"`
	Page   int32  `form:"page" binding:"required"`
}
type SearchProductsResponse struct {
	Products 	[]db.GetProductsWithSellerByQueryRow 	`json:"products"`
	Next  	 	bool							  		`json:"next"`
	Query  	 	string 									`json:"keyword"`
}
func (server *Server) searchProducts(ctx *gin.Context) {
	var req SearchProductsRequest
	var res SearchProductsResponse

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}		

	result, err := server.store.SearchProductsTx(ctx, db.SearchProductsParams{
		Query:  req.Query,
		Limit:  req.Limit,
		Offset: (req.Page - 1) * req.Limit,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	nextProducts, err := server.store.SearchProductsTx(ctx, db.SearchProductsParams{
		Query:  req.Query,
		Limit:  req.Limit,
		Offset: req.Page * req.Limit,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if len(nextProducts.Products) == 0 {
		res.Next = false
	} else {
		res.Next = true
	}

	res.Products = result.Products
	ctx.JSON(http.StatusOK, res)
}

func (server *Server) hotSearchTexts(ctx *gin.Context) {
	searchTexts, err := server.store.HotSearchTexts(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, searchTexts)
}
