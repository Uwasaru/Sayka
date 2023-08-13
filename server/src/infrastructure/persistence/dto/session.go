package dto

import (
	"github.com/Uwasaru/Sayka/domain/entity"
	"time"
)

type SessionDto struct {
	ID          string    `db:"id"`
	SessionID   string    `db:"session_id"`
	UserID      string    `db:"user_id"`
	Token       string    `db:"token"`
	TokenExpire time.Time `db:"token_expire"`
}

type SessionsDto []*SessionDto

func SessionDtoToEntity(dto *SessionDto) *entity.Session {
	return &entity.Session{
		ID:          dto.ID,
		SessionID:   dto.SessionID,
		UserID:      dto.UserID,
		Token:       dto.Token,
		TokenExpire: dto.TokenExpire,
	}
}

func SessionEntityToDto(s *entity.Session) SessionDto {
	return SessionDto{
		ID:          s.ID,
		SessionID:   s.SessionID,
		UserID:      s.UserID,
		Token:       s.Token,
		TokenExpire: s.TokenExpire,
	}
}

func SessionsDtoToEntity(dtos *SessionsDto) *entity.Sessions {
	sessions := &entity.Sessions{}
	for _, dto := range *dtos {
		session := SessionDtoToEntity(dto)
		*sessions = append(*sessions, session)
	}
	return sessions
}
