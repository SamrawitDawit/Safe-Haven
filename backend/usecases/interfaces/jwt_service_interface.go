package interfaces

import (
	"backend/domain"

	"github.com/dgrijalva/jwt-go"
)

type JwtServiceInterface interface {
	GenerateToken(user *domain.User) (string, string, *domain.CustomError)
	ValidateToken(token string) (*jwt.Token, *domain.CustomError)
	GenerateResetToken(email string, code int64) (string, *domain.CustomError)
	ExtractTokenClaims(token *jwt.Token) (jwt.MapClaims, *domain.CustomError)
}
