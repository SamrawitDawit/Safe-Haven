package router

import (
	"backend/delivery/config"
	"backend/delivery/controllers"
	"backend/infrastructure"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type RouterControllers struct {
	AuthController *controllers.AuthContoller
}

type RouterServices struct {
	JwtService *infrastructure.JWTService
}

func NewRouter(db *mongo.Database, routerControllers *RouterControllers, routerServices *RouterServices) {
	router := gin.Default()
	// jwtService := routerServices.JwtService
	router.POST("/register", routerControllers.AuthController.Register)
	router.POST("/login", routerControllers.AuthController.Login)
	router.POST("/refresh-token", routerControllers.AuthController.RefreshToken)
	router.POST("/forgot-password", routerControllers.AuthController.ForgotPassword)
	router.POST("/reset-password", routerControllers.AuthController.ResetPassword)
	router.GET("/auth/google", routerControllers.AuthController.GoogleLogin)
	router.GET("/auth/google/callback", routerControllers.AuthController.HandleGoogleCallback)

	router.Run(":" + config.ENV.PORT)
}
