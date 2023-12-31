package usecase

import (
	"context"

	"github.com/Uwasaru/Sayka/domain/entity"
	"github.com/Uwasaru/Sayka/domain/repository"
)

var _ IFavoriteUsecase = &FavoriteUsecase{}

// FavoriteUsecaseは投稿に関するユースケースを担当します
type FavoriteUsecase struct {
	fr repository.FavoriteRepository
}

// IFavoriteUsecaseは投稿に関するユースケースを担当します
type IFavoriteUsecase interface {
	GetByID(ctx context.Context, id string) (*entity.Favorite, error)
	GetByUserID(ctx context.Context, userID string) (*entity.Favorites, error)
	GetBySaykaID(ctx context.Context, saykaID string) (entity.FavoriteUsers, error)
	GetCountBySaykaID(ctx context.Context, saykaID string) (int, error)
	GetAll(ctx context.Context) (*entity.Favorites, error)
	CreateFavorite(ctx context.Context, favorite *entity.Favorite) error
	DeleteFavorite(ctx context.Context, id string) error
	DeleteFavoriteBySaykaIDUserID(ctx context.Context, saykaID string, userID string) error
}

// NewFavoriteUsecaseは新しいFavoriteUsecaseを初期化し構造体のポインタを返します
func NewFavoriteUsecase(fr repository.FavoriteRepository) IFavoriteUsecase {
	return &FavoriteUsecase{
		fr: fr,
	}
}

// GetByIDはIDを指定して投稿を取得します
func (fu *FavoriteUsecase) GetByID(ctx context.Context, id string) (*entity.Favorite, error) {
	return fu.fr.GetByID(ctx, id)
}

// GetByUserIDはUserIDを指定して投稿を取得します
func (fu *FavoriteUsecase) GetByUserID(ctx context.Context, userID string) (*entity.Favorites, error) {
	return fu.fr.GetByUserID(ctx, userID)
}

// GetBySaykaIDはSaykaIDを指定して投稿を取得します
func (fu *FavoriteUsecase) GetBySaykaID(ctx context.Context, saykaID string) (entity.FavoriteUsers, error) {
	return fu.fr.GetBySaykaID(ctx, saykaID)
}

func (fu *FavoriteUsecase) GetCountBySaykaID(ctx context.Context, saykaID string) (int, error) {
	return fu.fr.GetCountBySaykaID(ctx, saykaID)
}

// GetAllは全ての投稿を取得します
func (fu *FavoriteUsecase) GetAll(ctx context.Context) (*entity.Favorites, error) {
	return fu.fr.GetAll(ctx)
}

// Createは投稿を作成します
func (fu *FavoriteUsecase) CreateFavorite(ctx context.Context, favorite *entity.Favorite) error {
	return fu.fr.CreateFavorite(ctx, favorite)
}

// DeleteFavoriteは投稿を削除します
func (fu *FavoriteUsecase) DeleteFavorite(ctx context.Context, id string) error {
	return fu.fr.DeleteFavorite(ctx, id)
}

// DeleteFavoriteBySaykaIDUserIDは投稿を削除します
func (fu *FavoriteUsecase) DeleteFavoriteBySaykaIDUserID(ctx context.Context, saykaID string, userID string) error {
	return fu.fr.DeleteFavoriteBySaykaIDUserID(ctx, saykaID, userID)
}
