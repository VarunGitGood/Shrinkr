/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"shrinkr/util"

	"github.com/spf13/cobra"
)

const (
	allURLs = "http://127.0.0.1:3000/shrinkr/links/mappings"
	linkURL = "http://127.0.0.1:3000/shrinkr/links/"
	userURL = "http://127.0.0.1:3000/shrinkr/user/info"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "a set of user commands",
	Long: `User commands such as :
	- Info
	- All URLs
	- Delete URL
	- Info URL`,
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		info, _ := cmd.Flags().GetBool("info")
		deleteURL, _ := cmd.Flags().GetString("delete")
		infoURL, _ := cmd.Flags().GetString("info-url")
		if all {
			AllURLs()
		} else if info {
			UserInfo()
		} else if deleteURL != "" {
			DeleteURL(deleteURL)
		} else if infoURL != "" {
			InfoURL(infoURL)
		} else {
			util.PTextCYAN("Please specify a valid flag for user command use --help for more info")
		}
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
	// -a flag for all urls created by the user
	userCmd.Flags().BoolP("all", "a", false, "Get all urls created by the user")

	// -i flag for user info
	userCmd.Flags().BoolP("info", "i", false, "Get user info")

	// -d flag for delete URL
	userCmd.Flags().StringP("delete", "d", "", "Delete a URL by key")

	// -iu flag for info of a particular URL
	userCmd.Flags().StringP("info-url", "u", "", "Get info of a particular URL by key")
}

func AllURLs() {
	resp, err := util.Get(allURLs, "", true)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var all struct {
		Data   []LinkInfo `json:"data"`
		Status string     `json:"status"`
	}
	err = json.Unmarshal(body, &all)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
	}
	util.PTextCYAN("All URLs created by you: \n")
	for _, v := range all.Data {
		util.PTextGREEN("{")
		util.PTextCYAN("	Key : " + v.Key)
		util.PTextCYAN("	Long URL : " + v.LongURL)
		util.PTextCYAN("	Description : " + v.Description)
		util.PTextCYAN("	Expiration : " + fmt.Sprint(v.Expiration))
		util.PTextCYAN("	Created At : " + v.Created)
		util.PTextGREEN("},")
	}
}

func UserInfo() {
	resp, err := util.Get(userURL, "", true)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var Info struct {
		Data   User   `json:"data"`
		Status string `json:"status"`
	}
	err = json.Unmarshal(body, &Info)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
	}
	var user User
	user = Info.Data
	util.PTextCYAN("User Info: \n")
	util.PTextGREEN("{")
	util.PTextCYAN("	Verified Email : " + user.Username)
	util.PTextCYAN("	Joined On : " + user.Joined)
	util.PTextCYAN("	Total Links Created : " + fmt.Sprint(user.Links))
	util.PTextGREEN("}")
}

func DeleteURL(deleteURL string) {
	resp, err := util.Delete(linkURL+deleteURL, true)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var delete struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	}
	err = json.Unmarshal(body, &delete)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
	}
	if delete.Status == "success" {
		util.PTextCYAN(delete.Message)
	} else {
		util.PTextRED(delete.Message)
	}
}

func InfoURL(infoURL string) {
	resp, err := util.Get(linkURL+infoURL, "", true)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var info struct {
		Data   LinkInfo `json:"data"`
		Status string   `json:"status"`
	}
	err = json.Unmarshal(body, &info)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
	}
	if info.Status == "success" {
		util.PTextCYAN("Info for URL: " + infoURL)
		util.PTextGREEN("{")
		util.PTextCYAN("	Key : " + info.Data.Key)
		util.PTextCYAN("	Long URL : " + info.Data.LongURL)
		util.PTextCYAN("	Description : " + info.Data.Description)
		util.PTextCYAN("	Expiration : " + fmt.Sprint(info.Data.Expiration))
		util.PTextCYAN("	Created At : " + info.Data.Created)
		util.PTextGREEN("}")
	} else {
		util.PTextRED("Error: " + info.Data.Key + " does not exist")
	}
}
