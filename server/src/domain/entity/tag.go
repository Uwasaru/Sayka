package entity

// Tagはタグを表します

type Tag struct {
	ID      int    `json:"id"`
	SaykaID string `json:"sayka_id"`
	Name    string `json:"name"`
}

type Tags []*Tag
