package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

// SignUP returns userprofile
func SignUP(uname, password string) {
	url := "https://identitytoolkit.googleapis.com/v1/accounts:signUp"
	key := "AIzaSyBCfZSG0cOs_SNKtW1PG2-LRPE9S3LTcmA"
	requestBody, err := json.Marshal(map[string]string{
		"email":             uname,
		"password":          password,
		"returnSecureToken": "true",
	})
	if err != nil {
		panic(err)
	}
	resp, err := http.Post(url+"?key="+key, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 200 {
		SetUserData(body)
	}
}

// Login returns userprofile
func Login(uname, password string) {
	url := "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword"
	key := "AIzaSyBCfZSG0cOs_SNKtW1PG2-LRPE9S3LTcmA"
	requestBody, err := json.Marshal(map[string]string{
		"email":             uname,
		"password":          password,
		"returnSecureToken": "true",
	})
	if err != nil {
		panic(err)
	}
	resp, err := http.Post(url+"?key="+key, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		fmt.Println("Login Failed invalid email or password")
	} else {
		SetUserData(body)
		fmt.Println("Login Successful")
	}
}

// Logout returns bool
func Logout() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Remove(filepath.Join(home + "/.box/box.db"))
	fmt.Println("Logout Successful")
}

// LoginStatus returns
func LoginStatus() {
	row := GetUserData()
	if row["localId"] != "" {
		fmt.Println("Logged into Box App")
	} else {
		fmt.Println("Logged out of Box App")
	}
}
