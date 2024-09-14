package interfaces

import (
	"backend/domain"

	"github.com/dgrijalva/jwt-go"
)

type JwtServiceInterface interface {
	GenerateToken(user *domain.User) (string, string, error)
	ValidateToken(token string) (*jwt.Token, error)
	GenerateResetToken(email string, code int64) (string, error)
	ExtractTokenClaims(token *jwt.Token) (jwt.MapClaims, error)
}
