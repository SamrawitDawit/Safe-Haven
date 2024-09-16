package interfaces

import "backend/domain"

type HashingServiceInterface interface {
	HashPassword(password string) (string, *domain.CustomError)
	CheckPasswordHash(hashedPassword string, password string) *domain.CustomError
}
