package usecase

import (
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
	// GetByIDはIDを指定して投稿を取得します
	GetByID(id string) (*entity.Tag, error)
	// GetByNameはNameを指定して投稿を取得します
	GetByName(name string) (*entity.Tag, error)
	// GetAllは全ての投稿を取得します
	GetAll() (*entity.Tags, error)
	// Createは投稿を作成します
	CreateTag(tag *entity.Tag) error
	// Deleteは投稿を削除します
	DeleteTag(id string) error
}

// NewTagUsecaseは新しいTagUsecaseを初期化し構造体のポインタを返します
func NewTagUsecase(tr repository.TagRepository) ITagUsecase {
	return &TagUsecase{
		tr: tr,
	}
}

// GetByIDはIDを指定して投稿を取得します
func (tu *TagUsecase) GetByID(id string) (*entity.Tag, error) {
	return tu.tr.GetByID(id)
}

// GetByNameはNameを指定して投稿を取得します
func (tu *TagUsecase) GetByName(name string) (*entity.Tag, error) {
	return tu.tr.GetByName(name)
}

// GetAllは全ての投稿を取得します
func (tu *TagUsecase) GetAll() (*entity.Tags, error) {
	return tu.tr.GetAll()
}

// Createは投稿を作成します
func (tu *TagUsecase) CreateTag(tag *entity.Tag) error {
	return tu.tr.CreateTag(tag)
}

// Deleteは投稿を削除します
func (tu *TagUsecase) DeleteTag(id string) error {
	return tu.tr.DeleteTag(id)
}
