package infrastructure

import (
	"backend/domain"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService struct {
	JwtSecret string
}

func (s *JWTService) GenerateTokenForUser(user *domain.User) (string, string, error) {
	claims := jwt.MapClaims{
		"id":       user.ID,
		"role":     user.Role,
		"category": user.Category,
		"exp":      time.Now().Add(time.Minute * 5).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err := accessToken.SignedString([]byte(s.JwtSecret))
	if err != nil {
		return "", "", err
	}

	refreshClaims := jwt.MapClaims{
		"id":          user.ID,
		"isAnonymous": false,
		"exp":         time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(s.JwtSecret))
	if err != nil {
		return "", "", err
	}
	return accessTokenString, refreshTokenString, nil
}

func (s *JWTService) GenerateTokenForAnonymousUser(anonUser *domain.AnonymousUser) (string, string, error) {
	claims := jwt.MapClaims{
		"id":             anonUser.ID,
		"differentiator": anonUser.AnonymousDifferentiator,
		"exp":            time.Now().Add(time.Minute * 5).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err := accessToken.SignedString([]byte(s.JwtSecret))
	if err != nil {
		return "", "", err
	}

	refreshClaims := jwt.MapClaims{
		"id":          anonUser.ID,
		"isAnonymous": true,
		"exp":         time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(s.JwtSecret))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func (s *JWTService) ValidateToken(token string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token")
		}
		return []byte(s.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return parsedToken, nil
}

func (s *JWTService) ExtractTokenClaims(token *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}
