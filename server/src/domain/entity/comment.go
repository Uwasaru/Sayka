package entity

import "time"

// Commentはコメント情報を表します
type Comment struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	PostID    string    `json:"post_id"`
	Content   string    `json:"content"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type Comments []*Comment
