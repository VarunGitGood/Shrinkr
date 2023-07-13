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

var (
	addURL   = util.GetURL("addURL")
	baseURL  = util.GetURL("baseURL")
	checkURL = util.GetURL("checkURL")
	isCustom bool
)

var cqs = []*survey.Question{
	{
		Name: "shortURL",
		Prompt: &survey.Input{Message: "Enter the short URL: ",
			Help: "Leave blank for a random short URL"},
		Validate: func(val interface{}) error {
			if val.(string) == "" {
				return nil
			}
			if !CheckURL(val.(string)) {
				return errors.New("That short URL is already taken")
			}
			return nil
		},
	},
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
		isCustom, _ = cmd.Flags().GetBool("custom")
		if util.Authenticated() {
			Add(isCustom)
		} else {
			util.PTextCYAN("Please login first")
		}
	},
}

func CheckURL(url string) bool {
	var Response struct {
		status string `json:"status"`
	}
	resp, err := util.Get(checkURL+url, url, false)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&Response)
	if err != nil {
		fmt.Println(err)
	}
	if Response.status == "success" {
		return true
	} else {
		return false
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
	// -c flag for custom short URL provides a boolean value for isCustom
	addCmd.Flags().BoolP("custom", "c", false, "Add a custom short URL")
}

func Add(cust bool) {
	answers := struct {
		LongURL     string `json:"longURL" survey:"longURL"`
		Description string `json:"description" survey:"description"`
		Expiration  int    `json:"expiration" survey:"expiration"`
		ShortURL    string `json:"shortURL" survey:"shortURL"`
	}{}
	var url string
	if cust {
		err := survey.Ask(cqs, &answers)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		url = answers.ShortURL
	} else {
		err := survey.Ask(qs, &answers)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		shortURL := util.GenerateShortURL()
		answers.ShortURL = shortURL
		url = shortURL
	}
	answersJSON, err := json.Marshal(answers)
	resp, err := util.Post(addURL, bytes.NewReader(answersJSON), true)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	util.PTextCYAN("Shortened URL: " + baseURL + url)
}
