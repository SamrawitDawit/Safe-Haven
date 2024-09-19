package infrastructure

import (
	"backend/usecases/interfaces"
	"backend/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(JwtService interfaces.JwtServiceInterface) gin.HandlerFunc {
	return func(context *gin.Context) {
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
		if err != nil || token == nil || !token.Valid {
			res := utils.ErrorResponse(http.StatusUnauthorized, "Invalid token", "Invalid token")
			context.JSON(http.StatusUnauthorized, res)
			context.Abort()
			return
		}

		claims, err := JwtService.ExtractTokenClaims(token)
		if err != nil || claims == nil {
			res := utils.ErrorResponse(http.StatusUnauthorized, "Invalid token", "Invalid token")
			context.JSON(http.StatusUnauthorized, res)
			context.Abort()
			return
		}

		role, roleOk := claims["role"].(string)
		id, idOk := claims["id"].(string)

		if !roleOk || !idOk {
			res := utils.ErrorResponse(http.StatusUnauthorized, "Invalid token", "Invalid token")
			context.JSON(http.StatusUnauthorized, res)
			context.Abort()
			return
		}

		context.Set("role", role)
		context.Set("id", id)

		// Log the role for debugging
		fmt.Println("role:", role)

		context.Next()
	}
}

func AdminMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		role, exists := context.Get("role")
		if !exists {
			res := utils.ErrorResponse(http.StatusUnauthorized, "Invalid token", "Invalid token")
			context.JSON(http.StatusUnauthorized, res)
			context.Abort()
		} else if role != "admin" {
			res := utils.ErrorResponse(http.StatusForbidden, "Sorry, you must be an admin", "Sorry, you must be an admin")
			context.JSON(http.StatusForbidden, res)
			context.Abort()
		}
		context.Next()
	}
}
func CounselorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		role, exists := context.Get("role")
		if !exists || role != "counselor" {
			res := utils.ErrorResponse(http.StatusForbidden, "Sorry, you must be a counselor", "Sorry, you must be a counselor")
			context.JSON(http.StatusForbidden, res)
			context.Abort()
		}
		defer context.Next()
	}
}
