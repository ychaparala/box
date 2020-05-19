package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	fmt.Println(string(body))
}
