package entity


// Tagはタグを表します

type Tag struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
}

type Tags []*Tag