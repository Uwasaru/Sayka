package json

import (
	"github.com/Uwasaru/Sayka/domain/entity"
)

type FavoriteJson struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	SaykaID string `json:"sayka_id"`
}

func FavoriteJsonToEntity(favoriteJson *FavoriteJson) *entity.Favorite {
	return &entity.Favorite{
		ID:      favoriteJson.ID,
		UserID:  favoriteJson.UserID,
		SaykaID: favoriteJson.SaykaID,
	}
}

func FavoriteEntityToJson(favorite *entity.Favorite) *FavoriteJson {
	return &FavoriteJson{
		ID:      favorite.ID,
		UserID:  favorite.UserID,
		SaykaID: favorite.SaykaID,
	}
}
