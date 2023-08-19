package entity

import (
	"time"
)

// Saykaは投稿を表します
type Sayka struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	User        *User     `json:"user"`
	Title       string    `json:"title"`
	GithubUrl   string    `json:"github_url"`
	AppUrl      string    `json:"app_url"`
	SlideUrl    string    `json:"slide_url"`
	ArticleUrl  string    `json:"article_url"`
	FigmaUrl    string    `json:"figma_url"`
	Description string    `json:"description"`
	Tags        []string  `json:"tags"`
	Favorites   int       `json:"favorites"`
	Comments    int       `json:"comments"`
	IsFavorite  bool      `json:"is_favorite"`
	IsMe        bool      `json:"is_me"`
	CreatedAt   time.Time `json:"created_at"`
}

type Saykas []*Sayka

type Me struct {
	User 	 *User     `json:"user"`
	SaykaCount int       `json:"sayka_count"`
	FavoritedCount int       `json:"favorited_count"`
}
