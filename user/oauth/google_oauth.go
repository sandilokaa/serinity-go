package oauth

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
	RedirectURL:  "http://localhost:8080/api/v1/session/oauth/callback",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

func GetGoogleOauthConfig() *oauth2.Config {
	return googleOauthConfig
}

func GetLoginGoogleURL(state string) string {
	return googleOauthConfig.AuthCodeURL(state)
}
