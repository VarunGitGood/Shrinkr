/*
Copyright © 2023 Varun Singh varun7singh10@gmail.com

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"shrinkr/util"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

const (
	addURL  = "https://shrinkr-da1u.onrender.com/shrinkr/links/addurl"
	baseURL = "https://shrinkr-da1u.onrender.com/shrinkr/"
)

var qs = []*survey.Question{
	{
		Name:     "longURL",
		Prompt:   &survey.Input{Message: "Enter the URL to be shortened: "},
		Validate: survey.Required,
	},
	{
		Name:     "description",
		Prompt:   &survey.Input{Message: "Enter the description for the shortened URL: "},
		Validate: survey.Required,
	},
	{
		Name: "expiration",
		Prompt: &survey.Input{Message: "Enter the expiration time for the shortened URL: ",
			Help:    "Leave blank for no expiration",
			Default: "0"},
		Validate: func(val interface{}) error {
			if val.(string) < "0" {
				return errors.New("Expiration time cannot be negative")
			}
			if !util.IsInt(val.(string)) {
				return errors.New("Enter a valid integer")
			}
			return nil
		},
	},
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new URL to shrinkr to be shortened",
	Long: `Add a new URL to shrinkr to be shortened,also can specify the expiration time for the shortened URL
	along with password protection and Expiration Clicks`,
	Run: func(cmd *cobra.Command, args []string) {
		if util.Authenticated() {
			Add()
		} else {
			util.PTextCYAN("Please login first")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func Add() {
	answers := struct {
		LongURL     string `json:"longURL" survey:"longURL"`
		Description string `json:"description" survey:"description"`
		Expiration  int    `json:"expiration" survey:"expiration"`
		ShortURL    string `json:"shortURL" survey:"shortURL"`
	}{}
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	shortURL := util.GenerateShortURL()
	answers.ShortURL = shortURL
	answersJSON, err := json.Marshal(answers)
	resp, err := util.Post(addURL, bytes.NewReader(answersJSON), true)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	util.PTextCYAN("Shortened URL: " + baseURL + shortURL)
}
