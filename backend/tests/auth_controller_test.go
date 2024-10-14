package tests

import (
	"backend/delivery/controllers"
	"backend/domain"
	"backend/tests/mocks"
	"backend/usecases/dto"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	// "cloud.google.com/go/auth/credentials/idtoken"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	// "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AuthControllerTestSuite struct {
	suite.Suite
	mockUsecase  *mocks.AuthUseCaseInterface
	controller   *controllers.AuthController
	recorder     *httptest.ResponseRecorder
}

func (suite *AuthControllerTestSuite) SetupTest() {
	suite.mockUsecase = new(mocks.AuthUseCaseInterface)
	suite.recorder = httptest.NewRecorder()
	suite.controller = controllers.NewAuthController(suite.mockUsecase)

	// Set Gin mode to TestMode to avoid debug output during tests
	gin.SetMode(gin.TestMode)
}

func (suite *AuthControllerTestSuite) TestRegister_Success() {
	registerDTO := dto.RegisterDTO{
		Email:    "test@example.com",
		Password: "password123",
		FullName: "John Doe",
		Category: "general",
		Language: "Amharic",
	}
	user := &domain.User{}
	suite.mockUsecase.On("Register", registerDTO).Return(user, nil)
	body, _ := json.Marshal(registerDTO)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Create a Gin context and call the controller method
	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req

	suite.controller.Register(c)

	assert.Equal(suite.T(), http.StatusCreated, suite.recorder.Code)
	assert.Contains(suite.T(), suite.recorder.Body.String(), "User registered successfully")
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *AuthControllerTestSuite) TestRegister_InvalidRequest() {
	invalidDTO := map[string]string{
		"Email": "invalid-email-format",
	}
	body, _ := json.Marshal(invalidDTO)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Create a Gin context and call the controller method
	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req

	suite.controller.Register(c)

	assert.Equal(suite.T(), http.StatusBadRequest, suite.recorder.Code)
	assert.Contains(suite.T(), suite.recorder.Body.String(), "Invalid request")
}

func (suite *AuthControllerTestSuite) TestRegister_Failure() {
	registerDTO := dto.RegisterDTO{
		Email:    "test@example.com",
		Password: "password123",
		FullName: "John Doe",
		Category: "general",
		Language: "Amharic",
	}
	suite.mockUsecase.On("Register", registerDTO).Return(nil, domain.ErrUserCreationFailed)
	body, _ := json.Marshal(registerDTO)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Create a Gin context and call the controller method
	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req

	suite.controller.Register(c)

	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)
	assert.Contains(suite.T(), suite.recorder.Body.String(), "Registration failed")
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *AuthControllerTestSuite) TestLogin_Success() {
	loginDTO := dto.LoginDTO{
		Email:    "test@example.com",
		Password: "password123",
	}
	acToken := "access-token"
	refToken := "refresh-token"

	suite.mockUsecase.On("Login", loginDTO).Return(&domain.User{}, acToken, refToken, nil)
	body, _ := json.Marshal(loginDTO)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req

	suite.controller.Login(c)

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)
	assert.Contains(suite.T(), suite.recorder.Body.String(), "Login successful")
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *AuthControllerTestSuite) TestLogin_InvalidRequest() {
	invalidDTO := map[string]string{
		"Email": "invalid-email-format",
	}
	body, _ := json.Marshal(invalidDTO)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req

	suite.controller.Login(c)

	assert.Equal(suite.T(), http.StatusBadRequest, suite.recorder.Code)
	assert.Contains(suite.T(), suite.recorder.Body.String(), "Invalid request")
}

