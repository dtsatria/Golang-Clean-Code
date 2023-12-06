package middleware

import (
	"final-project-booking-room/config"
	"final-project-booking-room/utils/common"
	"net/http"

	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	RequireToken(roles ...string) gin.HandlerFunc
}

type authMiddleware struct {
	jwtService common.JwtToken
}

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func (a *authMiddleware) RequireToken(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var aH authHeader
		if err := ctx.ShouldBindHeader(&aH); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		tokenString := strings.Replace(aH.AuthorizationHeader, "Bearer ", "", -1)
		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		claims, err := a.jwtService.VerifyToken(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		ctx.Set(config.UserSesion, claims["userId"])
		ctx.Set(config.RoleSesion, claims["role"])
		// fmt.Println("claims :", claims)

		var validRole bool
		for _, role := range roles {
			if role == claims["role"] {
				validRole = true
				break
			}
		}

		if !validRole {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden Resourece"})
			return
		}

		ctx.Next()
	}
}

func NewAuthMiddleware(jwtService common.JwtToken) AuthMiddleware {
	return &authMiddleware{jwtService: jwtService}
}
