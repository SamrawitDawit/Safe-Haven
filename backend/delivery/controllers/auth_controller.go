package controllers

import (
	"backend/delivery/config"
	"backend/domain"
	"backend/usecases"
	"backend/usecases/dto"
	"backend/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"
)

type AuthController struct {
	userUsecase usecases.AuthUseCaseInterface
}

func NewAuthController(userUsecase usecases.AuthUseCaseInterface) *AuthController {
	return &AuthController{
		userUsecase: userUsecase,
	}
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var registerDTO dto.RegisterDTO
	if err := c.BindJSON(&registerDTO); err != nil {
		utils.LogError("Error binding request", err)
		res := utils.ErrorResponse(http.StatusBadRequest, "Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}
	created_user, err := ctrl.userUsecase.Register(registerDTO)
	if err != nil {
		res := utils.ErrorResponse(err.StatusCode, "Registration failed", err.Message)
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	c.JSON(http.StatusCreated, utils.SuccessResponse(http.StatusCreated, "User registered successfully", created_user))
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var loginDTO dto.LoginDTO
	if err := c.BindJSON(&loginDTO); err != nil {
		utils.LogError("Error binding request", err)
		res := utils.ErrorResponse(http.StatusBadRequest, "Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}
	user, acToken, refToken, err := ctrl.userUsecase.Login(loginDTO)
	if err != nil {
		res := utils.ErrorResponse(err.StatusCode, "Login Failed", err.Message)
		c.JSON(http.StatusUnauthorized, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "Login successful", map[string]interface{}{"user": user, "accessToken": acToken, "refreshToken": refToken}))
}
func (ctrl *AuthController) RefreshToken(c *gin.Context) {
	var refreshToken string
	if err := c.BindJSON(&refreshToken); err != nil {
		utils.LogError("Error binding request", err)
		res := utils.ErrorResponse(http.StatusBadRequest, "Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	newAccessToken, newRefreshToken, err := ctrl.userUsecase.RefreshToken(refreshToken)
	if err != nil {
		res := utils.ErrorResponse(err.StatusCode, "Token refresh failed", err.Message)
		c.JSON(http.StatusUnauthorized, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "Token refreshed successfully", map[string]string{"accessToken": newAccessToken, "refreshToken": newRefreshToken}))
}

func (ctrl *AuthController) ForgotPassword(c *gin.Context) {
	var email string
	if err := c.BindJSON(&email); err != nil {
		utils.LogError("Error binding request", err)
		res := utils.ErrorResponse(http.StatusBadRequest, "Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err := ctrl.userUsecase.ForgotPassword(email)
	if err != nil {
		res := utils.ErrorResponse(err.StatusCode, "Password reset email sending failed", err.Message)
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "Password reset email sending successful", nil))
}

func (ctrl *AuthController) ResetPassword(c *gin.Context) {
	var ResetPassworddto dto.ResetPasswordRequestDTO
	if err := c.BindJSON(&ResetPassworddto); err != nil {
		utils.LogError("Error binding request", err)
		res := utils.ErrorResponse(http.StatusBadRequest, "Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err := ctrl.userUsecase.ResetPassword(ResetPassworddto.Token, ResetPassworddto.NewPassword)
	if err != nil {
		res := utils.ErrorResponse(err.StatusCode, "Password reset failed", err.Message)
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "Password reset successful", nil))
}
func (ctrl *AuthController) HandleGoogleCallback(c *gin.Context) {
	// Continue with the Google OAuth callback process
	code := c.Query("code")
	state := c.Query("state")
	var token *oauth2.Token
	var err error

	client_type := c.Query("client")
	fmt.Println("client", client_type)
	var redirectURI string
	if client_type == "mobile" {
		redirectURI = "https://3613-196-189-114-138.ngrok-free.app/auth/google/callback?client=mobile"
	} else {
		redirectURI = "http://localhost:8080/auth/google/callback"
	}

	// Create googleConfig dynamically based on the redirectURI
	googleConfig := config.GoogleConfig(redirectURI)
	if client_type == "mobile" {
		token, err = googleConfig.Exchange(c, code, oauth2.SetAuthURLParam("code_verifier", state))
		if err != nil {
			utils.LogError("Error exchanging code for token", err)
			res := utils.ErrorResponse(http.StatusInternalServerError, "Google login failed", err.Error())
			c.JSON(http.StatusInternalServerError, res)
			return
		}
	} else {

		// Retrieve the state from the cookie
		storedState, err := c.Cookie("oauthstate")
		if err != nil {
			utils.LogError("Invalid state token", err)
			res := utils.ErrorResponse(http.StatusBadRequest, "State token missing or expired", "Failed to retrieve state token")
			c.JSON(http.StatusBadRequest, res)
			return
		}

		// Validate the state token
		if state != storedState {
			res := utils.ErrorResponse(http.StatusBadRequest, "Invalid state token", "CSRF protection failed")
			c.JSON(http.StatusBadRequest, res)
			return
		}

		// Continue with the Google OAuth callback process
		code := c.Query("code")
		token, err = googleConfig.Exchange(c, code)
		if err != nil {
			utils.LogError("Google Exchange failed", err)
			res := utils.ErrorResponse(http.StatusInternalServerError, "Google login failed", err.Error())
			c.JSON(http.StatusInternalServerError, res)
			return
		}
	}
	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		res := utils.ErrorResponse(http.StatusInternalServerError, "Google login failed", "Failed to get id_token")
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	payload, err := idtoken.Validate(c, idToken, googleConfig.ClientID)
	if err != nil {
		utils.LogError("Error token validation", err)
		res := utils.ErrorResponse(http.StatusInternalServerError, "Google login failed", err.Error())
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	userInfo := &domain.User{
		FullName:     payload.Claims["name"].(string),
		Email:        payload.Claims["email"].(string),
		GoogleSignin: true,
		Role:         "regular",
		Active:       true,
	}

	picture, ok := payload.Claims["picture"].(string)
	if ok {
		userInfo.ImageURL = picture
	}

	user, accessToken, refreshToken, cerr := ctrl.userUsecase.HandleGoogleCallback(userInfo)
	if cerr != nil {
		res := utils.ErrorResponse(cerr.StatusCode, "Google login failed", cerr.Message)
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	c.SetCookie("access_token", accessToken, 3600, "/", "localhost", true, true) // Secure, HttpOnly, SameSite=Lax
	c.SetCookie("refresh_token", refreshToken, 7200, "/", "localhost", true, true)
	c.Redirect(http.StatusFound, "http://localhost:3000/auth/dummy-dashboard?accessToken="+accessToken+"&refreshToken="+refreshToken+"&user="+string(user.ID.String()))

}

func (ctrl *AuthController) GoogleLogin(c *gin.Context) {
	// Generate the state token
	stateToken, err := utils.GenerateStateToken()
	if err != nil {
		res := utils.ErrorResponse(http.StatusInternalServerError, "Error generating state token", err.Error())
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	// Store the state token in a cookie
	c.SetCookie("oauthstate", stateToken, 3600, "/", "localhost", false, true)

	// Redirect to Google with the state token
	googleConfig := config.GoogleConfig("http://localhost:8080/auth/google/callback")
	url := googleConfig.AuthCodeURL(stateToken)
	c.Redirect(http.StatusTemporaryRedirect, url)
}
