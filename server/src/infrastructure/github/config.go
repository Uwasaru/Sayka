package github

import (
	"os"

	"golang.org/x/oauth2"
)

type Github struct {
	Config *oauth2.Config
}

func NewGithub(redirecturl string) *Github {
	return newGithub(redirecturl)
}

func newGithub(redirecturl string) *Github {
	github := &Github{
		Config: &oauth2.Config{
			ClientID:     os.Getenv("GITHUB_ID"),
			ClientSecret: os.Getenv("GITHUB_SECRET"),
			Endpoint:     endpoint,
			Scopes:       []string{"user"},
			RedirectURL:  redirecturl,
		},
	}
	return github
}

var endpoint = oauth2.Endpoint{
	AuthURL:  "https://github.com/login/oauth/authorize",
	TokenURL: "https://github.com/login/oauth/access_token",
}
