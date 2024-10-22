package config

import (
	"context"
	"fmt"
	"os"
	"reflect"

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
	CASE_COLLECTION      string
	ENCRYPT_KEY          string
	PROJECT_ID           string
	RECAPTCHA_KEY        string
	RECAPTCHA_ACTION     string
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
	ENV.CASE_COLLECTION = os.Getenv("CASE_COLLECTION")
	ENV.ENCRYPT_KEY = os.Getenv("ENCRYPT_KEY")
	ENV.PROJECT_ID = os.Getenv("PROJECT_ID")
	ENV.RECAPTCHA_KEY = os.Getenv("RECAPTCHA_KEY")
	ENV.RECAPTCHA_ACTION = os.Getenv("RECAPTCHA_ACTION")

	v := reflect.ValueOf(ENV)
	typeOfEnv := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i).String()
		fieldName := typeOfEnv.Field(i).Name

		if fieldValue == "" {
			return fmt.Errorf("%s is not set", fieldName)
		}
	}
	return nil
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
