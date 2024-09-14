package main

import (
	"backend/delivery/config"
	"backend/delivery/controllers"
	"backend/delivery/router"
	"backend/infrastructure"
	"backend/repositories"
	"backend/usecases"
	"log"

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
	emailService := infrastructure.EmailService{
		Host:     config.ENV.EMAIL_HOST,
		UserName: config.ENV.EMAIL_USERNAME,
		Password: config.ENV.EMAIL_PASSWORD,
		Email:    config.ENV.EMAIL,
	}
	userRepo := repositories.NewUserRepo(db, config.ENV.USER_COLLECTION)
	authUsecase := usecases.NewAuthUseCase(userRepo, &jwtService, &emailService, &pwdService)
	authController := controllers.NewAuthController(authUsecase, config.GoogleOAuthConfig)

	router.NewRouter(
		db,
		&router.RouterControllers{
			AuthController: authController,
		},
		&router.RouterServices{
			JwtService: &jwtService,
		},
	)
}
