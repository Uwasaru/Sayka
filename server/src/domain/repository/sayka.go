package repository

import (
	"context"

	"github.com/Uwasaru/Sayka/domain/entity"
)

// SaykaRepositoryは投稿に関する永続化を担当します
type SaykaRepository interface {
	GetByID(ctx context.Context, id string) (*entity.Sayka, error)
	GetByUserID(ctx context.Context, userID string) (*entity.Saykas, error)
	GetAll(ctx context.Context) (*entity.Saykas, error)
	GetTimeLine(ctx context.Context, id string, tag string) (*entity.Saykas, error)
	CreateSayka(ctx context.Context, sayka *entity.Sayka) error
	UpdateSayka(ctx context.Context, sayka *entity.Sayka) error
	DeleteSayka(ctx context.Context, id string) error
}
