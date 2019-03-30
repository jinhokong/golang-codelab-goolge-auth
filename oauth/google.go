package oauth

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOAuthConf *oauth2.Config

func init() {
	GoogleOAuthConf = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "http://localhost:1333/auth/callback",
		Scopes: []string{
			"https://googleapis.com/auth/userinfo.email",
			"https://googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

func GoogleAuthorizationURL(state string) string {
	return GoogleOAuthConf.AuthCodeURL(state)
}
