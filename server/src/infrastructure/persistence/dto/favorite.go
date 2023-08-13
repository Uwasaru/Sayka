package dto

import (
	"github.com/Uwasaru/Sayka/domain/entity"
)

type FavoriteDto struct {
	ID     string `db:"id"`
	UserID string `db:"user_id"`
	PostID string `db:"post_id"`
}

type FavoritesDto []*FavoriteDto

func FavoriteDtoToEntity(dto *FavoriteDto) *entity.Favorite {
	return &entity.Favorite{
		ID:     dto.ID,
		UserID: dto.UserID,
		PostID: dto.PostID,
	}
}

func FavoriteEntityToDto(f *entity.Favorite) FavoriteDto {
	return FavoriteDto{
		ID:     f.ID,
		UserID: f.UserID,
		PostID: f.PostID,
	}
}

func FavoritesDtoToEntity(dtos *FavoritesDto) *entity.Favorites {
	favorites := &entity.Favorites{}
	for _, dto := range *dtos {
		favorite := FavoriteDtoToEntity(dto)
		*favorites = append(*favorites, favorite)
	}
	return favorites
}
