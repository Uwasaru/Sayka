package repository

import (
	"github.com/Uwasaru/Sayka/domain/entity"
)

// PostRepositoryは投稿に関する永続化を担当します
type PostRepository interface {
	// GetByIDはIDを指定して投稿を取得します
	GetByID(id string) (*entity.Post, error)
	// GetByUserIDはUserIDを指定して投稿を取得します
	GetByUserID(userID string) (*entity.Posts, error)
	// GetAllは全ての投稿を取得します
	GetAll() (*entity.Posts, error)
	// Createは投稿を作成します
	CreatePost(post *entity.Post) error
	// UpdatePostは投稿を更新します
	UpdatePost(post *entity.Post) error
	// DeletePostは投稿を削除します
	DeletePost(id string) error
}
