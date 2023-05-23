package middleware

import (
	"net/http"
	"rakamin-academy/sdk/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func IsAuthenticated(ctx *gin.Context) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		header := ctx.Request.Header.Get("Authorization")

		if header == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You must be logged in first."})
			return
		}

		tokenParts := strings.SplitN(header, " ", 2)

		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token."})
			return
		}

		tokenString := tokenParts[1]

		tokenClaims, err := jwt.DecodeToken(tokenString)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid JWT token"})
		}

		userID := tokenClaims["id"].(uint)

		ctx.Set("user", userID)

		ctx.Next()

	}

}
