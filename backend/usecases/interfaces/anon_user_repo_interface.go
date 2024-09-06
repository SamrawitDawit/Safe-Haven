package interfaces

import "backend/domain"

type AnonymousUserRepoInterface interface {
	CreateAnonymousUser(anonUser *domain.AnonymousUser) error
	GetUserByAnonymousDifferentiator(differentiator string) (*domain.AnonymousUser, error)
}