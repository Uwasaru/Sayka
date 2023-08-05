package entity

// Favoriteはユーザー情報を表します
type Favorite struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	PostID string `json:"post_id"`
}

type Favorites []*Favorite
