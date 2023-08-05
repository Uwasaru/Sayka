package entity

import (
	"time"
)

// Userはユーザー情報を表します
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	GithubUrl string    `json:"github_url"`
	CreatedAt time.Time `json:"created_at"`
}

type Users []*User
