package util

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const env = "prod"

// create a key value map for urls
var URLS = map[string]string{
	"checkURL_prod": "https://shrinkr-da1u.onrender.com/shrinkr/links/check/",
	"addURL_prod":   "https://shrinkr-da1u.onrender.com/shrinkr/links/addurl",
	"baseURL_prod":  "https://shrinkr-da1u.onrender.com/shrinkr/",
	"loginURL_prod": "https://shrinkr-da1u.onrender.com/shrinkr/login",
	"tokenURL_prod": "https://shrinkr-da1u.onrender.com/shrinkr/token",
	"checkURL_dev":  "http://localhost:8080/shrinkr/links/check/",
	"addURL_dev":    "http://localhost:8080/shrinkr/links/addurl",
	"baseURL_dev":   "http://localhost:8080/shrinkr/",
	"loginURL_dev":  "http://localhost:8080/shrinkr/login",
	"tokenURL_dev":  "http://localhost:8080/shrinkr/token",
}

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

func GetURL(key string) string {
	flag := ""
	if env == "dev" {
		flag = flag + "_dev"
	} else {
		flag = flag + "_prod"
	}
	url := URLS[key+flag]
	return url
}
