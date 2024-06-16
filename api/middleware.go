package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zzooman/zapp-server/token"
)

const (
	AUTH_TOKEN = "auth_token"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenInCookie, err := ctx.Cookie(AUTH_TOKEN)
		if tokenInCookie != "" {
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
		} else {
			// Parse the token from the Authorization header
			authHeader := ctx.GetHeader("Authorization")
			if authHeader == "" {
				ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("Authorization header is missing")))
				ctx.Abort()
				return
			}

			// Split the header value to get the token
			splitToken := strings.Split(authHeader, "Bearer ")
			if len(splitToken) != 2 {
				ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("Invalid token format")))
				ctx.Abort()
				return
			}

			// Get the token value
			tokenValue := splitToken[1]

			// Verify the token
			token, err := tokenMaker.VerifyToken(tokenValue)
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
}
