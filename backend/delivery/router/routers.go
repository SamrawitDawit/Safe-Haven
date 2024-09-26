package router

import (
	"backend/delivery/config"
	"backend/delivery/controllers"
	"backend/infrastructure"

	"github.com/gin-gonic/gin"
)

type RouterControllers struct {
	AuthController *controllers.AuthController
	CaseController *controllers.CaseController
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
	caseRoutes := router.Group("/case")
	{
		caseRoutes.POST("/submit", routerControllers.CaseController.CreateCase)
		caseRoutes.PUT("/update", routerControllers.CaseController.UpdateCase)
		caseRoutes.DELETE("/delete", routerControllers.CaseController.DeleteCase)
		caseRoutes.POST("/", routerControllers.CaseController.GetAllCases)
		caseRoutes.GET("/case_id", routerControllers.CaseController.GetCaseByID)
		caseRoutes.GET("/submitter_id", routerControllers.CaseController.GetCasesBySubmitterID)
		caseRoutes.GET("/counselor_id", routerControllers.CaseController.GetCasesByCounselorID)
		caseRoutes.GET("/status", routerControllers.CaseController.GetCasesByStatus)
		
	}
	caseRoutes.Use(infrastructure.AuthMiddleware(routerServices.JwtService))

	// Server run on defined port
	router.Run(":" + config.ENV.PORT)
}
