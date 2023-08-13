package repository

import (
	"context"
	"github.com/Uwasaru/Sayka/domain/entity"
)

// SessionRepositoryはセッションに関する永続化を担当します

type SessionRepository interface {
	// GetByIDはIDを指定してセッションを取得します
	GetByID(ctx context.Context, id string) (*entity.Session, error)
	// GetByUserIDはUserIDを指定してセッションを取得します
	GetByUserID(ctx context.Context, userID string) (*entity.Session, error)
	// GetAllは全てのセッションを取得します
	GetAll(ctx context.Context) (*entity.Sessions, error)
	// Createはセッションを作成します
	CreateSession(ctx context.Context, session *entity.Session) error
	// Updateはセッションを更新します
	UpdateSession(ctx context.Context, session *entity.Session) error
	// Deleteはセッションを削除します
	DeleteSession(ctx context.Context, id string) error
}
