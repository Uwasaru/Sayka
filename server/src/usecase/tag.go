package usecase

import (
	"context"

	"github.com/Uwasaru/Sayka/domain/entity"
	"github.com/Uwasaru/Sayka/domain/repository"
)

var _ ITagUsecase = &TagUsecase{}

// TagUsecaseは投稿に関するユースケースを担当します
type TagUsecase struct {
	tr repository.TagRepository
}

// ITagUsecaseは投稿に関するユースケースを担当します
type ITagUsecase interface {
	GetByID(ctx context.Context, id string) (*entity.Tag, error)
	GetByName(ctx context.Context, name string) (*entity.Tag, error)
	GetBySaykaID(ctx context.Context, saykaID string) (*entity.Tags, error)
	GetAll(ctx context.Context) (*entity.Tags, error)
	CreateTag(ctx context.Context, tag *entity.Tag) error
	DeleteTag(ctx context.Context, id string) error
}

// NewTagUsecaseは新しいTagUsecaseを初期化し構造体のポインタを返します
func NewTagUsecase(tr repository.TagRepository) ITagUsecase {
	return &TagUsecase{
		tr: tr,
	}
}

// GetByIDはIDを指定して投稿を取得します
func (tu *TagUsecase) GetByID(ctx context.Context, id string) (*entity.Tag, error) {
	return tu.tr.GetByID(ctx, id)
}

// GetByNameはNameを指定して投稿を取得します
func (tu *TagUsecase) GetByName(ctx context.Context, name string) (*entity.Tag, error) {
	return tu.tr.GetByName(ctx, name)
}

// GetBySaykaIDはSaykaIDを指定して投稿を取得します
func (tu *TagUsecase) GetBySaykaID(ctx context.Context, saykaID string) (*entity.Tags, error) {
	return tu.tr.GetBySaykaID(ctx, saykaID)
}

// GetAllは全ての投稿を取得します
func (tu *TagUsecase) GetAll(ctx context.Context) (*entity.Tags, error) {
	return tu.tr.GetAll(ctx)
}

// Createは投稿を作成します
func (tu *TagUsecase) CreateTag(ctx context.Context, tag *entity.Tag) error {
	return tu.tr.CreateTag(ctx, tag)
}

// Deleteは投稿を削除します
func (tu *TagUsecase) DeleteTag(ctx context.Context, id string) error {
	return tu.tr.DeleteTag(ctx, id)
}
