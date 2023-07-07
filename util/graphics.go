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

