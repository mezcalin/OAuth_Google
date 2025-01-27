# small web server with Google Cloud OAuth2

## Getting started
go get github.com/joho/godotenv

## Google SSO via OAuth2
### Anlagen der ClientID
-> Registrieren bei https://cloud.google.com
-> Unter https://console.cloud.google.com/apis/credentials einen API-Schlüssel anlegen und dann anschließend eine OAuth 2.0-Client-ID

## Credentials hinterlegen
example.env kopieren in .env und die ClientId sowie den Clientschlüssel vom Google SSO eintragen

## Starten des Servers
*go run .*
Server läuft unter localhost:8080
