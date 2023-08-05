package usecase

import (
	"github.com/Uwasaru/Sayka/domain/entity"
	"github.com/Uwasaru/Sayka/domain/repository"
)

var _ IPostUsecase = &PostUsecase{}

// PostUsecaseは投稿に関するユースケースを担当します
type PostUsecase struct {
	pr repository.PostRepository
}

// IPostUsecaseは投稿に関するユースケースを担当します
type IPostUsecase interface {
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

// NewPostUsecaseは新しいPostUsecaseを初期化し構造体のポインタを返します
func NewPostUsecase(pr repository.PostRepository) IPostUsecase {
	return &PostUsecase{
		pr: pr,
	}
}

// GetByIDはIDを指定して投稿を取得します
func (pu *PostUsecase) GetByID(id string) (*entity.Post, error) {
	return pu.pr.GetByID(id)
}

// GetByUserIDはUserIDを指定して投稿を取得します
func (pu *PostUsecase) GetByUserID(userID string) (*entity.Posts, error) {
	return pu.pr.GetByUserID(userID)
}

// GetAllは全ての投稿を取得します
func (pu *PostUsecase) GetAll() (*entity.Posts, error) {
	return pu.pr.GetAll()
}

// Createは投稿を作成します
func (pu *PostUsecase) CreatePost(post *entity.Post) error {
	return pu.pr.CreatePost(post)
}

// UpdatePostは投稿を更新します
func (pu *PostUsecase) UpdatePost(post *entity.Post) error {
	return pu.pr.UpdatePost(post)
}

// DeletePostは投稿を削除します
func (pu *PostUsecase) DeletePost(id string) error {
	return pu.pr.DeletePost(id)
}
