package oauth

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

var googleOauthConfig = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
	RedirectURL:  "http://localhost:8080/api/v1/sessions/oauth/callback",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

func GetGoogleOauthConfig() *oauth2.Config {
	return googleOauthConfig
}

func GetLoginGoogleURL(state string) string {
	url := googleOauthConfig.AuthCodeURL(state)
	return url
}
