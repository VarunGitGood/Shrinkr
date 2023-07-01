/*
Copyright © 2023 Varun Singh varun7singh10@gmail.com

*/
package cmd

import (
	"encoding/json"
	"fmt"
	// "github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to your account using Google OAuth",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting login process...")
		Login()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func GetLoginData() *loginDTO {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://127.0.0.1:3000/shrinkr/login", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var login loginDTO
	err = json.Unmarshal(body, &login)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
	}
	return &login
}

func Login() {
	done := make(chan bool)
	login := GetLoginData()
	state := login.State
	fmt.Println("Click here to login using Google:\n", login.Url)

	// creating a seperate go routine to handle the callback so as to not block the main thread
	go func() {
		http.HandleFunc("/shrinkr/callback", func(w http.ResponseWriter, r *http.Request) {
			HandleCallback(w, r, state, done)
		})

		if err := http.ListenAndServe(":9000", nil); err != nil {
			fmt.Println(err)
			done <- true // Signal completion in case of error
		}
	}()

	<-done
}

func HandleCallback(
	w http.ResponseWriter,
	r *http.Request,
	state string,
	done chan bool,
) error {
	if r.URL.Query().Get("state") != state {
		return fmt.Errorf("invalid oauth state")
	}
	// unblock the Login() call
	fmt.Fprintf(w, "Login Successful!")
	fmt.Println("Login Successful!")
	done <- true
	// return html
	return nil
}

type loginDTO struct {
	Url   string `json:"url"`
	State string `json:"state"`
}
