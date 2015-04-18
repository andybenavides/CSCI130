package main

import (
	"code.google.com/p/goauth2/oauth"
	"code.google.com/p/google-api-go-client/drive/v2"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Settings for authorization.
var config = &oauth.Config{
	ClientId:     "706961739316-ihs3q7f5obrc65lsps3hrq1ni5cfdt57.apps.googleusercontent.com",
	ClientSecret: "f7ef265nzJglTDS4yqf-3ilP",
	Scope:        "https://www.googleapis.com/auth/drive",
	RedirectURL:  "urn:ietf:wg:oauth:2.0:oob",
	AuthURL:      "https://accounts.google.com/o/oauth2/auth",
	TokenURL:     "https://accounts.google.com/o/oauth2/token",
}

// Uploads a file to Google Drive
func main() {

	// Generate a URL to visit for authorization.
	authUrl := config.AuthCodeURL("state")
	log.Printf("Go to the following link in your browser: %v\n", authUrl)
	t := &oauth.Transport{
		Config:    config,
		Transport: http.DefaultTransport,
	}

	// Read the code, and exchange it for a token.
	log.Printf("Enter verification code: ")
	var code string
	fmt.Scanln(&code)
	_, err := t.Exchange(code)
	if err != nil {
		log.Fatalf("An error occurred exchanging the code: %v\n", err)
	}

	// Create a new authorized Drive client.
	svc, err := drive.New(t.Client())
	if err != nil {
		log.Fatalf("An error occurred creating Drive client: %v\n", err)
	}

	// Define the metadata for the file we are going to create.
	f := &drive.File{
		Title:       "My Document",
		Description: "My test document",
	}

	// Read the file data that we are going to upload.
	m, err := os.Open("document.txt")
	if err != nil {
		log.Fatalf("An error occurred reading the document: %v\n", err)
	}

	// Make the API request to upload metadata and file data.
	r, err := svc.Files.Insert(f).Media(m).Do()
	if err != nil {
		log.Fatalf("An error occurred uploading the document: %v\n", err)
	}
	log.Printf("Created: ID=%v, Title=%v\n", r.Id, r.Title)
}
