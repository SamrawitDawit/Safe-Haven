package interfaces

type EmailServiceInterface interface {
	SendResetPasswordEmail(email string, resetToken string) error
}
