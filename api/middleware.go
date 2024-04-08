package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zzooman/zapp-server/token"
)

const (
  AUTH_TOKEN = "auth_token"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenInCookie, err := ctx.Cookie(AUTH_TOKEN)
		if err != nil { 
			ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("auth_token is empty")))
			ctx.Abort()
			return
		}		
		token, err := tokenMaker.VerifyToken(tokenInCookie)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
			ctx.Abort()
			return
		}

		ctx.Set(AUTH_TOKEN, token)
		ctx.Next()
	}
}

// const (
// 	AUTH_HEADER_KEY  ="Authorization"
// 	AUTH_TYPE_BEARER = "Bearer"
// 	AUTH_PAYLOAD_KEY = "auth_payload"
// )

// func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		if ctx.GetHeader(AUTH_HEADER_KEY) == "" {
// 			ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("Authorization header is required")))
// 			ctx.Abort()
// 			return
// 		}

// 		authrizationHeader := ctx.GetHeader(AUTH_HEADER_KEY)
// 		fullToken := strings.Fields(authrizationHeader)
// 		if(len(fullToken) != 2 || fullToken[0] != AUTH_TYPE_BEARER) {
// 			ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("Invalid token")))
// 			ctx.Abort()
// 			return
// 		}
// 		payload, err := tokenMaker.VerifyToken(fullToken[1])
// 		if err != nil {
// 			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
// 			ctx.Abort()
// 			return
// 		}

// 		ctx.Set(AUTH_PAYLOAD_KEY, payload)
// 		ctx.Next()
// 	}
// }
		
