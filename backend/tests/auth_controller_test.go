package tests

import (
	"backend/delivery/controllers"
	"backend/tests/mocks"
	"backend/usecases/dto"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	// "cloud.google.com/go/auth/credentials/idtoken"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/oauth2"
)

type AuthControllerTestSuite struct {
	suite.Suite
	mockUsecase *mocks.AuthUseCaseInterface
	controller  *controllers.AuthContoller
	recorder    *httptest.ResponseRecorder
	googleConfig *oauth2.Config
}

func (suite *AuthControllerTestSuite) SetupTest() {
	suite.mockUsecase = new(mocks.AuthUseCaseInterface)
	suite.recorder = httptest.NewRecorder()
	suite.googleConfig = &oauth2.Config{} // Mock or setup a minimal OAuth2 config if needed
	suite.controller = controllers.NewAuthController(suite.mockUsecase, suite.googleConfig)

	// Set Gin mode to TestMode to avoid debug output during tests
	gin.SetMode(gin.TestMode)
}

func (suite *AuthControllerTestSuite) TestRegister_Success() {
	registerDTO := dto.RegisterDTO{
		Email:    "test@example.com",
		Password: "password123",
		FullName: "John Doe",
		UserType: "normal",
	}
	suite.mockUsecase.On("Register", registerDTO).Return(nil)
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
		UserType: "normal",
	}
	suite.mockUsecase.On("Register", registerDTO).Return(errors.New("registration error"))
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
		UserType: "normal",
	}
	acToken := "access-token"
	refToken := "refresh-token"

	suite.mockUsecase.On("Login", loginDTO).Return(acToken, refToken, nil)
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
		UserType: "normal",
	}
	suite.mockUsecase.On("Login", loginDTO).Return("", "", errors.New("login failed"))
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
	suite.mockUsecase.On("RefreshToken", refreshToken).Return("", "", errors.New("token refresh failed"))
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
	suite.mockUsecase.On("ForgotPassword", email).Return(errors.New("failed to send email"))
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
		Token:      "valid-token",
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
		Token:      "invalid-token",
		NewPassword: "newPassword123",
	}

	suite.mockUsecase.On("ResetPassword", resetDTO.Token, resetDTO.NewPassword).Return(errors.New("reset failed"))
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

// func (suite *AuthControllerTestSuite) TestHandleGoogleCallback_Success() {
// 	code := "valid-code"
// 	token := &oauth2.Token{
// 		AccessToken:  "access-token",
// 		RefreshToken: "refresh-token",
// 	}
// 	idToken := "valid-id-token"
// 	payload := &idtoken.Payload{
// 		Claims: map[string]interface{}{
// 			"email": "test@example.com",
// 			"name":  "John Doe",
// 		},
// 	}

// 	suite.mockUsecase.On("HandleGoogleCallback", mock.Anything).Return("access-token", "refresh-token", nil)
// 	suite.mockUsecase.On("Exchange", mock.Anything, code).Return(token, nil)
// 	suite.mockUsecase.On("Validate", mock.Anything, idToken, suite.googleConfig.ClientID).Return(payload, nil)

// 	req, _ := http.NewRequest(http.MethodGet, "/callback?code="+code, nil)
// 	c, _ := gin.CreateTestContext(suite.recorder)
// 	c.Request = req

// 	suite.controller.HandleGoogleCallback(c)

// 	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)
// 	assert.Contains(suite.T(), suite.recorder.Body.String(), "Google login successful")
// 	suite.mockUsecase.AssertExpectations(suite.T())
// }

// func (suite *AuthControllerTestSuite) TestHandleGoogleCallback_ExchangeFailure() {
// 	code := "invalid-code"

// 	suite.mockUsecase.On("Exchange", mock.Anything, code).Return(nil, errors.New("exchange failed"))

// 	req, _ := http.NewRequest(http.MethodGet, "/callback?code="+code, nil)
// 	c, _ := gin.CreateTestContext(suite.recorder)
// 	c.Request = req

// 	suite.controller.HandleGoogleCallback(c)

// 	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)
// 	assert.Contains(suite.T(), suite.recorder.Body.String(), "Google login failed")
// 	suite.mockUsecase.AssertExpectations(suite.T())
// }

// func (suite *AuthControllerTestSuite) TestHandleGoogleCallback_ValidationFailure() {
// 	code := "valid-code"
// 	token := &oauth2.Token{
// 		AccessToken:  "access-token",
// 		RefreshToken: "refresh-token",
// 	}
// 	idToken := "valid-id-token"

// 	suite.mockUsecase.On("Exchange", mock.Anything, code).Return(token, nil)
// 	suite.mockUsecase.On("Validate", mock.Anything, idToken, suite.googleConfig.ClientID).Return(nil, errors.New("validation failed"))

// 	req, _ := http.NewRequest(http.MethodGet, "/callback?code="+code, nil)
// 	c, _ := gin.CreateTestContext(suite.recorder)
// 	c.Request = req

// 	suite.controller.HandleGoogleCallback(c)

// 	assert.Equal(suite.T(), http.StatusInternalServerError, suite.recorder.Code)
// 	assert.Contains(suite.T(), suite.recorder.Body.String(), "Google login failed")
// 	suite.mockUsecase.AssertExpectations(suite.T())
// }

func (suite *AuthControllerTestSuite) TearDownTest() {
	suite.mockUsecase.AssertExpectations(suite.T())
}

func TestAuthControllerTestSuite(t *testing.T) {
	suite.Run(t, new(AuthControllerTestSuite))
}
