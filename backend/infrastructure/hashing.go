package infrastructure

import (
	"backend/domain"

	"golang.org/x/crypto/bcrypt"
)

type HashingService struct{}

func (s *HashingService) Hash(value string) (string, *domain.CustomError) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", domain.ErrPasswordHashingFailed
	}
	return string(bytes), nil
}

func (s *HashingService) CheckHash(value, hash string) *domain.CustomError {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))
	if err != nil {
		return domain.ErrInvalidPassword
	}
	return nil
}