func (suite *AuthControllerTestSuite) TestLogin_Failure() {
	loginDTO := dto.LoginDTO{
		Email:    "test@example.com",
		Password: "password123",
	}
	suite.mockUsecase.On("Login", loginDTO).Return(nil, "", "", domain.ErrInvalidCredentials)
	body, _ := json.Marshal(loginDTO)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req

	suite.controller.Login(c)

	assert.Equal(suite.T(), http.StatusUnauthorized, suite.recorder.Code)
	assert.Contains(suite.T(), suite.recorder.Body.String(), "Login Failed")
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *AuthControllerTestSuite) TestRefreshToken_Success() {
	refreshToken := "old-refresh-token"
	newAccessToken := "new-access-token"
	newRefreshToken := "new-refresh-token"

	suite.mockUsecase.On("RefreshToken", refreshToken).Return(newAccessToken, newRefreshToken, nil)
	body, _ := json.Marshal(refreshToken)
	req, _ := http.NewRequest(http.MethodPost, "/refresh-token", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req

	suite.controller.RefreshToken(c)

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)
	assert.Contains(suite.T(), suite.recorder.Body.String(), "Token refreshed successfully")
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *AuthControllerTestSuite) TestRefreshToken_InvalidRequest() {
	body, _ := json.Marshal(0)
	req, _ := http.NewRequest(http.MethodPost, "/refresh-token", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req

	suite.controller.RefreshToken(c)

	assert.Equal(suite.T(), http.StatusBadRequest, suite.recorder.Code)
	assert.Contains(suite.T(), suite.recorder.Body.String(), "Invalid request")
}

func (suite *AuthControllerTestSuite) TestRefreshToken_Failure() {
	refreshToken := "old-refresh-token"
	suite.mockUsecase.On("RefreshToken", refreshToken).Return("", "", domain.ErrRefreshTokenGenerationFailed)
	body, _ := json.Marshal(refreshToken)
	req, _ := http.NewRequest(http.MethodPost, "/refresh-token", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req

	suite.controller.RefreshToken(c)

	assert.Equal(suite.T(), http.StatusUnauthorized, suite.recorder.Code)
	assert.Contains(suite.T(), suite.recorder.Body.String(), "Token refresh failed")
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *AuthControllerTestSuite) TestForgotPassword_Success() {
	email := "test@example.com"
	suite.mockUsecase.On("ForgotPassword", email).Return(nil)
	body, _ := json.Marshal(email)
	req, _ := http.NewRequest(http.MethodPost, "/forgot-password", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req

	suite.controller.ForgotPassword(c)

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)
	assert.Contains(suite.T(), suite.recorder.Body.String(), "Password reset email sending successful")
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *AuthControllerTestSuite) TestForgotPassword_Failure() {
	email := "test@example.com"
	suite.mockUsecase.On("ForgotPassword", email).Return(domain.ErrEmailSendingFailed)
	body, _ := json.Marshal(email)
	req, _ := http.NewRequest(http.MethodPost, "/forgot-password", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req

	suite.controller.ForgotPassword(c)

	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)
	assert.Contains(suite.T(), suite.recorder.Body.String(), "Password reset email sending failed")
	suite.mockUsecase.AssertExpectations(suite.T())
}
func (suite *AuthControllerTestSuite) TestResetPassword_Success() {
	resetDTO := dto.ResetPasswordRequestDTO{
		Token:       "valid-token",
		NewPassword: "newPassword123",
	}

	suite.mockUsecase.On("ResetPassword", resetDTO.Token, resetDTO.NewPassword).Return(nil)
	body, _ := json.Marshal(resetDTO)
	req, _ := http.NewRequest(http.MethodPost, "/reset-password", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req

	suite.controller.ResetPassword(c)

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)
	assert.Contains(suite.T(), suite.recorder.Body.String(), "Password reset successful")
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *AuthControllerTestSuite) TestResetPassword_InvalidRequest() {
	invalidDTO := map[string]string{
		"Token": "some-token",
	}
	body, _ := json.Marshal(invalidDTO)
	req, _ := http.NewRequest(http.MethodPost, "/reset-password", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req

	suite.controller.ResetPassword(c)

	assert.Equal(suite.T(), http.StatusBadRequest, suite.recorder.Code)
	assert.Contains(suite.T(), suite.recorder.Body.String(), "Invalid request")
}

func (suite *AuthControllerTestSuite) TestResetPassword_Failure() {
	resetDTO := dto.ResetPasswordRequestDTO{
		Token:       "invalid-token",
		NewPassword: "newPassword123",
	}

	suite.mockUsecase.On("ResetPassword", resetDTO.Token, resetDTO.NewPassword).Return(domain.ErrResetTokenGenerationFailed)
	body, _ := json.Marshal(resetDTO)
	req, _ := http.NewRequest(http.MethodPost, "/reset-password", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(suite.recorder)
	c.Request = req

	suite.controller.ResetPassword(c)

	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)
	assert.Contains(suite.T(), suite.recorder.Body.String(), "Password reset failed")
	suite.mockUsecase.AssertExpectations(suite.T())
}

func (suite *AuthControllerTestSuite) TearDownTest() {
	suite.mockUsecase.AssertExpectations(suite.T())
}

func TestAuthControllerTestSuite(t *testing.T) {
	suite.Run(t, new(AuthControllerTestSuite))
}
