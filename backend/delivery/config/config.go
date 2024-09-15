package config

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EnvironmentVariables struct {
	PORT                 string
	DB_URI               string
	DB_NAME              string
	JWT_SECRET           string
	EMAIL_HOST           string
	EMAIL_PORT           string
	EMAIL_USERNAME       string
	EMAIL_PASSWORD       string
	EMAIL                string
	GOOGLE_CLIENT_ID     string
	GOOGLE_CLIENT_SECRET string
	USER_COLLECTION      string
}

var ENV EnvironmentVariables

func LoadEnv() error {
	err := godotenv.Load("../.env")
	if err != nil {
		return err
	}
	ENV.PORT = os.Getenv("PORT")
	ENV.DB_URI = os.Getenv("DB_URI")
	ENV.DB_NAME = os.Getenv("DB_NAME")
	ENV.JWT_SECRET = os.Getenv("JWT_SECRET")
	ENV.EMAIL_HOST = os.Getenv("EMAIL_HOST")
	ENV.EMAIL_PORT = os.Getenv("EMAIL_PORT")
	ENV.EMAIL_USERNAME = os.Getenv("EMAIL_USERNAME")
	ENV.EMAIL_PASSWORD = os.Getenv("EMAIL_PASSWORD")
	ENV.EMAIL = os.Getenv("EMAIL")
	ENV.GOOGLE_CLIENT_ID = os.Getenv("GOOGLE_CLIENT_ID")
	ENV.GOOGLE_CLIENT_SECRET = os.Getenv("GOOGLE_CLIENT_SECRET")
	ENV.USER_COLLECTION = os.Getenv("USER_COLLECTION")

	switch {
	case ENV.PORT == "":
		return fmt.Errorf("PORT is not set")
	case ENV.DB_URI == "":
		return fmt.Errorf("DB_URI is not set")
	case ENV.DB_NAME == "":
		return fmt.Errorf("DB_NAME is not set")
	case ENV.JWT_SECRET == "":
		return fmt.Errorf("JWT_SECRET is not set")
	case ENV.EMAIL_HOST == "":
		return fmt.Errorf("EMAIL_HOST is not set")
	case ENV.EMAIL_USERNAME == "":
		return fmt.Errorf("EMAIL_USERNAME is not set")
	case ENV.EMAIL_PASSWORD == "":
		return fmt.Errorf("EMAIL_PASSWORD is not set")
	case ENV.EMAIL == "":
		return fmt.Errorf("EMAIL is not set")
	case ENV.GOOGLE_CLIENT_ID == "":
		return fmt.Errorf("GOOGLE_CLIENT_ID is not set")
	case ENV.GOOGLE_CLIENT_SECRET == "":
		return fmt.Errorf("GOOGLE_CLIENT_SECRET is not set")
	case ENV.USER_COLLECTION == "":
		return fmt.Errorf("USER_COLLECTION is not set")
	default:
		return nil
	}
}

func ConnectDB(clientOption *options.ClientOptions) (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}
