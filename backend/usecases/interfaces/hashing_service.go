package interfaces

import "backend/domain"

type HashingServiceInterface interface {
	Hash(value string) (string, *domain.CustomError)
	CheckHash(hash string, value string) *domain.CustomError
}
