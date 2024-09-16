package infrastructure

import (
	"backend/domain"

	"golang.org/x/crypto/bcrypt"
)

type HashingService struct{}

func (s *HashingService) HashPassword(password string) (string, *domain.CustomError) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", domain.ErrPasswordHashingFailed
	}
	return string(bytes), nil
}

func (s *HashingService) CheckPasswordHash(password, hash string) *domain.CustomError {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return domain.ErrInvalidPassword
	}
	return nil
}
