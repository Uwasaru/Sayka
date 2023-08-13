package repository

import (
	"context"
	"github.com/Uwasaru/Sayka/domain/entity"
)

// TagRepositoryはタグに関する永続化を担当します

type TagRepository interface {
	GetByID(ctx context.Context, id string) (*entity.Tag, error)
	GetByName(ctx context.Context, name string) (*entity.Tag, error)
	GetAll(ctx context.Context) (*entity.Tags, error)
	CreateTag(ctx context.Context, tag *entity.Tag) error
	DeleteTag(ctx context.Context, id string) error
}
