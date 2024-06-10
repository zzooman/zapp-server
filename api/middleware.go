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
			// Remove the auth_token cookie
			ctx.SetCookie(AUTH_TOKEN, "", -1, "", "", false, true)
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
			ctx.Abort()			
			return
		}

		ctx.Set(AUTH_TOKEN, token)
		ctx.Next()
	}
}
