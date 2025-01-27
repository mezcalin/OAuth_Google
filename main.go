package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Fehler beim Laden der .env-Datei: %v", err)
	}

	oauthConfig.ClientID = os.Getenv("GOOGLE_CLIENT_ID")
	oauthConfig.ClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)

	fmt.Println("Server läuft auf http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(res http.ResponseWriter, req *http.Request) {
	html := `<html><body><a href="/login">Mit Google anmelden</a></body></html>`
	fmt.Fprint(res, html)
}

func handleLogin(res http.ResponseWriter, req *http.Request) {
	url := oauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(res, req, url, http.StatusTemporaryRedirect)
}

func handleCallback(res http.ResponseWriter, req *http.Request) {
	// Überprüfen des OAuth-State-Strings
	if req.FormValue("state") != oauthStateString {
		http.Error(res, "Ungültiger OAuth-State", http.StatusBadRequest)
		return
	}

	// Den Auth-Code holen
	code := req.FormValue("code")
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(res, "Code-Austausch fehlgeschlagen", http.StatusInternalServerError)
		return
	}

	// Benutzerinformationen abrufen
	client := oauthConfig.Client(context.Background(), token)
	userInfoResp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(res, "Fehler beim Abrufen der Benutzerinformationen", http.StatusInternalServerError)
		return
	}
	defer userInfoResp.Body.Close()

	userInfo := struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}{}
	if err := json.NewDecoder(userInfoResp.Body).Decode(&userInfo); err != nil {
		http.Error(res, "Fehler beim Decodieren der Benutzerinformationen", http.StatusInternalServerError)
		return
	}

	// Benutzerinformationen anzeigen
	fmt.Fprintf(res, "Hallo %s, deine E-Mail ist: %s", userInfo.Name, userInfo.Email)
}
