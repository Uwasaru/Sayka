package entity

import (
	"time"
)

// Postは投稿を表します
type Post struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Title       string    `json:"title"`
	GithubUrl   string    `json:"github_url"`
	AppUrl      string    `json:"app_url"`
	SlideUrl    string    `json:"slide_url"`
	Description string    `json:"description"`
	Tags        []string  `json:"tags"`
	Favorites   int       `json:"favorites"`
	Comments    int       `json:"comments"`
	IsFavorite  bool      `json:"is_favorite"`
	CreatedAt   time.Time `json:"created_at"`
}

type Posts []*Post
