package tests

import (
	"backend/domain"
	"backend/infrastructure"
	"backend/tests/mocks"
	"backend/utils"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MiddlewareTestSuite struct {
	suite.Suite
	mockJwtService *mocks.JwtServiceInterface
	router         *gin.Engine
}

func (s *MiddlewareTestSuite) SetupTest() {
	s.mockJwtService = new(mocks.JwtServiceInterface)
	s.router = gin.Default()
	s.router.Use(gin.Logger())

	// Setup routes for testing
	authGroup := s.router.Group("/")
	authGroup.Use(infrastructure.AuthMiddleware(s.mockJwtService))
	authGroup.GET("/protected", func(c *gin.Context) {
		res := utils.SuccessResponse(200, "Success", "Success")
		c.JSON(http.StatusOK, res)
	})

	adminGroup := s.router.Group("/admin")
	adminGroup.Use(infrastructure.AuthMiddleware(s.mockJwtService), infrastructure.AdminMiddleWare())
	adminGroup.GET("/admin-protected", func(c *gin.Context) {
		res := utils.SuccessResponse(200, "Success", "Success")
		c.JSON(http.StatusOK, res)
	})
}

func (s *MiddlewareTestSuite) TestAuthMiddleware_Success() {
	token := &jwt.Token{Valid: true}
	claims := jwt.MapClaims{
		"id":   uuid.New().String(),
		"role": "regular",
	}

	s.mockJwtService.On("ValidateToken", mock.Anything).Return(token, nil)
	s.mockJwtService.On("ExtractTokenClaims", token).Return(claims, nil)

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer valid.token")
	w := httptest.NewRecorder()

	s.router.ServeHTTP(w, req)

	res := utils.SuccessResponse(200, "Success", "Success")
	resJSON, _ := json.Marshal(res)
	assert.Equal(s.T(), http.StatusOK, w.Code)
	assert.JSONEq(s.T(), string(resJSON), w.Body.String())
}

func (s *MiddlewareTestSuite) TestAuthMiddleware_Failure_NoToken() {
	req, _ := http.NewRequest("GET", "/protected", nil)
	w := httptest.NewRecorder()

	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusUnauthorized, w.Code)
	res := utils.ErrorResponse(401, "No token provided", "No token provided")
	resJSON, _ := json.Marshal(res)
	assert.JSONEq(s.T(), string(resJSON), w.Body.String())
}

func (s *MiddlewareTestSuite) TestAuthMiddleware_Failure_InvalidTokenFormat() {
	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "InvalidTokenFormat")
	w := httptest.NewRecorder()

	s.router.ServeHTTP(w, req)

	res := utils.ErrorResponse(401, "Invalid token", "Invalid token")
	resJSON, _ := json.Marshal(res)
	assert.Equal(s.T(), http.StatusUnauthorized, w.Code)
	assert.JSONEq(s.T(), string(resJSON), w.Body.String())
}

func (s *MiddlewareTestSuite) TestAuthMiddleware_Failure_ExpiredOrInvalidToken() {
	s.mockJwtService.On("ValidateToken", mock.Anything).Return(nil, &domain.CustomError{Message: "Invalid or expired token"})

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer invalid.token")
	w := httptest.NewRecorder()

	s.router.ServeHTTP(w, req)

	res := utils.ErrorResponse(401, "Invalid token", "Invalid token")
	resJSON, _ := json.Marshal(res)
	assert.Equal(s.T(), http.StatusUnauthorized, w.Code)
	assert.JSONEq(s.T(), string(resJSON), w.Body.String())
}

func (s *MiddlewareTestSuite) TestAdminMiddleware_Success() {
	token := &jwt.Token{Valid: true}
	claims := jwt.MapClaims{
		"id":   uuid.New().String(),
		"role": "admin",
	}

	s.mockJwtService.On("ValidateToken", mock.Anything).Return(token, nil)
	s.mockJwtService.On("ExtractTokenClaims", token).Return(claims, nil)

	req, _ := http.NewRequest("GET", "/admin/admin-protected", nil)
	req.Header.Set("Authorization", "Bearer valid.token")
	w := httptest.NewRecorder()

	s.router.ServeHTTP(w, req)
	res := utils.SuccessResponse(200, "Success", "Success")
	resJSON, _ := json.Marshal(res)
	assert.Equal(s.T(), http.StatusOK, w.Code)
	assert.JSONEq(s.T(), string(resJSON), w.Body.String())
}

func (s *MiddlewareTestSuite) TestAdminMiddleware_Failure_NotAdmin() {
	token := &jwt.Token{Valid: true}
	claims := jwt.MapClaims{
		"id":   uuid.New().String(),
		"role": "regular",
	}

	s.mockJwtService.On("ValidateToken", mock.Anything).Return(token, nil)
	s.mockJwtService.On("ExtractTokenClaims", token).Return(claims, nil)

	req, _ := http.NewRequest("GET", "/admin/admin-protected", nil)
	req.Header.Set("Authorization", "Bearer valid.token")
	w := httptest.NewRecorder()

	s.router.ServeHTTP(w, req)

	res := utils.ErrorResponse(403, "Sorry, you must be an admin", "Sorry, you must be an admin")
	resJSON, _ := json.Marshal(res)
	assert.Equal(s.T(), http.StatusForbidden, w.Code)
	assert.JSONEq(s.T(), string(resJSON), w.Body.String())
}

func TestMiddlewareTestSuite(t *testing.T) {
	suite.Run(t, new(MiddlewareTestSuite))
}
