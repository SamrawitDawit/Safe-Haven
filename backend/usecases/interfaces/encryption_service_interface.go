package interfaces

import "backend/domain"

type EncryptionServiceInterface interface {
	Encrypt(value string) (string, *domain.CustomError)
	Decrypt(value string) (string, *domain.CustomError)
}
