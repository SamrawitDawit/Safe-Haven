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
		caseRoutes.PUT("/update/:case_id", routerControllers.CaseController.UpdateCase)
		caseRoutes.DELETE("/delete/:case_id", routerControllers.CaseController.DeleteCase)
		caseRoutes.GET("/", routerControllers.CaseController.GetAllCases)
		caseRoutes.GET("/id/:case_id", routerControllers.CaseController.GetCaseByID)                       // Separate case_id route
		caseRoutes.GET("/submitter/:submitter_id", routerControllers.CaseController.GetCasesBySubmitterID) // Separate submitter route
		caseRoutes.GET("/counselor/:counselor_id", routerControllers.CaseController.GetCasesByCounselorID) // Separate counselor route
		caseRoutes.GET("/status/:status", routerControllers.CaseController.GetCasesByStatus)               // Separate status route
	}

	caseRoutes.Use(infrastructure.AuthMiddleware(routerServices.JwtService))

	// Server run on defined port
	router.Run(":" + config.ENV.PORT)
}
