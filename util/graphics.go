package util

import (
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func Spinner(prefix string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[26], 400*time.Millisecond)
	s.Color("cyan")
	startingMessage := color.New(color.FgCyan, color.Bold)
	s.Prefix = startingMessage.Sprint(prefix)
	return s
}

func Text(text string) string {
	blue := color.New(color.FgCyan, color.Bold)
	return blue.Sprint(text)
}

func PTextCYAN(text string) {
	blue := color.New(color.FgCyan, color.Bold)
	blue.Println(text)
}

func PTextGREEN(text string) {
	green := color.New(color.FgGreen, color.Bold)
	green.Println(text)
}

func PTextRED(text string) {
	red := color.New(color.FgRed, color.Bold)
	red.Println(text)
}
