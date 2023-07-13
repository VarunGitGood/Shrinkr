/*
Copyright © 2023 Varun Singh varun7singh10@gmail.com
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"runtime"
	"shrinkr/util"

	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to your account using Google OAuth",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		Login()
	},
}

var (
	loginURL = util.GetURL("loginURL")
	tokenURL = util.GetURL("tokenURL")
)

type loginDTO struct {
	Url   string `json:"url"`
	State string `json:"state"`
}

type tokenDTO struct {
	Token string `json:"token"`
	Name  string `json:"name"`
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func GetLoginData() *loginDTO {
	resp, err := util.Get(loginURL, "", false)
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
	util.OpenBrowser(login.Url, runtime.GOOS)

	s := util.Spinner("Click on the link above to login")
	s.Start()
	fmt.Print("\n")
	util.PTextGREEN(login.Url + "\n")

	go func() {
		http.HandleFunc("/shrinkr/callback", func(w http.ResponseWriter, r *http.Request) {
			HandleCallback(w, r, state, done)
		})
		server := &http.Server{Addr: ":9000"}
		if err := server.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
		<-done
		if err := server.Close(); err != nil {
			fmt.Println("Error shutting down the server:", err)
		}
	}()
	<-done
	s.Stop()
}

func WriteToFile(token string) {
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
	if r.URL.Query().Get("state") != state {
		return
	}
	resp, err := util.Get(tokenURL, "code="+r.URL.Query().Get("code"), false)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var tokenData tokenDTO
	err = json.Unmarshal(body, &tokenData)
	WriteToFile(tokenData.Token)
	fmt.Println("\nLogged in as " + util.Text(tokenData.Name) + "\n")
	http.ServeFile(w, r, "cmd/public/check.html")
	// Unblock the Login() call
	done <- true
}
