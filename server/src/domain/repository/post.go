package repository

import (
	"context"

	"github.com/Uwasaru/Sayka/domain/entity"
)

// PostRepositoryは投稿に関する永続化を担当します
type PostRepository interface {
	GetByID(ctx context.Context, id string) (*entity.Post, error)
	GetByUserID(ctx context.Context, userID string) (*entity.Posts, error)
	GetAll(ctx context.Context) (*entity.Posts, error)
	GetTimeLine(ctx context.Context, id string) (*entity.Posts, error)
	CreatePost(ctx context.Context, post *entity.Post) error
	UpdatePost(ctx context.Context, post *entity.Post) error
	DeletePost(ctx context.Context, id string) error
}
