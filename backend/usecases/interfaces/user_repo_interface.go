package interfaces

import (
	"backend/domain"

	"github.com/google/uuid"
)

type UserRepositoryInterface interface {
	CreateUser(user *domain.User) error
	UpdateUser(user *domain.User) error
	GetUserByEmail(email string) (*domain.User, error)
	GetUserByID(id uuid.UUID) (*domain.User, error)
	GetUserByPhoneNumber(phoneNumber string) (*domain.User, error)
	GetUserByAnonymousDifferentiator(differentiator string) (*domain.User, error)
	GetUsersCount() (int, error)
}
