package util

import (
	"github.com/google/uuid"
	"strings"
)

func GenerateShortURL() string {
	
	uuid := uuid.New()
	strings.Replace(uuid.String(), "-", "", -1)
	return uuid.String()[:8]
}

func IsInt(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}


