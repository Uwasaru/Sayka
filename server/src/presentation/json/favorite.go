package json

import (
	"github.com/Uwasaru/Sayka/domain/entity"
)

type FavoriteJson struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	PostID string `json:"post_id"`
}

func FavoriteJsonToEntity(favoriteJson *FavoriteJson) *entity.Favorite {
	return &entity.Favorite{
		ID:     favoriteJson.ID,
		UserID: favoriteJson.UserID,
		PostID: favoriteJson.PostID,
	}
}

func FavoriteEntityToJson(favorite *entity.Favorite) *FavoriteJson {
	return &FavoriteJson{
		ID:     favorite.ID,
		UserID: favorite.UserID,
		PostID: favorite.PostID,
	}
}
