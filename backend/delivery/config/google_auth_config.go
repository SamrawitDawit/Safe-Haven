package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GoogleConfig() *oauth2.Config {
	LoadEnv()
	var (
		GoogleOAuthConfig = &oauth2.Config{
			ClientID:     ENV.GOOGLE_CLIENT_ID,
			ClientSecret: ENV.GOOGLE_CLIENT_SECRET,
			RedirectURL:  "http://localhost:8080/auth/google/callback",
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile", "openid"},
			Endpoint:     google.Endpoint,
		}
	)
	return GoogleOAuthConfig

}

var GoogleOAuthConfig = GoogleConfig()
