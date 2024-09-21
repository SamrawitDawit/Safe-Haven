package infrastructure

import (
	"backend/domain"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService struct {
	JwtSecret string
}

func (s *JWTService) GenerateToken(user *domain.User) (string, string, *domain.CustomError) {
	claims := jwt.MapClaims{
		"id":       user.ID,
		"role":     user.Role,
		"category": user.Category,
		"exp":      time.Now().Add(time.Minute * 5),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err := accessToken.SignedString([]byte(s.JwtSecret))
	if err != nil {
		return "", "", domain.ErrTokenGenerationFailed
	}

	refreshClaims := jwt.MapClaims{
		"id":  user.ID.String(),
		"exp": time.Now().Add(time.Hour * 24 * 7),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(s.JwtSecret))
	if err != nil {
		return "", "", domain.ErrRefreshTokenGenerationFailed
	}
	return accessTokenString, refreshTokenString, nil
}

func (s *JWTService) ValidateToken(token string) (*jwt.Token, *domain.CustomError) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, domain.ErrUnexpectedSigningMethod
		}
		return []byte(s.JwtSecret), nil
	})
	if err != nil {
		return nil, domain.ErrTokenParsingFailed
	}
	return parsedToken, nil
}

func (s *JWTService) ExtractTokenClaims(token *jwt.Token) (jwt.MapClaims, *domain.CustomError) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, domain.ErrTokenParsingFailed
	}
	return claims, nil
}

func (s *JWTService) GenerateResetToken(email string, code int64) (string, *domain.CustomError) {
	claims := jwt.MapClaims{
		"email": email,
		"code":  code,
		"exp":   time.Now().Add(time.Minute * 5).Unix(),
	}

	resetToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	resetTokenString, err := resetToken.SignedString([]byte(s.JwtSecret))
	if err != nil {
		return "", domain.ErrResetTokenGenerationFailed
	}
	return resetTokenString, nil
}
