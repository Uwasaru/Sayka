package dto

import (
	"time"

	"github.com/Uwasaru/Sayka/domain/entity"
)

type SaykaDto struct {
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

type SaykasDto []*SaykaDto

func SaykaDtoToEntity(dto *SaykaDto) *entity.Sayka {
	return &entity.Sayka{
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

func SaykaEntityToDto(p *entity.Sayka) SaykaDto {
	return SaykaDto{
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

func SaykasDtoToEntity(dtos *SaykasDto) *entity.Saykas {
	saykas := &entity.Saykas{}
	for _, dto := range *dtos {
		sayka := SaykaDtoToEntity(dto)
		*saykas = append(*saykas, sayka)
	}
	return saykas
}
