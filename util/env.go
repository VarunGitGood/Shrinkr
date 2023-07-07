package util

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Env(key string) string {
	err := godotenv.Load("secrets.env")

	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
