package persistence

import (
	"context"

	"github.com/Uwasaru/Sayka/domain/entity"
	"github.com/Uwasaru/Sayka/domain/repository"
	"github.com/Uwasaru/Sayka/infrastructure/database"
	d "github.com/Uwasaru/Sayka/infrastructure/persistence/dto"
)

var _ repository.SessionRepository = &SessionRepository{}

// SessionRepositoryはセッションに関する永続化を担当します
type SessionRepository struct {
	conn *database.Conn
}

// NewSessionRepositoryは新しいSessionRepositoryを初期化し構造体のポインタを返します
func NewSessionRepository(conn *database.Conn) *SessionRepository {
	return &SessionRepository{
		conn: conn,
	}
}

// GetByIDはIDを指定してセッションを取得します
func (sr *SessionRepository) GetByID(ctx context.Context, id string) (*entity.Session, error) {
	query := `
	SELECT *
	FROM sessions
	WHERE id = ?
	`
	var dto d.SessionDto
	err := sr.conn.DB.GetContext(ctx, &dto, query, id)
	if err != nil {
		return nil, err
	}
	return d.SessionDtoToEntity(&dto), nil
}

// GetByUserIDはUserIDを指定してセッションを取得します
func (sr *SessionRepository) GetByUserID(ctx context.Context, userID string) (*entity.Session, error) {
	query := `
	SELECT *
	FROM sessions
	WHERE user_id = ?
	`
	var dto d.SessionDto
	err := sr.conn.DB.GetContext(ctx, &dto, query, userID)
	if err != nil {
		return nil, err
	}
	return d.SessionDtoToEntity(&dto), nil
}

// GetAllは全てのセッションを取得します
func (sr *SessionRepository) GetAll(ctx context.Context) (*entity.Sessions, error) {
	query := `
	SELECT *
	FROM sessions
	`
	rows, err := sr.conn.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var sessions entity.Sessions
	for rows.Next() {
		var dto d.SessionDto
		err := rows.Scan(&dto)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, d.SessionDtoToEntity(&dto))
	}
	return &sessions, nil
}

// Createはセッションを作成します
func (sr *SessionRepository) CreateSession(ctx context.Context, session *entity.Session) error {
	query := `
	INSERT INTO sessions (session_id, user_id, token, token_expire)
	VALUES (?, ?, ?, ?)
	`
	_, err := sr.conn.DB.ExecContext(ctx, query, session.SessionID, session.UserID, session.Token, session.TokenExpire)
	if err != nil {
		return err
	}
	return nil
}

// Updateはセッションを更新します
func (sr *SessionRepository) UpdateSession(ctx context.Context, session *entity.Session) error {
	query := `
	UPDATE sessions
	SET token = ?, token_expire = ?
	WHERE id = ?
	`
	_, err := sr.conn.DB.ExecContext(ctx, query, session.Token, session.TokenExpire, session.ID)
	if err != nil {
		return err
	}
	return nil
}

// Deleteはセッションを削除します
func (sr *SessionRepository) DeleteSession(ctx context.Context, id string) error {
	query := `
	DELETE FROM sessions
	WHERE id = ?
	`
	_, err := sr.conn.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
