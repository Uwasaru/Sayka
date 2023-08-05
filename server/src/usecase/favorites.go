package usecase

import (
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
	// GetByIDはIDを指定して投稿を取得します
	GetByID(id string) (*entity.Favorite, error)
	// GetByUserIDはUserIDを指定して投稿を取得します
	GetByUserID(userID string) (*entity.Favorites, error)
	// GetByPostIDはPostIDを指定して投稿を取得します
	GetByPostID(postID string) (*entity.Favorites, error)
	// GetAllは全ての投稿を取得します
	GetAll() (*entity.Favorites, error)
	// Createは投稿を作成します
	CreateFavorite(favorite *entity.Favorite) error
	// DeleteFavoriteは投稿を削除します
	DeleteFavorite(id string) error
}

// NewFavoriteUsecaseは新しいFavoriteUsecaseを初期化し構造体のポインタを返します
func NewFavoriteUsecase(fr repository.FavoriteRepository) IFavoriteUsecase {
	return &FavoriteUsecase{
		fr: fr,
	}
}

// GetByIDはIDを指定して投稿を取得します
func (fu *FavoriteUsecase) GetByID(id string) (*entity.Favorite, error) {
	return fu.fr.GetByID(id)
}

// GetByUserIDはUserIDを指定して投稿を取得します
func (fu *FavoriteUsecase) GetByUserID(userID string) (*entity.Favorites, error) {
	return fu.fr.GetByUserID(userID)
}

// GetByPostIDはPostIDを指定して投稿を取得します
func (fu *FavoriteUsecase) GetByPostID(postID string) (*entity.Favorites, error) {
	return fu.fr.GetByPostID(postID)
}

// GetAllは全ての投稿を取得します
func (fu *FavoriteUsecase) GetAll() (*entity.Favorites, error) {
	return fu.fr.GetAll()
}

// Createは投稿を作成します
func (fu *FavoriteUsecase) CreateFavorite(favorite *entity.Favorite) error {
	return fu.fr.CreateFavorite(favorite)
}

// DeleteFavoriteは投稿を削除します
func (fu *FavoriteUsecase) DeleteFavorite(id string) error {
	return fu.fr.DeleteFavorite(id)
}
