package infrastructure

import (
	"backend/usecases/interfaces"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(JwtService interfaces.JwtServiceInterface) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer context.Next()

		authHeader := context.GetHeader("Authorization")

		if authHeader == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header is required"})
			context.Abort()
			return
		}

		authPart := strings.Split(authHeader, " ")

		if len(authPart) != 2 || strings.ToLower(authPart[0]) != "bearer" {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Authorization header"})
			context.Abort()
			return
		}

		token, err := JwtService.CheckToken(authPart[1])

		if token == nil || !token.Valid {
			errMsg := "Invalid or expired token"

			if err != nil {
				errMsg = err.Error()
			}
			context.JSON(http.StatusUnauthorized, gin.H{"error": errMsg})
			context.Abort()
			return
		}

		claims, ok := JwtService.FindClaim(token)
		if !ok {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			context.Abort()
			return
		}
		role := claims["role"]
		id := claims["id"]

		if role == nil || id == nil {
			context.JSON(401, gin.H{"error": "Invalid token claims"})
			context.Abort()
			return
		}
		context.Set("role", role)
		context.Set("id", id)

	}
}

func AdminMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer context.Next()
		role, exists := context.Get("role")
		if !exists || role != "admin" {
			context.JSON(http.StatusForbidden, gin.H{"message": "Sorry, you must be an admin"})
			context.Abort()
			return
		}
	}
}
func CounselorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer context.Next()
		role, exists := context.Get("role")
		if !exists || role != "counselor" {
			context.JSON(http.StatusForbidden, gin.H{"message": "Sorry, you must be a counselor"})
			context.Abort()
			return
		}
	}
}
