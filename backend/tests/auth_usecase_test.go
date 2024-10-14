package tests

import (
	"backend/domain"
	"backend/tests/mocks"
	"backend/usecases"
	"backend/usecases/dto"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AuthUseCaseTestSuite struct {
	suite.Suite
	authUseCase        usecases.AuthUseCaseInterface
	mockUserRepo       *mocks.UserRepositoryInterface
	mockJwtService     *mocks.JwtServiceInterface
	mockEmailService   *mocks.EmailServiceInterface
	mockHashingService *mocks.HashingServiceInterface
	mockEncryptService *mocks.EncryptionServiceInterface
}

func (suite *AuthUseCaseTestSuite) SetupTest() {
	suite.mockUserRepo = new(mocks.UserRepositoryInterface)
	suite.mockJwtService = new(mocks.JwtServiceInterface)
	suite.mockEmailService = new(mocks.EmailServiceInterface)
	suite.mockHashingService = new(mocks.HashingServiceInterface)
	suite.mockEncryptService = new(mocks.EncryptionServiceInterface)
	suite.authUseCase = usecases.NewAuthUseCase(
		suite.mockUserRepo,
		suite.mockJwtService,
		suite.mockEmailService,
		suite.mockHashingService,
		suite.mockEncryptService,
	)
}

func (suite *AuthUseCaseTestSuite) TestRegister_Success() {
	// Arrange
	registerDTO := dto.RegisterDTO{
		FullName: "Test User",
		Email:    "test@example.com",
		Password: "passworD123!",
		Category: "general",
		Language: "Amharic",
	}

	// Expect user repo to not find an existing user by email
	suite.mockUserRepo.On("GetUserByEmail", registerDTO.Email).Return(nil, domain.ErrUserNotFound)

	suite.mockUserRepo.On("GetUsersCount").Return(0, nil)

	// Expect password hashing to succeed
	suite.mockHashingService.On("Hash", registerDTO.Password).Return("hashed_password", nil)

	// Expect CreateUser to be called
	suite.mockUserRepo.On("CreateUser", mock.Anything).Return(nil)

	// Act
	user, err := suite.authUseCase.Register(registerDTO)

	// Assert
	suite.NotNil(user)
	suite.Nil(err)
	suite.mockUserRepo.AssertCalled(suite.T(), "CreateUser", mock.MatchedBy(func(user *domain.User) bool {
		return user.Role == "admin" && user.Verified == true
	}))
}

func (suite *AuthUseCaseTestSuite) TestRegister_Email_Exists_Failure() {
	// Arrange
	registerDTO := dto.RegisterDTO{
		FullName: "Test User",
		Email:    "test@example.com",
		Password: "passworD123!",
		Category: "general",
		Language: "Amharic",
	}

	// Expect user repo to not find an existing user by email
	suite.mockUserRepo.On("GetUserByEmail", registerDTO.Email).Return(&domain.User{GoogleSignin: false}, domain.ErrUserNotFound)

	// Act
	user, err := suite.authUseCase.Register(registerDTO)

	// Assert
	suite.Nil(user)
	suite.Equal(err, domain.ErrUserEmailExists)
}

func (suite *AuthUseCaseTestSuite) TestRegister_EmailExists_GoogleSignin_Success() {
	// Arrange
	registerDTO := dto.RegisterDTO{
		FullName: "Test User",
		Email:    "test@example.com",
		Password: "passworD123!",
		Category: "general",
		Language: "Amharic",
	}

	existingUser := &domain.User{
		Email:        registerDTO.Email,
		GoogleSignin: true,
	}

	// Mocking the situation where user exists with Google sign-in
	suite.mockUserRepo.On("GetUserByEmail", registerDTO.Email).Return(existingUser, nil)
	suite.mockUserRepo.On("GetUsersCount").Return(0, nil)
	suite.mockHashingService.On("Hash", registerDTO.Password).Return("hashed_password", nil)
	suite.mockUserRepo.On("UpdateUser", mock.Anything).Return(nil)

	// Act
	user, err := suite.authUseCase.Register(registerDTO)

	// Assert
	suite.NotNil(user)
	suite.Nil(err)
	suite.mockUserRepo.AssertCalled(suite.T(), "UpdateUser", mock.Anything)
}

func (suite *AuthUseCaseTestSuite) TestRegister_PhoneNumberExists_Failure() {
	// Arrange
	registerDTO := dto.RegisterDTO{
		FullName:    "Test User",
		PhoneNumber: "1234567890",
		Password:    "passworD123!",
		Category:    "general",
		Language:    "Amharic",
	}

	existingUser := &domain.User{
		PhoneNumber: registerDTO.PhoneNumber,
	}

	// Expect user repo to return an existing user by phone number
	suite.mockUserRepo.On("GetUserByPhoneNumber", registerDTO.PhoneNumber).Return(existingUser, nil)

	// Act
	user, err := suite.authUseCase.Register(registerDTO)

	// Assert
	suite.Nil(user)
	suite.Equal(err, domain.ErrUserPhoneNumberExists)
}

func (suite *AuthUseCaseTestSuite) TestLogin_Success() {
	// Arrange
	loginDTO := dto.LoginDTO{
		Email:    "test@example.com",
		Password: "password123",
	}

	user := &domain.User{
		Email:    loginDTO.Email,
		Password: "hashed_password",
	}

	suite.mockUserRepo.On("GetUserByEmail", loginDTO.Email).Return(user, nil)
	suite.mockHashingService.On("CheckHash", loginDTO.Password, user.Password).Return(nil)
	suite.mockUserRepo.On("UpdateUserFields", mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("map[string]interface {}")).Return(nil)
	suite.mockJwtService.On("GenerateToken", user).Return("access_token", "refresh_token", nil)
	suite.mockEncryptService.On("Encrypt", mock.Anything).Return("encrypted", nil)

	// Act
	user, accessToken, refreshToken, err := suite.authUseCase.Login(loginDTO)

	// Assert
	suite.Nil(err)
	suite.Equal("access_token", accessToken)
	suite.Equal("refresh_token", refreshToken)
	suite.NotNil(user)
}

func (suite *AuthUseCaseTestSuite) TestLogin_UserNotFoundByEmail_Failure() {
	// Arrange
	loginDTO := dto.LoginDTO{
		Email:    "notfound@example.com",
		Password: "password123",
	}

	// Mock user not found
	suite.mockUserRepo.On("GetUserByEmail", loginDTO.Email).Return(nil, domain.ErrUserNotFound)

	// Act
	_, _, _, err := suite.authUseCase.Login(loginDTO)

	// Assert
	suite.Equal(err, domain.ErrUserNotFound)
}

func (suite *AuthUseCaseTestSuite) TestLogin_InvalidPassword_Failure() {
	// Arrange
	loginDTO := dto.LoginDTO{
		Email:    "test@example.com",
		Password: "wrongpassword",
	}

	user := &domain.User{
		Email:    loginDTO.Email,
		Password: "hashed_password",
	}

	// Mock user found by email
	suite.mockUserRepo.On("GetUserByEmail", loginDTO.Email).Return(user, nil)
	// Mock password check failure
	suite.mockHashingService.On("CheckHash", loginDTO.Password, user.Password).Return(domain.ErrInvalidCredentials)

	// Act
	_, _, _, err := suite.authUseCase.Login(loginDTO)

	// Assert
	suite.Equal(err, domain.ErrInvalidCredentials)
}

func (suite *AuthUseCaseTestSuite) TestLogin_GoogleSigninAttemptedNormalLogin_Failure() {
	// Arrange
	loginDTO := dto.LoginDTO{
		Email:    "googleuser@example.com",
		Password: "password123",
	}

	user := &domain.User{
		Email:        loginDTO.Email,
		GoogleSignin: true,
	}

	// Mock user found by email
	suite.mockUserRepo.On("GetUserByEmail", loginDTO.Email).Return(user, nil)

	// Act
	_, _, _, err := suite.authUseCase.Login(loginDTO)

	// Assert
	suite.Equal(err, domain.ErrInvalidCredentials)
}

func (suite *AuthUseCaseTestSuite) TestForgotPassword_Success() {
	// Arrange
	email := "test@example.com"
	user := &domain.User{
		Email: email,
	}

	suite.mockUserRepo.On("GetUserByEmail", email).Return(user, nil)
	suite.mockJwtService.On("GenerateResetToken", user.Email, mock.Anything).Return("reset_token", nil)
	suite.mockEmailService.On("SendResetPasswordEmail", user.Email, "reset_token").Return(nil)
	suite.mockUserRepo.On("UpdateUserFields", mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("map[string]interface {}")).Return(nil)
	suite.mockEncryptService.On("Encrypt", mock.Anything).Return("encrypted", nil)
	suite.mockHashingService.On("Hash", mock.Anything).Return("hashed", nil)

	// Act
	err := suite.authUseCase.ForgotPassword(email)

	// Assert
	suite.Nil(err)
	suite.mockUserRepo.AssertCalled(suite.T(), "UpdateUserFields", mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("map[string]interface {}"))
}

func (suite *AuthUseCaseTestSuite) TestForgotPassword_UserNotFound_Failure() {
	// Arrange
	email := "notfound@example.com"

	// Mock user not found
	suite.mockUserRepo.On("GetUserByEmail", email).Return(nil, domain.ErrUserNotFound)

	// Act
	err := suite.authUseCase.ForgotPassword(email)

	// Assert
	suite.Equal(err, domain.ErrUserNotFound)
}
func (suite *AuthUseCaseTestSuite) TestRefreshToken_Success() {
	// Arrange
	refreshToken := "valid_refresh_token"
	user := &domain.User{
		ID:           uuid.New(),
		RefreshToken: "valid_refresh_token",
	}

	// Mock token validation and user retrieval
	validToken := &jwt.Token{Valid: true}
	suite.mockJwtService.On("ValidateToken", refreshToken).Return(validToken, nil)
	suite.mockJwtService.On("ExtractTokenClaims", validToken).Return(jwt.MapClaims{"id": user.ID.String()}, nil)
	suite.mockUserRepo.On("GetUserByID", user.ID).Return(user, nil)
	suite.mockJwtService.On("ValidateToken", user.RefreshToken).Return(validToken, nil)
	suite.mockJwtService.On("GenerateToken", user).Return("access_token", "refresh_token", nil)
	suite.mockUserRepo.On("UpdateUserFields", mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("map[string]interface {}")).Return(nil)
	suite.mockUserRepo.On("GetUserByIDWithLock", user.ID).Return(user, nil)
	suite.mockEncryptService.On("Decrypt", "valid_refresh_token").Return("valid_refresh_token", nil)
	suite.mockEncryptService.On("Encrypt", mock.AnythingOfType("string")).Return("new_refresh_token", nil)

	// Act
	accToken, refToken, err := suite.authUseCase.RefreshToken(refreshToken)

	// Assert
	suite.Nil(err)
	suite.Equal("access_token", accToken)
	suite.Equal("refresh_token", refToken)
	suite.mockUserRepo.AssertCalled(suite.T(), "UpdateUserFields", mock.Anything, mock.Anything, mock.Anything)
}
func (suite *AuthUseCaseTestSuite) TestRefreshToken_InvalidToken_Failure() {
	// Arrange
	invalidToken := "invalid_refresh_token"

	// Mock invalid token validation
	suite.mockJwtService.On("ValidateToken", invalidToken).Return(nil, domain.ErrInvalidToken)

	// Act
	_, _, err := suite.authUseCase.RefreshToken(invalidToken)

	// Assert
	suite.Equal(err, domain.ErrInvalidToken)
}
func (suite *AuthUseCaseTestSuite) TestRefreshToken_TokenMismatch_Failure() {
	// Arrange
	refreshToken := "valid_refresh_token"
	user := &domain.User{
		ID:           uuid.New(),
		RefreshToken: "different_token",
	}

	// Mock token validation and user retrieval
	validToken := &jwt.Token{Valid: true}
	suite.mockJwtService.On("ValidateToken", refreshToken).Return(validToken, nil)
	suite.mockJwtService.On("ExtractTokenClaims", validToken).Return(jwt.MapClaims{"id": user.ID.String()}, nil)
	suite.mockUserRepo.On("GetUserByID", user.ID).Return(user, nil)
	suite.mockJwtService.On("ValidateToken", user.RefreshToken).Return(validToken, nil)
	suite.mockUserRepo.On("GetUserByIDWithLock", user.ID).Return(user, nil)
	suite.mockUserRepo.On("UpdateUserFields", mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("map[string]interface {}")).Return(nil)
	suite.mockEncryptService.On("Decrypt", "different_token").Return("invalid_refresh_token", nil)

	// Act
	_, _, err := suite.authUseCase.RefreshToken(refreshToken)

	// Assert
	suite.Equal(err, domain.ErrInvalidRefreshToken)
}
func (suite *AuthUseCaseTestSuite) TestResetPassword_Success() {
	// Arrange
	token := "valid_token"
	newPassword := "new_passworD!1"
	user := &domain.User{
		Email:      "test@example.com",
		ResetCode:  "12345",
		ResetToken: token,
		Password:   "old_password",
	}
	suite.mockEncryptService.On("Decrypt", "1234").Return(1234, nil)
	suite.mockEncryptService.On("Decrypt", "valid_token").Return("valid_token", nil)
	claims := jwt.MapClaims{
		"code":  1234.5, // claims code as float64
		"email": user.Email,
	}

	validToken := &jwt.Token{Valid: true}
	suite.mockJwtService.On("ValidateToken", token).Return(validToken, nil)
	suite.mockJwtService.On("ExtractTokenClaims", validToken).Return(claims, nil)
	suite.mockUserRepo.On("GetUserByEmail", user.Email).Return(user, nil)
	suite.mockHashingService.On("Hash", newPassword).Return("hashed_new_password", nil)
	suite.mockHashingService.On("CheckHash", mock.Anything, mock.Anything).Return(nil)
	suite.mockEncryptService.On("Decrypt", "1234").Return(1234, nil)
	suite.mockEncryptService.On("Decrypt", "valid_token").Return("valid_token", nil)
	suite.mockUserRepo.On("UpdateUserFields", mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("map[string]interface {}")).Return(nil)

	// Act
	err := suite.authUseCase.ResetPassword(token, newPassword)

	// Assert
	suite.Nil(err)

}

func (suite *AuthUseCaseTestSuite) TestResetPassword_InvalidToken() {
	// Arrange
	token := "invalid_token"
	newPassword := "new_passworD1!"

	suite.mockJwtService.On("ValidateToken", token).Return(nil, domain.ErrInvalidToken)

	// Act
	err := suite.authUseCase.ResetPassword(token, newPassword)

	// Assert
	suite.Equal(err, domain.ErrInvalidToken)
}
func (suite *AuthUseCaseTestSuite) TestResetPassword_InvalidTokenClaims() {
	// Arrange
	token := "valid_token"
	newPassword := "new_passworD1!"

	validToken := &jwt.Token{Valid: true}
	suite.mockJwtService.On("ValidateToken", token).Return(validToken, nil)
	suite.mockJwtService.On("ExtractTokenClaims", validToken).Return(nil, domain.ErrInvalidTokenClaims)

	// Act
	err := suite.authUseCase.ResetPassword(token, newPassword)

	// Assert
	suite.Equal(err, domain.ErrInvalidTokenClaims)
}
func (suite *AuthUseCaseTestSuite) TestResetPassword_UserNotFound() {
	// Arrange
	token := "valid_token"
	newPassword := "new_passworD1!"

	claims := jwt.MapClaims{
		"code":  123.45,
		"email": "test@example.com",
	}

	validToken := &jwt.Token{Valid: true}
	suite.mockJwtService.On("ValidateToken", token).Return(validToken, nil)
	suite.mockJwtService.On("ExtractTokenClaims", validToken).Return(claims, nil)
	suite.mockUserRepo.On("GetUserByEmail", "test@example.com").Return(nil, domain.ErrUserNotFound)

	// Act
	err := suite.authUseCase.ResetPassword(token, newPassword)

	// Assert
	suite.Equal(err, domain.ErrUserNotFound)
}
func (suite *AuthUseCaseTestSuite) TestResetPassword_CodeMismatch() {
	// Arrange
	token := "valid_token"
	newPassword := "new_passworD1!"
	user := &domain.User{
		Email:      "test@example.com",
		ResetCode:  "54321", // Different reset code
		ResetToken: token,
	}

	claims := jwt.MapClaims{
		"code":  123.45, // Code in token is different
		"email": user.Email,
	}

	validToken := &jwt.Token{Valid: true}
	suite.mockJwtService.On("ValidateToken", token).Return(validToken, nil)
	suite.mockJwtService.On("ExtractTokenClaims", validToken).Return(claims, nil)
	suite.mockEncryptService.On("Decrypt", "valid_token").Return("valid_token", nil)
	suite.mockHashingService.On("CheckHash", mock.Anything, mock.Anything).Return(domain.ErrInvalidResetCode)
	suite.mockUserRepo.On("GetUserByEmail", user.Email).Return(user, nil)

	// Act
	err := suite.authUseCase.ResetPassword(token, newPassword)

	// Assert
	suite.Equal(err, domain.ErrInvalidResetCode)
}
func (suite *AuthUseCaseTestSuite) TestHandleGoogleCallback_UserExists_NonGoogleSignin_Failure() {
	// Arrange
	user := &domain.User{
		Email:        "test@example.com",
		GoogleSignin: false, // User exists but not signed in with Google
	}

	suite.mockUserRepo.On("GetUserByEmail", user.Email).Return(user, nil)

	// Act
	_, _, _, err := suite.authUseCase.HandleGoogleCallback(user)

	// Assert
	suite.Equal(err, domain.ErrUserEmailExists)
}
func (suite *AuthUseCaseTestSuite) TestHandleGoogleCallback_UserExists_GoogleSignin_Success() {
	// Arrange
	user := &domain.User{
		Email:        "test@example.com",
		GoogleSignin: true,
	}

	suite.mockUserRepo.On("GetUserByEmail", user.Email).Return(user, nil)
	suite.mockJwtService.On("GenerateToken", user).Return("access_token", "refresh_token", nil)
	suite.mockUserRepo.On("UpdateUser", user).Return(nil)
	suite.mockUserRepo.On("UpdateUserFields", mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("map[string]interface {}")).Return(nil)
	suite.mockEncryptService.On("Encrypt", mock.Anything).Return("encrypted", nil)

	// Act
	returned_user, accessToken, refreshToken, err := suite.authUseCase.HandleGoogleCallback(user)

	// Assert
	suite.Nil(err)
	suite.NotNil(returned_user)
	suite.Equal("access_token", accessToken)
	suite.Equal("refresh_token", refreshToken)
}
func (suite *AuthUseCaseTestSuite) TestHandleGoogleCallback_NewUser_Success() {
	// Arrange
	user := &domain.User{
		Email:        "newuser@example.com",
		GoogleSignin: true,
	}

	suite.mockUserRepo.On("GetUserByEmail", user.Email).Return(nil, domain.ErrUserNotFound)
	suite.mockJwtService.On("GenerateToken", user).Return("access_token", "refresh_token", nil)
	suite.mockUserRepo.On("CreateUser", mock.Anything).Return(nil)
	suite.mockEncryptService.On("Encrypt", mock.Anything).Return("encrypted", nil)
	suite.mockUserRepo.On("UpdateUserFields", mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("map[string]interface {}")).Return(nil)

	// Act
	returned_user, accessToken, refreshToken, err := suite.authUseCase.HandleGoogleCallback(user)

	// Assert
	suite.Nil(err)
	suite.NotNil(returned_user)
	suite.Equal("access_token", accessToken)
	suite.Equal("refresh_token", refreshToken)
}
func (suite *AuthUseCaseTestSuite) TestHandleGoogleCallback_GenerateTokenFailure() {
	// Arrange
	user := &domain.User{
		Email:        "newuser@example.com",
		GoogleSignin: true,
	}

	suite.mockUserRepo.On("GetUserByEmail", user.Email).Return(nil, domain.ErrUserNotFound)
	suite.mockJwtService.On("GenerateToken", user).Return("", "", domain.ErrTokenGenerationFailed)

	// Act
	_, _, _, err := suite.authUseCase.HandleGoogleCallback(user)

	// Assert
	suite.Equal(err, domain.ErrTokenGenerationFailed)
}

func (suite *AuthUseCaseTestSuite) TestHandleGoogleCallback_Success() {
	// Arrange
	user := &domain.User{
		Email: "test@example.com",
	}

	suite.mockUserRepo.On("GetUserByEmail", user.Email).Return(nil, domain.ErrUserNotFound)
	suite.mockJwtService.On("GenerateToken", user).Return("access_token", "refresh_token", nil)
	suite.mockEncryptService.On("Encrypt", mock.Anything).Return("encrypted", nil)
	suite.mockUserRepo.On("UpdateUserFields", mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("map[string]interface {}")).Return(nil)

	suite.mockUserRepo.On("CreateUser", mock.Anything).Return(nil)

	// Act
	returned_user, accessToken, refreshToken, err := suite.authUseCase.HandleGoogleCallback(user)

	// Assert
	suite.Nil(err)
	suite.NotNil(returned_user)
	suite.Equal("access_token", accessToken)
	suite.Equal("refresh_token", refreshToken)
}
func TestAuthUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(AuthUseCaseTestSuite))
}
