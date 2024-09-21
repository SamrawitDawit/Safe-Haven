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
	ErrUserNotFound          = NewCustomError("User not found", http.StatusNotFound)
	ErrUserEmailExists       = NewCustomError("Email already exists", 409)
	ErrUserPhoneNumberExists = NewCustomError("Phone number already exists", 409)
	ErrUserCreationFailed    = NewCustomError("Failed to create user", http.StatusInternalServerError)
	ErrUserUpdateFailed      = NewCustomError("Failed to update user", http.StatusInternalServerError)
	ErrUserFetchFailed       = NewCustomError("Failed to fetch users", http.StatusInternalServerError)
	ErrUserCountFailed       = NewCustomError("Failed to count users", http.StatusInternalServerError)
	ErrRaceCondition         = NewCustomError("Concurrency error", http.StatusInternalServerError)
)

// auth-related errors
var (
	ErrInvalidCredentials      = NewCustomError("Invalid email or password", http.StatusUnauthorized)
	ErrInvalidToken            = NewCustomError("Invalid token", http.StatusUnauthorized)
	ErrInvalidRefreshToken     = NewCustomError("Invalid refresh token", http.StatusUnauthorized)
	ErrInvalidResetCode        = NewCustomError("Invalid reset code", http.StatusBadRequest)
	ErrUnAuthorized            = NewCustomError("Unauthorized access", http.StatusUnauthorized)
	ErrUnexpectedSigningMethod = NewCustomError("Unexpected signing method", http.StatusInternalServerError)
	ErrEmailOrPhoneRequired    = NewCustomError("Email or phone number is required", http.StatusBadRequest)
)

// JWT-related errors
var (
	ErrTokenGenerationFailed        = NewCustomError("Failed to generate token", http.StatusInternalServerError)
	ErrRefreshTokenGenerationFailed = NewCustomError("Failed to generate refresh token", http.StatusInternalServerError)
	ErrTokenParsingFailed           = NewCustomError("Invalid token", http.StatusUnauthorized)
	ErrResetTokenGenerationFailed   = NewCustomError("Failed to generate reset token", http.StatusInternalServerError)
	ErrInvalidTokenClaims           = NewCustomError("Invalid token claims", http.StatusUnauthorized)
	ErrInvalidResetToken            = NewCustomError("Invalid reset token", http.StatusBadRequest)
)
var (
	ErrPasswordHashingFailed = NewCustomError("Failed to hash password", http.StatusInternalServerError)
	ErrEmailSendingFailed    = NewCustomError("Failed to send email", http.StatusInternalServerError)
	ErrInvalidPassword       = NewCustomError("Invalid password", http.StatusUnauthorized)
	ErrResetTokenAlreadySent = NewCustomError("Reset token already sent, Try again after 5 minutes", http.StatusBadRequest)
)

//Encryption-related errors
var (
	ErrEncryptionFailed = NewCustomError("Failed to encrypt data", http.StatusInternalServerError)
	ErrDecryptionFailed = NewCustomError("Failed to decrypt data", http.StatusInternalServerError)
)
