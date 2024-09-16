package domain

import "net/http"

type CustomError struct {
	Message    string
	StatusCode int
}

// Error implements error.
func (c *CustomError) Error() string {
	return c.Message
}

func NewCustomError(message string, statusCode int) *CustomError {
	return &CustomError{
		Message:    message,
		StatusCode: statusCode,
	}
}

var (
	ErrUserNotFound           = NewCustomError("User not found", http.StatusNotFound)
	ErrUserEmailExists        = NewCustomError("Email already exists", 409)
	ErrUserPhoneNumberExists  = NewCustomError("Phone number already exists", 409)
	ErrUserAnonymousExists    = NewCustomError("Anonymous differentiator already exists", 409)
	ErrUserCreationFailed     = NewCustomError("Failed to create user", http.StatusInternalServerError)
	ErrUserUpdateFailed       = NewCustomError("Failed to update user", http.StatusInternalServerError)
	ErrUserPromotionFailed    = NewCustomError("Failed to promote user", http.StatusInternalServerError)
	ErrUserFetchFailed        = NewCustomError("Failed to fetch users", http.StatusInternalServerError)
	ErrUserCursorDecodeFailed = NewCustomError("Failed to decode user data", http.StatusInternalServerError)
	ErrUserCountFailed        = NewCustomError("Failed to count users", http.StatusInternalServerError)
)

// auth-related errors
var (
	ErrUserTokenUpdateFailed   = NewCustomError("Failed to update user tokens", http.StatusInternalServerError)
	ErrInvalidCredentials      = NewCustomError("Invalid email or password", http.StatusUnauthorized)
	ErrInvalidToken            = NewCustomError("Invalid token", http.StatusUnauthorized)
	ErrInvalidRefreshToken     = NewCustomError("Invalid refresh token", http.StatusUnauthorized)
	ErrInvalidResetCode        = NewCustomError("Invalid reset code", http.StatusBadRequest)
	ErrUnAuthorized            = NewCustomError("Unauthorized access", http.StatusUnauthorized)
	ErrUnexpectedSigningMethod = NewCustomError("Unexpected signing method", http.StatusInternalServerError)
)

// JWT-related errors
var (
	ErrTokenGenerationFailed        = NewCustomError("Failed to generate token", http.StatusInternalServerError)
	ErrRefreshTokenGenerationFailed = NewCustomError("Failed to generate refresh token", http.StatusInternalServerError)
	ErrTokenParsingFailed           = NewCustomError("Invalid token", http.StatusUnauthorized)
	ErrResetTokenGenerationFailed   = NewCustomError("Failed to generate reset token", http.StatusInternalServerError)
	ErrInvalidTokenClaims           = NewCustomError("Invalid token claims", http.StatusUnauthorized)
)
var (
	// Password-related errors
	ErrPasswordHashingFailed = NewCustomError("Failed to hash password", http.StatusInternalServerError)

	// Email-related errors
	ErrEmailSendingFailed = NewCustomError("Failed to send email", http.StatusInternalServerError)
	ErrInvalidPassword    = NewCustomError("Invalid password", http.StatusUnauthorized)
)
