package repository

import (
	"context"
	"github.com/Uwasaru/Sayka/domain/entity"
)

type TagPostRepository interface {
	GetByID(ctx context.Context, id string) (*entity.TagPost, error)
	GetByPostID(ctx context.Context, postID string) (*entity.TagPosts, error)
	GetAll(ctx context.Context) (*entity.TagPosts, error)
	CreateTagPost(ctx context.Context, tagPost *entity.TagPost) error
	DeleteTagPost(ctx context.Context, id string) error
}
