package controllers

import (
	"backend/usecases"
	"backend/usecases/dto"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthContoller struct {
	userUsecase usecases.AuthUseCaseInterface
}

func NewAuthController(userUsecase usecases.AuthUseCaseInterface) *AuthContoller {
	return &AuthContoller{
		userUsecase: userUsecase,
	}
}

func (ctrl *AuthContoller) Register(c *gin.Context) {
	var registerDTO dto.RegisterDTO
	if err := c.BindJSON(&registerDTO); err != nil {
		res := utils.ErrorResponse("Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusCreated, utils.SuccessResponse("User registered successfully", nil))
}

func (ctrl *AuthContoller) Login(c *gin.Context) {
	var loginDTO dto.LoginDTO
	if err := c.BindJSON(&loginDTO); err != nil {
		res := utils.ErrorResponse("Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	tokens, err := ctrl.userUsecase.Login(loginDTO)
	if err != nil {
		res := utils.ErrorResponse("Login Failed", err.Error())
		c.JSON(http.StatusUnauthorized, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse("Login successful", tokens))
}

func (ctrl *AuthContoller) AnonymousRegister(c *gin.Context) {
	var anonUserDTO dto.RegisterDTO
	if err := c.BindJSON(&anonUserDTO); err != nil {
		res := utils.ErrorResponse("Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err := ctrl.userUsecase.Register(anonUserDTO)
	if err != nil {
		res := utils.ErrorResponse("Registration failed", err.Error())
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	c.JSON(http.StatusCreated, utils.SuccessResponse("Anonymous user registered successfully", nil))
}

func (ctrl *AuthContoller) AnonymousLogin(c *gin.Context) {
	var anonUserDTO dto.LoginDTO
	if err := c.BindJSON(&anonUserDTO); err != nil {
		res := utils.ErrorResponse("Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	tokens, err := ctrl.userUsecase.AnonymousLogin(anonUserDTO)
	if err != nil {
		res := utils.ErrorResponse("Login Failed", err.Error())
		c.JSON(http.StatusUnauthorized, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse("Login successful", tokens))
}

func (ctrl *AuthContoller) RefreshToken(c *gin.Context) {
	var refreshToken string
	if err := c.BindJSON(&refreshToken); err != nil {
		res := utils.ErrorResponse("Invalid request", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	newAccessToken, newRefreshToken, err := ctrl.userUsecase.RefreshToken(refreshToken)
	if err != nil {
		res := utils.ErrorResponse("Token refresh failed", err.Error())
		c.JSON(http.StatusUnauthorized, res)
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse("Token refreshed successfully", map[string]string{"accessToken": newAccessToken, "refreshToken": newRefreshToken}))
}
