/*
Copyright © 2023 Varun Singh varun7singh10@gmail.com

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"shrinkr/util"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

const (
	addURL = "http://127.0.0.1:3000/shrinkr/links/addurl"
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
	Long: `Add a new URL to shrinkr to be shortened,
	also can specify the expiration time for the shortened URL
	along with password protection and Expiration Clicks`,
	Run: func(cmd *cobra.Command, args []string) {
		Add()
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
	baseURL := "http://127.0.0.1:3000/shrinkr/"
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	shortURL := util.GenerateShortURL()
	answers.ShortURL = shortURL
	answersJSON, err := json.Marshal(answers)
	client := &http.Client{}
	req, err := http.NewRequest("POST", addURL, bytes.NewBuffer(answersJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+util.GetToken())
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	util.PTextCYAN("Shortened URL: " + baseURL + shortURL)
}
