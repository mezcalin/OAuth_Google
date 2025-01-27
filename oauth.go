package main

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var oauthConfig = &oauth2.Config{
	ClientID:     "",                               // from https://console.developers.google.com/project/<your-project-id>/apiui/credential
	ClientSecret: "",                               // from https://console.developers.google.com/project/<your-project-id>/apiui/credential
	RedirectURL:  "http://localhost:8080/callback", // Redirect URI
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

var oauthStateString = "randomString" // Sicherheitsmechanismus
