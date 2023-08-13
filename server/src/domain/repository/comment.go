package repository

import (
	"context"
	"github.com/Uwasaru/Sayka/domain/entity"
)

// CommentRepositoryはコメントに関する永続化を担当します
type CommentRepository interface {
	GetByID(ctx context.Context, id string) (*entity.Comment, error)
	GetByUserID(ctx context.Context, userID string) (*entity.Comments, error)
	GetByPostID(ctx context.Context, postID string) (*entity.Comments, error)
	GetAll(ctx context.Context) (*entity.Comments, error)
	CreateComment(ctx context.Context, comment *entity.Comment) error
	DeleteComment(ctx context.Context, id string) error
}
