/*
Copyright © 2023 Varun Singh varun7singh10@gmail.com

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// init a http client to make requests to the shrinkr server as it will be used in multiple commands

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "shrinkr",
	Short: "Shinkr",
	Long:  `Shrinkr is a simple to use URL shortener with Google OAuth support and blazing fast performance.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
