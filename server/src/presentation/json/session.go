package json

import (
	"github.com/Uwasaru/Sayka/domain/entity"
	"time"
)

type SessionJson struct {
	ID          string    `json:"id"`
	SessionID   string    `json:"session_id"`
	UserID      string    `json:"user_id"`
	Token       string    `json:"token"`
	TokenExpire time.Time `json:"token_expire"`
}

func SessionJsonToEntity(sessionJson *SessionJson) *entity.Session {
	return &entity.Session{
		ID:          sessionJson.ID,
		SessionID:   sessionJson.SessionID,
		UserID:      sessionJson.UserID,
		Token:       sessionJson.Token,
		TokenExpire: sessionJson.TokenExpire,
	}
}

func SessionEntityToJson(session *entity.Session) *SessionJson {
	return &SessionJson{
		ID:          session.ID,
		SessionID:   session.SessionID,
		UserID:      session.UserID,
		Token:       session.Token,
		TokenExpire: session.TokenExpire,
	}
}