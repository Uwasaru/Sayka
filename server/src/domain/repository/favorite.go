package repository

import (
	"github.com/Uwasaru/Sayka/domain/entity"
)

// FavoriteRepositoryは投稿に関する永続化を担当します
type FavoriteRepository interface {
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
