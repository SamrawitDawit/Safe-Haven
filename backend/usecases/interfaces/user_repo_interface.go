package interfaces

import "backend/domain"

type UserRepositoryInterface interface {
	CreateUser(user *domain.User) error
	GetUserByEmail(email string) (*domain.User, error)
	GetUserByPhoneNumber(phoneNumber string) (*domain.User, error)
	GetUserByAnonymousDifferentiator(differentiator string) (*domain.User, error)
}
