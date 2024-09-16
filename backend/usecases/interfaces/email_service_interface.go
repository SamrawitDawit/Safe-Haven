package interfaces

import "backend/domain"

type EmailServiceInterface interface {
	SendResetPasswordEmail(email string, resetToken string) *domain.CustomError
}
