package entity

// Favoriteはユーザー情報を表します
type Favorite struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	SaykaID string `json:"sayka_id"`
}

type FavoriteUsers []string

type Favorites []*Favorite
