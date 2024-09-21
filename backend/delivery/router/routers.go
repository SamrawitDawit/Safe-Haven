package router

import (
	"backend/delivery/config"
	"backend/delivery/controllers"
	"backend/infrastructure"

	"github.com/gin-gonic/gin"
)

type RouterControllers struct {
	AuthController *controllers.AuthController
}

type RouterServices struct {
	JwtService *infrastructure.JWTService
}

func NewRouter(routerControllers *RouterControllers, routerServices *RouterServices) {
	router := gin.Default()

	// Middlewares
	router.Use(infrastructure.CorsMiddleware())

	// Auth routes
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", routerControllers.AuthController.Register)
		authRoutes.POST("/login", routerControllers.AuthController.Login)
		authRoutes.POST("/refresh-token", routerControllers.AuthController.RefreshToken)
		authRoutes.POST("/forgot-password", routerControllers.AuthController.ForgotPassword)
		authRoutes.POST("/reset-password", routerControllers.AuthController.ResetPassword)
		authRoutes.GET("/google", routerControllers.AuthController.GoogleLogin)
		authRoutes.GET("/google/callback", routerControllers.AuthController.HandleGoogleCallback)
	}

	// Protected routes example
	protectedRoutes := router.Group("/user")
	protectedRoutes.Use(infrastructure.AuthMiddleware(routerServices.JwtService))

	// Server run on defined port
	router.Run(":" + config.ENV.PORT)
}
