package interfaces

import (
	"backend/domain"

	"github.com/google/uuid"
)

type UserRepositoryInterface interface {
	CreateUser(user *domain.User) *domain.CustomError
	UpdateUser(user *domain.User) *domain.CustomError
	GetUserByEmail(email string) (*domain.User, *domain.CustomError)
	GetUserByID(id uuid.UUID) (*domain.User, *domain.CustomError)
	GetUserByPhoneNumber(phoneNumber string) (*domain.User, *domain.CustomError)
	GetUserByAnonymousDifferentiator(differentiator string) (*domain.User, *domain.CustomError)
	GetUsersCount() (int, *domain.CustomError)
}
