package util

import (
	"os/exec"
	"strings"

	"github.com/google/uuid"
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

func OpenBrowser(url string) {
	cmd := exec.Command("xdg-open", url)
	cmd.Run()
}
