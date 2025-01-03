package main

import (
	"backend/delivery/config"
	"backend/delivery/controllers"
	"backend/delivery/router"
	"backend/infrastructure"
	"backend/repositories"
	"backend/usecases"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	clientOption := options.Client().ApplyURI(config.ENV.DB_URI)
	client, err := config.ConnectDB(clientOption)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	db := client.Database(config.ENV.DB_NAME)
	jwtService := infrastructure.JWTService{JwtSecret: config.ENV.JWT_SECRET}
	pwdService := infrastructure.HashingService{}
	encryptService := infrastructure.EncryptionService{Key: config.ENV.ENCRYPT_KEY}
	emailService := infrastructure.EmailService{
		Host:     config.ENV.EMAIL_HOST,
		Port:     config.ENV.EMAIL_PORT,
		UserName: config.ENV.EMAIL_USERNAME,
		Password: config.ENV.EMAIL_PASSWORD,
		Email:    config.ENV.EMAIL,
	}
	recaptchaService := infrastructure.RecaptchaService{}
	userRepo := repositories.NewUserRepo(db, config.ENV.USER_COLLECTION)
	caseRepo := repositories.NewCaseRepo(db, config.ENV.CASE_COLLECTION)
	authUsecase := usecases.NewAuthUseCase(userRepo, &jwtService, &emailService, &pwdService, &encryptService)
	caseUsecase := usecases.NewCaseUseCase(caseRepo, &encryptService)
	authController := controllers.NewAuthController(authUsecase)
	caseController := controllers.NewCaseController(caseUsecase, &recaptchaService)
	wsController := infrastructure.SetupWebSocketServer()
	http.HandleFunc("/ws", wsController.ServeWS)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	router.NewRouter(
		&router.RouterControllers{
			AuthController: authController,
			CaseController: caseController,
		},
		&router.RouterServices{
			JwtService: &jwtService,
		},
	)
}
