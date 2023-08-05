package entity

import (
	"time"
)

// Userはユーザー情報を表します
type User struct {
	ID        string    `json:"id"`
	GithubID	string    `json:"github_id"`
	GithubIcon	string    `json:"github_icon"`
	AccessToken	string    `json:"access_token"`
	RefleshToken	string    `json:"reflesh_token"`
	TokenExpire time.Time `json:"token_expire"`
	CreatedAt time.Time `json:"created_at"`
}

type Users []*User
