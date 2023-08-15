package repository

import (
	"context"
	"time"

	"github.com/Uwasaru/Sayka/domain/entity"
	"golang.org/x/oauth2"
)

type AuthRepository interface {
	StoreState(ctx context.Context, authState *entity.AuthStates) error
	StoreSession(ctx context.Context, sessionId string, userId string) error
	FindStateByState(ctx context.Context, state string) (*entity.AuthStates, error)
	StoreORUpdateToken(ctx context.Context, userId string, token *oauth2.Token) error
	DeleteState(ctx context.Context, state string) error
	GetTokenByUserID(ctx context.Context, userId string) (*oauth2.Token, error)
	UpdateToken(ctx context.Context, userId string, token *oauth2.Token) error
	StoreToken(ctx context.Context, userId string, token *oauth2.Token) error
	GetUserIdFromSession(ctx context.Context, sessionId string) (string, error)
	DeleteSession(ctx context.Context, sessionID string) error
	GetExpiryFromSession(ctx context.Context, sessionId string) (time.Time, error)
}
