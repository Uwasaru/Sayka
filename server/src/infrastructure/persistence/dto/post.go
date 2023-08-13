package dto

import (
	"time"
	"github.com/Uwasaru/Sayka/domain/entity"
)

type PostDto struct {
	ID     string `db:"id"`
	UserID string `db:"user_id"`
	Title string `db:"title"`
	GithubUrl string `db:"github_url"`
	AppUrl string `db:"app_url"`
	SlideUrl string `db:"slide_url"`
	Description string `db:"description"`
	CreatedAt time.Time `db:"created_at"`
}

type PostsDto []*PostDto

func PostDtoToEntity(dto *PostDto) *entity.Post {
	return &entity.Post{
		ID: dto.ID,
		UserID: dto.UserID,
		Title: dto.Title,
		GithubUrl: dto.GithubUrl,
		AppUrl: dto.AppUrl,
		SlideUrl: dto.SlideUrl,
		Description: dto.Description,
		CreatedAt: dto.CreatedAt,
	}
}

func PostEntityToDto(p *entity.Post) PostDto {
	return PostDto{
		ID: p.ID,
		UserID: p.UserID,
		Title: p.Title,
		GithubUrl: p.GithubUrl,
		AppUrl: p.AppUrl,
		SlideUrl: p.SlideUrl,
		Description: p.Description,
		CreatedAt: p.CreatedAt,
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
