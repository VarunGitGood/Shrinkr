package util

import (
	"fmt"
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

func OpenBrowser(url string, OS string) {
	var cmd *exec.Cmd
	switch OS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "mac":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
