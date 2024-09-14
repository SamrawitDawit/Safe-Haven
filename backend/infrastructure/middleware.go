package infrastructure

import (
	"backend/usecases/interfaces"
	"backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(JwtService interfaces.JwtServiceInterface) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer context.Next()

		authHeader := context.GetHeader("Authorization")

		if authHeader == "" {
			res := utils.ErrorResponse(http.StatusUnauthorized, "No token provided", "No token provided")
			context.JSON(http.StatusUnauthorized, res)
			context.Abort()
			return
		}

		authPart := strings.Split(authHeader, " ")

		if len(authPart) != 2 || strings.ToLower(authPart[0]) != "bearer" {
			res := utils.ErrorResponse(http.StatusUnauthorized, "Invalid token", "Invalid token")
			context.JSON(http.StatusUnauthorized, res)
			context.Abort()
			return
		}

		token, err := JwtService.ValidateToken(authPart[1])

		if token == nil || !token.Valid {
			res := utils.ErrorResponse(http.StatusUnauthorized, "Invalid token", "Invalid token")
			context.JSON(http.StatusUnauthorized, res)
			context.Abort()
		}

		claims, err := JwtService.ExtractTokenClaims(token)
		if err != nil {
			res := utils.ErrorResponse(http.StatusUnauthorized, "Invalid token", "Invalid token")
			context.JSON(http.StatusUnauthorized, res)
			context.Abort()
		}
		role := claims["role"]
		id := claims["id"]

		if role == nil || id == nil {
			res := utils.ErrorResponse(http.StatusUnauthorized, "Invalid token", "Invalid token")
			context.JSON(http.StatusUnauthorized, res)
			context.Abort()
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
			res := utils.ErrorResponse(http.StatusForbidden, "Sorry, you must be an admin", "Sorry, you must be an admin")
			context.JSON(http.StatusForbidden, res)
			context.Abort()
		}
	}
}
func CounselorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer context.Next()
		role, exists := context.Get("role")
		if !exists || role != "counselor" {
			res := utils.ErrorResponse(http.StatusForbidden, "Sorry, you must be a counselor", "Sorry, you must be a counselor")
			context.JSON(http.StatusForbidden, res)
			context.Abort()
		}
	}
}
