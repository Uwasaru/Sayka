package entity

// Tagはタグを表します

type Tag struct {
	ID   int `json:"id"`
	PostID string `json:"post_id"`
	Name string `json:"name"`
}

type Tags []*Tag
