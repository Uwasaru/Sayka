package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Uwasaru/Sayka/domain/entity"
	"github.com/Uwasaru/Sayka/domain/repository"
	mycontext "github.com/Uwasaru/Sayka/utils/context"
	"golang.org/x/oauth2"

	"github.com/Uwasaru/Sayka/utils"
)

type Auth struct {
	authRepo   repository.AuthRepository
	githubRepo repository.GithubRepository
	userRepo   repository.UserRepository
}
type IAuthUsecase interface {
	GetAuthURL(ctx context.Context, redirectURL string) (url string, state string, err error)
	Authorization(ctx context.Context, state, code string) (string, string, error)
	DeleteSession(ctx context.Context, sessionID string) error
	CheckSessionExpiry(ctx context.Context, sessionID string) (bool, error)
	GetUserIdFromSession(ctx context.Context, sessionId string) (int, error)
	GetTokenByUserId(ctx context.Context, userId int) (*oauth2.Token, error)
	RefreshAccessToken(ctx context.Context, userId int, token *oauth2.Token) (*oauth2.Token, error)
}

func NewAuthUsecase(ra repository.AuthRepository, rg repository.GithubRepository, ur repository.UserRepository) IAuthUsecase {
	return &Auth{authRepo: ra, githubRepo: rg, userRepo: ur}
}

func (uc *Auth) GetAuthURL(ctx context.Context, redirectURL string) (url string, resstate string, err error) {
	state := utils.Ulid()
	st := &entity.AuthStates{State: state, RedirectURL: redirectURL}
	if err := uc.authRepo.StoreState(ctx, st); err != nil {
		return "", "", fmt.Errorf("storeState: %w", err)
	}
	return uc.githubRepo.GetAuthURL(state), state, nil
}

func (uc *Auth) Authorization(ctx context.Context, state, code string) (string, string, error) {
	storedState, err := uc.authRepo.FindStateByState(ctx, state)
	if err != nil {
		return "", "", fmt.Errorf("findStateByState: %w", err)
	}
	// codeを使ってtokenを取得
	token, err := uc.githubRepo.Exchange(ctx, code)
	if err != nil {
		return storedState.RedirectURL, "", fmt.Errorf("exchange: %w", err)
	}
	// ctxにtokenをセット
	ctx = mycontext.SetToken(ctx, token)
	userId, err := uc.createUserIfNotExists(ctx)
	if err != nil {
		return storedState.RedirectURL, "", fmt.Errorf("createUserIfNotExists: %w", err)
	}
	// tokenをdbに保存
	if err := uc.authRepo.StoreORUpdateToken(ctx, userId, token); err != nil {
		return storedState.RedirectURL, "", fmt.Errorf("storeORUpdateToken: %w", err)
	}
	// sessionをdbに保存
	sessionID := utils.Ulid()
	if err := uc.authRepo.StoreSession(ctx, sessionID, userId); err != nil {
		return storedState.RedirectURL, "", fmt.Errorf("storeSession: %w", err)
	}
	// stateを削除
	if err := uc.authRepo.DeleteState(ctx, state); err != nil {
		return storedState.RedirectURL, sessionID, nil
	}
	return storedState.RedirectURL, sessionID, nil
}

func (uc *Auth) createUserIfNotExists(ctx context.Context) (int, error) {
	// githubからユーザー情報を取得
	user, err := uc.githubRepo.GetMe(ctx)
	if err != nil {
		return 0, fmt.Errorf("getMe: %w", err)
	}

	// ユーザが既に存在するか確認
	_, err = uc.userRepo.GetUser(context.Background(), user.ID)
	if err != nil && err == sql.ErrNoRows {
		user, err = uc.userRepo.CreateUser(context.Background(), user)
		if err != nil {
			return 0, err
		}
	}
	if err != nil {
		return 0, err
	}
	return user.ID, err
}

func (uc *Auth) DeleteSession(ctx context.Context, sessionID string) error {
	if sessionID == "" {
		return fmt.Errorf("sessionid is empty")
	}
	err := uc.authRepo.DeleteSession(ctx, sessionID)
	return err
}

func (uc *Auth) CheckSessionExpiry(ctx context.Context, sessionID string) (bool, error) {
	if sessionID == "" {
		return false, fmt.Errorf("sessionid is empty")
	}
	expiry, err := uc.authRepo.GetExpiryFromSession(ctx, sessionID)
	if err != nil {
		return false, fmt.Errorf("GetExpiryFromSession: %w", err)
	}
	if expiry.Before(time.Now()) {
		return false, nil
	}
	return true, nil
}

func (uc *Auth) GetUserIdFromSession(ctx context.Context, sessionId string) (int, error) {
	userId, err := uc.authRepo.GetUserIdFromSession(ctx, sessionId)
	if err != nil {
		return 0, fmt.Errorf("getUserIDFromSession: %w", err)
	}
	return userId, nil
}

func (uc *Auth) GetTokenByUserId(ctx context.Context, userId int) (*oauth2.Token, error) {
	token, err := uc.authRepo.GetTokenByUserID(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("getTokenByUserID: %w", err)
	}
	return token, nil
}
func (uc *Auth) RefreshAccessToken(ctx context.Context, userId int, token *oauth2.Token) (*oauth2.Token, error) {
	if token.Valid() {
		return token, nil
	}
	newToken, err := uc.githubRepo.Refresh(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("refresh: %w", err)
	}
	if err := uc.authRepo.StoreORUpdateToken(ctx, userId, newToken); err != nil {
		return nil, fmt.Errorf("storeORUpdateToken: %w", err)
	}
	return newToken, nil
}
