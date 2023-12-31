package repository

import (
	"context"

	"github.com/Uwasaru/Sayka/domain/entity"
)

// FavoriteRepositoryは投稿に関する永続化を担当します
type FavoriteRepository interface {
	GetByID(ctx context.Context, id string) (*entity.Favorite, error)
	GetByUserID(ctx context.Context, userID string) (*entity.Favorites, error)
	GetCountBySaykaID(ctx context.Context, saykaID string) (int, error)
	GetBySaykaID(ctx context.Context, saykaID string) (entity.FavoriteUsers, error)
	GetAllCountByUserID(ctx context.Context, userID string) (int, error)
	GetAll(ctx context.Context) (*entity.Favorites, error)
	CreateFavorite(ctx context.Context, favorite *entity.Favorite) error
	DeleteFavorite(ctx context.Context, id string) error
	DeleteFavoriteBySaykaIDUserID(ctx context.Context, saykaID string, userID string) error
}
