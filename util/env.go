package util

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func SecretsEnv(key string) string {
	err := godotenv.Load("secrets.env")

	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}

func GetToken() string {
	return SecretsEnv("TOKEN")
}

func UrlEnv(key string) string {
	err := godotenv.Load("urls.env")
	if err != nil {
		fmt.Print(err)
	}
	var url string
	if os.Getenv("ENV") == "dev" {
		url = key + "dev"
	} else {
		url = key + "prod"
	}
	return os.Getenv(url)
}

func GetURL(key string) string {
	return UrlEnv(key)
}
