package json

import (
	"time"

	"github.com/Uwasaru/Sayka/domain/entity"
)

type PostJson struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
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
	CreatedAt   time.Time `json:"created_at"`
}

type PostRequestJson struct {
	LastPostID string `json:"last_post_id"`
	PostAmount int    `json:"post_amount"`
}

func PostJsonToEntity(postJson *PostJson) *entity.Post {
	return &entity.Post{
		ID:          postJson.ID,
		UserID:      postJson.UserID,
		Title:       postJson.Title,
		GithubUrl:   postJson.GithubUrl,
		AppUrl:      postJson.AppUrl,
		SlideUrl:    postJson.SlideUrl,
		ArticleUrl:  postJson.ArticleUrl,
		FigmaUrl:    postJson.FigmaUrl,
		Description: postJson.Description,
		Tags:        postJson.Tags,
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
		ArticleUrl:  post.ArticleUrl,
		FigmaUrl:    post.FigmaUrl,
		Description: post.Description,
		Tags:        post.Tags,
		CreatedAt:   post.CreatedAt,
	}
}
