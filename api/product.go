package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/zzooman/zapp-server/db/sqlc"
	"github.com/zzooman/zapp-server/token"
)

type createProductRequest struct {
	Title   string   `json:"title" binding:"required"`
	Content string   `json:"content" binding:"required"`
	Price   string   `json:"price" binding:"required"`	
	Medias  []string `json:"medias" binding:"required"`
}

func (server *Server) createProduct(ctx *gin.Context) {
	var req createProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(AUTH_TOKEN).(*token.Payload)
	price, _ := strconv.ParseInt(req.Price, 10, 64)
	product, err := server.store.CreateProduct(ctx, db.CreateProductParams{
		Seller:    authPayload.Username,
		Title:     req.Title,
		Content:   req.Content,
		Price:     price,		
		Medias:     req.Medias,
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, product)
}

type Seller struct {
	Username string      `json:"username"`
	Email    pgtype.Text `json:"email"`
	Phone    pgtype.Text `json:"phone"`
	Profile  pgtype.Text `json:"profile"`
}
type ProductResponse struct {
	ID        int64               `json:"id"`
	Title     string              `json:"title"`
	Content   string              `json:"content"`
	Medias    []string            `json:"medias"`
	Price     int64               `json:"price"`	
	Views     pgtype.Int8         `json:"views"`
	CreatedAt pgtype.Timestamptz  `json:"created_at"`
	Seller    Seller              `json:"seller"`
	IsLiked   bool                `json:"isLiked"`
}

type getProductRequest struct {
	Id string `uri:"id" binding:"required"`	
}
func (server *Server) getProduct(ctx *gin.Context) {
	var req getProductRequest	
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id, err := strconv.ParseInt(req.Id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	productWithSeller, err := server.store.GetProductWithSellor(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.ViewProduct(ctx, productWithSeller.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	_, err = server.store.GetWishWithProduct(ctx, db.GetWishWithProductParams{
		ProductID:  productWithSeller.ID,
		Username: 	productWithSeller.Seller,
	})
	
	ctx.JSON(http.StatusOK, ProductResponse{
		ID:        	productWithSeller.ID,
		Title:     	productWithSeller.Title,
		Content:   	productWithSeller.Content,
		Medias:    	productWithSeller.Medias,
		Price:     	productWithSeller.Price,		
		Views:     	productWithSeller.Views,
		CreatedAt: 	productWithSeller.CreatedAt,
		Seller: Seller{
			Username: productWithSeller.Seller,
			Email:    productWithSeller.Email,
			Phone:    productWithSeller.Phone,
			Profile:  productWithSeller.Profile,			
		},
		IsLiked: err == nil,
	})
}

type getProductsRequest struct {
	Limit  	int32 	`form:"limit" binding:"required`
	Page 	int32 	`form:"page" binding:"required`
}
type getProductsResponse struct {
	Products []ProductResponse 	`json:"products"`
	Next  bool		   		`json:"next"`
}

func (server *Server) getProducts(ctx *gin.Context) {
	var req getProductsRequest
	var res getProductsResponse

	if err := ctx.BindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	// 게시글 & 작성자 정보 조회
	productsWithSeller, err := server.store.GetProductsWithSeller(ctx, db.GetProductsWithSellerParams{
		Limit:  req.Limit,
		Offset: (req.Page - 1) * req.Limit,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// 다음 페이지 존재 여부 확인
	nextProducts, err := server.store.GetProducts(ctx, db.GetProductsParams{
		Limit:  req.Limit,
		Offset: req.Page * req.Limit,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if len(nextProducts) == 0 {
		res.Next = false
	} else {
		res.Next = true	
	}

	// 채널 생성
	ch := make(chan struct {
		db.GetProductsWithSellerRow
		IsLiked bool
	}, len(productsWithSeller))

	var username string = ""
	auth, ok := ctx.Get(AUTH_TOKEN)
	if ok {
		username = auth.(*token.Payload).Username
	}

	// 비동기 처리
	for _, product := range productsWithSeller {
		go func(product db.GetProductsWithSellerRow) {			
			_, err := server.store.GetWishWithProduct(ctx, db.GetWishWithProductParams{
				ProductID:   product.ID,
				Username: username,
			})	
			ch <- struct {
				db.GetProductsWithSellerRow
				IsLiked bool
			}{product, err == nil}
		}(product)
	}

	// 결과 수집
	for range productsWithSeller {
		result := <-ch
		res.Products = append(res.Products, ProductResponse{
			ID:        	result.ID,
			Title:     	result.Title,
			Content:   	result.Content,
			Medias:     result.Medias,
			Price:     	result.Price,			
			Views:     	result.Views,
			CreatedAt: 	result.CreatedAt,
			Seller: Seller{
				Username: result.Seller,
				Email:    result.Email,
				Phone:    result.Phone,
				Profile:  result.Profile,
			},
			IsLiked: result.IsLiked,
		})
	}

	products := res.Products
	// 최신순 정렬
	for range products {
		for i := 0; i < len(products)-1; i++ {
			if products[i].CreatedAt.Time.Before(products[i+1].CreatedAt.Time) {				
				products[i], products[i+1] = products[i+1], products[i]
			}
		}
	}	
	ctx.JSON(http.StatusOK, res)	
}



func (server *Server) getProductsILiked(ctx *gin.Context) {
	var req getProductsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	username := ctx.MustGet(AUTH_TOKEN).(*token.Payload).Username
	products, err := server.store.GetProductsWithSellerThatILiked(ctx, db.GetProductsWithSellerThatILikedParams{
		Username: 	username,
		Limit:    	req.Limit,
		Offset:  	(req.Page - 1) * req.Limit,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, products)
}


func (server *Server) getProductsISold(ctx *gin.Context) {
	var req getProductsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	username := ctx.MustGet(AUTH_TOKEN).(*token.Payload).Username
	products, err := server.store.GetProductsWithSellerThatISold(ctx, db.GetProductsWithSellerThatISoldParams{
		Seller: 	username,
		Limit:    	req.Limit,
		Offset:  	(req.Page - 1) * req.Limit,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (server *Server) getProductsIBought(ctx *gin.Context) {
	var req getProductsRequest
	
	if err := ctx.ShouldBindQuery(&req); err != nil {
		fmt.Println("error 1", err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	username := ctx.MustGet(AUTH_TOKEN).(*token.Payload).Username
	products, err := server.store.GetProductsWithSellerThatIBought(ctx, db.GetProductsWithSellerThatIBoughtParams{		
		Buyer: 		username,
		Limit:    	req.Limit,
		Offset:  	(req.Page - 1) * req.Limit,
	})
	if err != nil {
		fmt.Println("error 2", err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}	
	ctx.JSON(http.StatusOK, products)
}