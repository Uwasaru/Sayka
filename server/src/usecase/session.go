package usecase

import (
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

// NewSessionUsecaseは新しいSessionUsecaseを初期化し構造体のポインタを返します
func NewSessionUsecase(sr repository.SessionRepository) ISessionUsecase {
	return &SessionUsecase{
		sr: sr,
	}
}

// GetByIDはIDを指定してセッションを取得します
func (su *SessionUsecase) GetByID(id string) (*entity.Session, error) {
	return su.sr.GetByID(id)
}

// GetByUserIDはUserIDを指定してセッションを取得します
func (su *SessionUsecase) GetByUserID(userID string) (*entity.Session, error) {
	return su.sr.GetByUserID(userID)
}

// GetAllは全てのセッションを取得します
func (su *SessionUsecase) GetAll() (*entity.Sessions, error) {
	return su.sr.GetAll()
}

// CreateSessionはセッションを作成します
func (su *SessionUsecase) CreateSession(session *entity.Session) error {
	return su.sr.CreateSession(session)
}

// UpdateSessionはセッションを更新します
func (su *SessionUsecase) UpdateSession(session *entity.Session) error {
	return su.sr.UpdateSession(session)
}

// DeleteSessionはセッションを削除します
func (su *SessionUsecase) DeleteSession(id string) error {
	return su.sr.DeleteSession(id)
}
