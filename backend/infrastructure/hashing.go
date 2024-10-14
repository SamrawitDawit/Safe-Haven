package infrastructure

import (
	"backend/domain"
	"backend/utils"

	"golang.org/x/crypto/bcrypt"
)

type HashingService struct{}

func (s *HashingService) Hash(value string) (string, *domain.CustomError) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		utils.LogError("Error hashing password", err)
		return "", domain.ErrPasswordHashingFailed
	}
	return string(bytes), nil
}

func (s *HashingService) CheckHash(value, hash string) *domain.CustomError {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))
	if err != nil {
		utils.LogError("Error checking hash", err)
		return domain.ErrInvalidPassword
	}
	return nil
}
