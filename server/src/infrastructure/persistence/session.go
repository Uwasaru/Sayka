package persistence

import (
	"database/sql"
	"github.com/Uwasaru/Sayka/domain/entity"
	"github.com/Uwasaru/Sayka/domain/repository"
)

var _ repository.SessionRepository = &SessionRepository{}

// SessionRepositoryはセッションに関する永続化を担当します
type SessionRepository struct {
	db *sql.DB
}

// NewSessionRepositoryは新しいSessionRepositoryを初期化し構造体のポインタを返します
func NewSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{
		db: db,
	}
}

// GetByIDはIDを指定してセッションを取得します
func (sr *SessionRepository) GetByID(id string) (*entity.Session, error) {
	session := &entity.Session{}
	err := sr.db.QueryRow("SELECT * FROM sessions WHERE id = ?", id).Scan(&session.ID, &session.SessionID, &session.UserID, &session.Token, &session.TokenExpire)
	if err != nil {
		return nil, err
	}
	return session, nil
}

// GetByUserIDはUserIDを指定してセッションを取得します
func (sr *SessionRepository) GetByUserID(userID string) (*entity.Session, error) {
	session := &entity.Session{}
	err := sr.db.QueryRow("SELECT * FROM sessions WHERE user_id = ?", userID).Scan(&session.ID, &session.SessionID, &session.UserID, &session.Token, &session.TokenExpire)
	if err != nil {
		return nil, err
	}
	return session, nil
}

// GetAllは全てのセッションを取得します
func (sr *SessionRepository) GetAll() (*entity.Sessions, error) {
	rows, err := sr.db.Query("SELECT * FROM sessions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sessions := &entity.Sessions{}
	for rows.Next() {
		session := &entity.Session{}
		err := rows.Scan(&session.ID, &session.SessionID, &session.UserID, &session.Token, &session.TokenExpire)
		if err != nil {
			return nil, err
		}
		*sessions = append(*sessions, session)
	}
	return sessions, nil
}

// Createはセッションを作成します
func (sr *SessionRepository) CreateSession(session *entity.Session) error {
	_, err := sr.db.Exec("INSERT INTO sessions (session_id, user_id, token, token_expire) VALUES (?, ?, ?, ?)", session.SessionID, session.UserID, session.Token, session.TokenExpire)
	if err != nil {
		return err
	}
	return nil
}

// Updateはセッションを更新します
func (sr *SessionRepository) UpdateSession(session *entity.Session) error {
	_, err := sr.db.Exec("UPDATE sessions SET session_id = ?, user_id = ?, token = ?, token_expire = ? WHERE id = ?", session.SessionID, session.UserID, session.Token, session.TokenExpire, session.ID)
	if err != nil {
		return err
	}
	return nil
}

// Deleteはセッションを削除します
func (sr *SessionRepository) DeleteSession(id string) error {
	_, err := sr.db.Exec("DELETE FROM sessions WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
