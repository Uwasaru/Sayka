package json

import (
	"time"

	"github.com/Uwasaru/Sayka/domain/entity"
)

type SaykaJson struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	User        *UserJson `json:"user"`
	Title       string    `json:"title"`
	GithubUrl   string    `json:"github_url"`
	AppUrl      string    `json:"app_url"`
	SlideUrl    string    `json:"slide_url"`
	ArticleUrl  string    `json:"article_url"`
	FigmaUrl    string    `json:"figma_url"`
	Description string    `json:"description"`
	Tags        []string  `json:"tags"`
	Favorites   int       `json:"favorite_count"`
	Comments    int       `json:"comment_count"`
	IsFavorite  bool      `json:"is_favorite"`
	IsMe        bool      `json:"is_me"`
	CreatedAt   time.Time `json:"created_at"`
}

type MeJson struct {
	User         *UserJson `json:"user"`
	SaykaCount   int       `json:"sayka_count"`
	FavoritedCount int       `json:"favorited_count"`
	IsMe 			bool      `json:"is_me"`
}

type SaykaRequestJson struct {
	LastSaykaID string `json:"last_sayka_id"`
	SaykaAmount int    `json:"sayka_amount"`
}

type SaykasJson []SaykaJson

func SaykaJsonToEntity(saykaJson *SaykaJson) *entity.Sayka {
	return &entity.Sayka{
		ID:          saykaJson.ID,
		UserID:      saykaJson.UserID,
		Title:       saykaJson.Title,
		GithubUrl:   saykaJson.GithubUrl,
		AppUrl:      saykaJson.AppUrl,
		SlideUrl:    saykaJson.SlideUrl,
		ArticleUrl:  saykaJson.ArticleUrl,
		FigmaUrl:    saykaJson.FigmaUrl,
		Description: saykaJson.Description,
		Tags:        saykaJson.Tags,
		Favorites:   saykaJson.Favorites,
		Comments:    saykaJson.Comments,
		IsFavorite:  saykaJson.IsFavorite,
		IsMe:        saykaJson.IsMe,
		CreatedAt:   saykaJson.CreatedAt,
	}
}

func SaykasEntityToJson(c entity.Saykas) *SaykasJson {
	var SaykasJson SaykasJson
	for _, glyph := range c {
		SaykasJson = append(SaykasJson, *SaykaEntityToJson(glyph))
	}

	return &SaykasJson
}

func SaykaEntityToJson(sayka *entity.Sayka) *SaykaJson {
	return &SaykaJson{
		ID:          sayka.ID,
		UserID:      sayka.UserID,
		User:        UserEntityToJson(sayka.User),
		Title:       sayka.Title,
		GithubUrl:   sayka.GithubUrl,
		AppUrl:      sayka.AppUrl,
		SlideUrl:    sayka.SlideUrl,
		ArticleUrl:  sayka.ArticleUrl,
		FigmaUrl:    sayka.FigmaUrl,
		Description: sayka.Description,
		Tags:        sayka.Tags,
		Favorites:   sayka.Favorites,
		Comments:    sayka.Comments,
		IsFavorite:  sayka.IsFavorite,
		IsMe:        sayka.IsMe,
		CreatedAt:   sayka.CreatedAt,
	}
}


func MeEntityToJson(me *entity.Me) *MeJson {
	return &MeJson{
		User: UserEntityToJson(me.User),
		SaykaCount: me.SaykaCount,
		FavoritedCount: me.FavoritedCount,
		IsMe: me.IsMe,
	}
}