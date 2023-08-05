package repository

import (
	"github.com/Uwasaru/Sayka/domain/entity"
)

// SessionRepositoryはセッションに関する永続化を担当します

type SessionRepository interface {
	// GetByIDはIDを指定してセッションを取得します
	GetByID(id string) (*entity.Session, error)
	// GetByUserIDはUserIDを指定してセッションを取得します
	GetByUserID(userID string) (*entity.Session, error)
	// GetAllは全てのセッションを取得します
	GetAll() (*entity.Sessions, error)
	// Createはセッションを作成します
	CreateSession(session *entity.Session) error
	// Updateはセッションを更新します
	UpdateSession(session *entity.Session) error
	// Deleteはセッションを削除します
	DeleteSession(id string) error
}