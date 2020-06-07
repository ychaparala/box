package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type (
	// GoogleDrive struct
	GoogleDrive struct {
		fileStorage map[string]string
	}
	// OneDrive struct
	OneDrive struct {
		fileStorage map[string]string
	}
	// Usage struct
	Usage struct {
		Email             string `json:"email"`
		Limit             int64  `json:"limit"`
		Usage             int64  `json:"usage"`
		UsageInDrive      int64  `json:"usageInDrive"`
		UsageInDriveTrash int64  `json:"usageInDriveTrash"`
	}
	// FileStorage defines list of available api's
	FileStorage interface {
		Connect() (*http.Client, bool)
		Disconnect()
		CreateFile()
		UpdateFile()
		DeleteFile()
		Usage() string
	}
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.


	// Get Token from web
	url := "https://us-central1-box-app-80870.cloudfunctions.net/app/getAccessToken"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+ GetAccessToken())
	resp, _ := client.Do(req)
	var tok *oauth2.Token
	if resp.StatusCode != 200 {
		tok = getTokenFromWeb(config)
		url := "https://us-central1-box-app-80870.cloudfunctions.net/app/putToken"
		client := &http.Client{}
		row := GetUserData()
		requestBody, err := json.Marshal(map[string]string{
			"email": row["email"],
			"provider": "googleDrive",
			"refresh_token": tok.AccessToken,
		})
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
		req.Header.Set("Authorization", "Bearer "+ GetAccessToken())
		req.Header.Set("Content-Type", "application/json")
		_ , err = client.Do(req)
		if err != nil {
			panic(err)
		}
	} else {
		defer resp.Body.Close()
		tok, _ = tokenFromFile("/Users/yuga/Downloads/token.json")
	}

	return config.Client(context.Background(), tok)
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// Connect func for GoogleDrive
func (gd GoogleDrive) Connect() (*http.Client, bool) {
	ls, _:=LoginStatus()
	if ls {
		url := "https://us-central1-box-app-80870.cloudfunctions.net/app/getClientConfig"

		client := &http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Bearer "+ GetAccessToken())
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}


		// If modifying these scopes, delete your previously saved token.json.
		config, err := google.ConfigFromJSON(b, drive.DriveAppdataScope)
		if err != nil {
			log.Fatalf("Unable to parse client secret file to config: %v", err)
		}
		//Google oauth client
		client = getClient(config)
		return client, true
	}
	return nil, false
}

// Usage func for GoogleDrive
func (gd GoogleDrive) Usage() *drive.About {
	client , success := gd.Connect()
	if success {
		srv, err := drive.New(client)
		if err != nil {
			log.Fatalf("Unable to retrieve Drive client: %v", err)
		}
		r, _ := srv.About.Get().Fields("storageQuota,user").Do()
		fmt.Println("Usage: "+strconv.Itoa(int(r.StorageQuota.Usage)))
		fmt.Println("Limit: "+strconv.Itoa(int(r.StorageQuota.Limit)))
		fmt.Println("UsageInDrive: "+strconv.Itoa(int(r.StorageQuota.UsageInDrive)))
		return r
	}
	return nil
}
