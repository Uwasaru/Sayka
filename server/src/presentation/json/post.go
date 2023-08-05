package json

import (
	"github.com/Uwasaru/Sayka/domain/entity"
	"time"
)

type PostJson struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Title       string    `json:"title"`
	GithubUrl   string    `json:"github_url"`
	AppUrl      string    `json:"app_url"`
	SlideUrl    string    `json:"slide_url"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func PostJsonToEntity(postJson *PostJson) *entity.Post {
	return &entity.Post{
		ID:          postJson.ID,
		UserID:      postJson.UserID,
		Title:       postJson.Title,
		GithubUrl:   postJson.GithubUrl,
		AppUrl:      postJson.AppUrl,
		SlideUrl:    postJson.SlideUrl,
		Description: postJson.Description,
		CreatedAt:   postJson.CreatedAt,
	}
}

func PostEntityToJson(post *entity.Post) *PostJson {
	return &PostJson{
		ID:          post.ID,
		UserID:      post.UserID,
		Title:       post.Title,
		GithubUrl:   post.GithubUrl,
		AppUrl:      post.AppUrl,
		SlideUrl:    post.SlideUrl,
		Description: post.Description,
		CreatedAt:   post.CreatedAt,
	}
}
