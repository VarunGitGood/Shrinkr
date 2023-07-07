/*
Copyright © 2023 Varun Singh varun7singh10@gmail.com
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"shrinkr/util"

	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to your account using Google OAuth",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		Login()
	},
}

const (
	loginURL = "http://127.0.0.1:3000/shrinkr/login"
	tokenURL = "http://127.0.0.1:3000/shrinkr/token"

)

type loginDTO struct {
	Url   string `json:"url"`
	State string `json:"state"`
}

type tokenDTO struct {
	Token string `json:"token"`
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
	req, err := http.NewRequest("GET", loginURL , nil)
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

	s := util.Spinner("Click on the link above to login ")
	s.Start()
	fmt.Println(login.Url + "\n")

	// creating a seperate go routine to handle the callback so as to not block the main thread
	go func() {
		// Handlers
		http.HandleFunc("/shrinkr/callback", func(w http.ResponseWriter, r *http.Request) {
			HandleCallback(w, r, state, done)
		})
		http.HandleFunc("/shrinkr/test", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Test")
		})
		server := &http.Server{Addr: ":9000"}
		if err := server.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
		// Wait for completion or other termination condition
		<-done
		if err := server.Close(); err != nil {
			fmt.Println("Error shutting down the server:", err)
		}
	}()
	<-done
	s.Stop()
}

func WriteToFile(token string) {
	// store the token in the .env file for future use with varible name TOKEN
	err := ioutil.WriteFile("secrets.env", []byte("TOKEN="+token), 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func HandleCallback(
	w http.ResponseWriter,
	r *http.Request,
	state string,
	done chan bool,
) {
	// Check the state that was sent back
	if r.URL.Query().Get("state") != state {
		return
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", tokenURL, nil)
	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("code", r.URL.Query().Get("code"))
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var token tokenDTO
	err = json.Unmarshal(body, &token)
	WriteToFile(token.Token)
	fmt.Println("Login Successful!")
	http.ServeFile(w, r, "cmd/public/check.html")
	// Unblock the Login() call
	done <- true
}


