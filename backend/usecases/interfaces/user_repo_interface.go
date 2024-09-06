package interfaces

import "backend/domain"

type UserRepositoryInterface interface {
	CreateNormalUser(user *domain.User) error
	GetUserByEmail(email string) (*domain.User, error)
	GetUserByPhoneNumber(phoneNumber string) (*domain.User, error)
}
