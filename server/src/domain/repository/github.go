package repository

import (
	"context"

	"github.com/Uwasaru/Sayka/domain/entity"
	"golang.org/x/oauth2"
)

type GithubRepository interface {
	// auth
	GetAuthURL(state string) string
	Exchange(ctx context.Context, code string) (*oauth2.Token, error)
	Refresh(ctx context.Context, token *oauth2.Token) (*oauth2.Token, error)
	// user
	GetMe(ctx context.Context) (*entity.User, error)
}
