package usecase

import (
	"context"
	"github.com/Uwasaru/Sayka/domain/entity"
	"github.com/Uwasaru/Sayka/domain/repository"
)

var _ ISessionUsecase = &SessionUsecase{}

// SessionUsecaseはセッションに関するユースケースを担当します
type SessionUsecase struct {
	sr repository.SessionRepository
}

// ISessionUsecaseはセッションに関するユースケースを担当します
type ISessionUsecase interface {
	GetByID(ctx context.Context, id string) (*entity.Session, error)
	GetByUserID(ctx context.Context, userID string) (*entity.Session, error)
	GetAll(ctx context.Context) (*entity.Sessions, error)
	CreateSession(ctx context.Context, session *entity.Session) error
	UpdateSession(ctx context.Context, session *entity.Session) error
	DeleteSession(ctx context.Context, id string) error
}

// NewSessionUsecaseは新しいSessionUsecaseを初期化し構造体のポインタを返します
func NewSessionUsecase(sr repository.SessionRepository) ISessionUsecase {
	return &SessionUsecase{
		sr: sr,
	}
}

// GetByIDはIDを指定してセッションを取得します
func (su *SessionUsecase) GetByID(ctx context.Context, id string) (*entity.Session, error) {
	return su.sr.GetByID(ctx, id)
}

// GetByUserIDはUserIDを指定してセッションを取得します
func (su *SessionUsecase) GetByUserID(ctx context.Context, userID string) (*entity.Session, error) {
	return su.sr.GetByUserID(ctx, userID)
}

// GetAllは全てのセッションを取得します
func (su *SessionUsecase) GetAll(ctx context.Context) (*entity.Sessions, error) {
	return su.sr.GetAll(ctx)
}

// CreateSessionはセッションを作成します
func (su *SessionUsecase) CreateSession(ctx context.Context, session *entity.Session) error {
	return su.sr.CreateSession(ctx, session)
}

// UpdateSessionはセッションを更新します
func (su *SessionUsecase) UpdateSession(ctx context.Context, session *entity.Session) error {
	return su.sr.UpdateSession(ctx, session)
}

// DeleteSessionはセッションを削除します
func (su *SessionUsecase) DeleteSession(ctx context.Context, id string) error {
	return su.sr.DeleteSession(ctx, id)
}
