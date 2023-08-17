package dto

import (
	"github.com/Uwasaru/Sayka/domain/entity"
	"time"
)

type PostDto struct {
	ID          string    `db:"id"`
	UserID      string    `db:"user_id"`
	Title       string    `db:"title"`
	GithubUrl   string    `db:"github_url"`
	AppUrl      string    `db:"app_url"`
	SlideUrl    string    `db:"slide_url"`
	ArticleUrl  string    `db:"article_url"`
	FigmaUrl    string    `db:"figma_url"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type PostsDto []*PostDto

func PostDtoToEntity(dto *PostDto) *entity.Post {
	return &entity.Post{
		ID:          dto.ID,
		UserID:      dto.UserID,
		Title:       dto.Title,
		GithubUrl:   dto.GithubUrl,
		AppUrl:      dto.AppUrl,
		SlideUrl:    dto.SlideUrl,
		ArticleUrl:  dto.ArticleUrl,
		FigmaUrl:    dto.FigmaUrl,
		Description: dto.Description,
		CreatedAt:   dto.CreatedAt,
	}
}

func PostEntityToDto(p *entity.Post) PostDto {
	return PostDto{
		ID:          p.ID,
		UserID:      p.UserID,
		Title:       p.Title,
		GithubUrl:   p.GithubUrl,
		AppUrl:      p.AppUrl,
		SlideUrl:    p.SlideUrl,
		ArticleUrl:  p.ArticleUrl,
		FigmaUrl:    p.FigmaUrl,
		Description: p.Description,
		CreatedAt:   p.CreatedAt,
	}
}

func PostsDtoToEntity(dtos *PostsDto) *entity.Posts {
	posts := &entity.Posts{}
	for _, dto := range *dtos {
		post := PostDtoToEntity(dto)
		*posts = append(*posts, post)
	}
	return posts
}
