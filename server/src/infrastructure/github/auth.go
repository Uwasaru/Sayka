package github

import (
	"context"
	"fmt"
	"time"

	"github.com/Uwasaru/Sayka/domain/repository"
	"github.com/patrickmn/go-cache"
	"golang.org/x/oauth2"
)

type Client struct {
	auth  *Github
	cache *cache.Cache
}

func NewClient(redirecturl string) repository.GithubRepository {
	auth := NewGithub(redirecturl)
	return &Client{auth: auth, cache: cache.New(10*time.Minute, 20*time.Minute)}
}

func (c *Client) GetAuthURL(state string) string {
	return c.auth.Config.AuthCodeURL(state)
}

func (c *Client) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := c.auth.Config.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("exchange: %w", err)
	}
	return token, nil
}

func (c *Client) Refresh(ctx context.Context, token *oauth2.Token) (*oauth2.Token, error) {
	newtoken, err := c.auth.Config.TokenSource(ctx, token).Token()
	if err != nil {
		return nil, fmt.Errorf("tokenSource: %w", err)
	}
	return newtoken, err
}
