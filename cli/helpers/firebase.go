package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
)

// SignUP returns bool
func SignUP(uname, password string) bool {
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
		PutUserData(body)
		return true
	}
	fmt.Println("Signup to Box App Failed, Invalid email or password")
	return false
}

// Login returns bool
func Login(uname, password string) bool {
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
		return false
	}
	PutUserData(body)
	fmt.Println("Login Successful " + html.UnescapeString("&#"+strconv.Itoa(128077)+";"))
	return true
}

//GetAccessToken returns firebase Access Token
func GetAccessToken() string{
	tokenURL := "https://securetoken.googleapis.com/v1/token"
	key := "AIzaSyBCfZSG0cOs_SNKtW1PG2-LRPE9S3LTcmA"
	row := GetUserData()
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token",row["refreshToken"])

	client := &http.Client{}
	req, _ := http.NewRequest("POST", tokenURL+"?key="+key, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var accessJSON map[string] interface{}
	if err := json.Unmarshal(body, &accessJSON); err != nil {
		panic(err)
	}
	return fmt.Sprint(accessJSON["access_token"])
}

// Logout returns void
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

// LoginStatus returns void
func LoginStatus() (bool, string) {
	row := GetUserData()
	if row["localId"] != "" {
		return true, row["email"]
	}
	return false, ""
}
