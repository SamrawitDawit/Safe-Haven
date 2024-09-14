package controllers

import (
	"backend/domain"
	"backend/usecases"
	"backend/usecases/dto"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"
)

type AuthContoller struct {
	userUsecase  usecases.AuthUseCaseInterface
	googleConfig *oauth2.Config
}

func NewAuthController(userUsecase usecases.AuthUseCaseInterface, googleConfig *oauth2.Config) *AuthContoller {
	return &AuthContoller{
		userUsecase:  userUsecase,
		googleConfig: googleConfig,
	}
}

func (ctrl *AuthContoller) Register(c *gin.Context) {
	var registerDTO dto.RegisterDTO
	if err := c.BindJSON(&registerDTO); err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest, "Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if dto.ValidateRegisterDTO(registerDTO) != nil {
		res := utils.ErrorResponse(http.StatusBadRequest, "Invalid request", "Invalid request")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	err := ctrl.userUsecase.Register(registerDTO)
	if err != nil {
		res := utils.ErrorResponse(http.StatusInternalServerError, "Registration failed", err.Error())
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	c.JSON(http.StatusCreated, utils.SuccessResponse(http.StatusCreated, "User registered successfully", nil))
}

func (ctrl *AuthContoller) Login(c *gin.Context) {
	var loginDTO dto.LoginDTO
	if err := c.BindJSON(&loginDTO); err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest, "Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if dto.ValidateLoginDTO(loginDTO) != nil {
		res := utils.ErrorResponse(http.StatusBadRequest, "Invalid request", "Invalid request")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	acToken, refToken, err := ctrl.userUsecase.Login(loginDTO)
	if err != nil {
		res := utils.ErrorResponse(http.StatusUnauthorized, "Login Failed", err.Error())
		c.JSON(http.StatusUnauthorized, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "Login successful", map[string]string{"accessToken": acToken, "refreshToken": refToken}))
}
func (ctrl *AuthContoller) RefreshToken(c *gin.Context) {
	var refreshToken string
	if err := c.BindJSON(&refreshToken); err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest, "Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	newAccessToken, newRefreshToken, err := ctrl.userUsecase.RefreshToken(refreshToken)
	if err != nil {
		res := utils.ErrorResponse(http.StatusUnauthorized, "Token refresh failed", err.Error())
		c.JSON(http.StatusUnauthorized, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "Token refreshed successfully", map[string]string{"accessToken": newAccessToken, "refreshToken": newRefreshToken}))
}

func (ctrl *AuthContoller) ForgotPassword(c *gin.Context) {
	var email string
	if err := c.BindJSON(&email); err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest, "Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err := ctrl.userUsecase.ForgotPassword(email)
	if err != nil {
		res := utils.ErrorResponse(http.StatusInternalServerError, "Password reset failed", err.Error())
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "Password reset successful", nil))
}

func (ctrl *AuthContoller) ResetPassword(c *gin.Context) {
	var token, newPassword string
	if err := c.BindJSON(&token); err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest, "Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if err := c.BindJSON(&newPassword); err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest, "Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err := ctrl.userUsecase.ResetPassword(token, newPassword)
	if err != nil {
		res := utils.ErrorResponse(http.StatusInternalServerError, "Password reset failed", err.Error())
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "Password reset successful", nil))
}

func (ctrl *AuthContoller) HandleGoogleCallback(c *gin.Context) {
	code := c.Query("code")
	token, err := ctrl.googleConfig.Exchange(c, code)
	if err != nil {
		res := utils.ErrorResponse(http.StatusInternalServerError, "Google login failed", err.Error())
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		res := utils.ErrorResponse(http.StatusInternalServerError, "Google login failed", "Failed to get id_token")
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	payload, err := idtoken.Validate(c, idToken, ctrl.googleConfig.ClientID)
	if err != nil {
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
	accessToken, refreshToken, err := ctrl.userUsecase.HandleGoogleCallback(userInfo)
	if err != nil {
		res := utils.ErrorResponse(http.StatusInternalServerError, "Google login failed", err.Error())
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(http.StatusOK, "Google login successful", map[string]string{"accessToken": accessToken, "refreshToken": refreshToken}))
}

func (ctrl *AuthContoller) GoogleLogin(c *gin.Context) {
	url := ctrl.googleConfig.AuthCodeURL("state-token")
	c.Redirect(http.StatusTemporaryRedirect, url)
}
