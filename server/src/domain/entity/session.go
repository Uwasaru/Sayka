package entity

import (
	"time"
)

// Sessionはセッションを表します
type Session struct {
	ID        string    `json:"id"`
	SessionID string    `json:"session_id"`
	UserID    string    `json:"user_id"`
	Token     string    `json:"token"`
	TokenExpire time.Time `json:"token_expire"`
}

// Sessionsはセッションの配列を表します
type Sessions []*Session