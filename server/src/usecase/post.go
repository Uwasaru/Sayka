package usecase

import (
	"context"
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
	GetByID(ctx context.Context, id string) (*entity.Post, error)
	// GetByUserIDはUserIDを指定して投稿を取得します
	GetByUserID(ctx context.Context, userID string) (*entity.Posts, error)
	// GetAllは全ての投稿を取得します
	GetAll(ctx context.Context) (*entity.Posts, error)
	// Createは投稿を作成します
	CreatePost(ctx context.Context, post *entity.Post) error
	// UpdatePostは投稿を更新します
	UpdatePost(ctx context.Context, post *entity.Post) error
	// DeletePostは投稿を削除します
	DeletePost(ctx context.Context, id string) error
}

// NewPostUsecaseは新しいPostUsecaseを初期化し構造体のポインタを返します
func NewPostUsecase(pr repository.PostRepository) IPostUsecase {
	return &PostUsecase{
		pr: pr,
	}
}

// GetByIDはIDを指定して投稿を取得します
func (pu *PostUsecase) GetByID(ctx context.Context, id string) (*entity.Post, error) {
	return pu.pr.GetByID(ctx, id)
}

// GetByUserIDはUserIDを指定して投稿を取得します
func (pu *PostUsecase) GetByUserID(ctx context.Context, userID string) (*entity.Posts, error) {
	return pu.pr.GetByUserID(ctx, userID)
}

// GetAllは全ての投稿を取得します
func (pu *PostUsecase) GetAll(ctx context.Context) (*entity.Posts, error) {
	return pu.pr.GetAll(ctx)
}

// Createは投稿を作成します
func (pu *PostUsecase) CreatePost(ctx context.Context, post *entity.Post) error {
	return pu.pr.CreatePost(ctx, post)
}

// UpdatePostは投稿を更新します
func (pu *PostUsecase) UpdatePost(ctx context.Context, post *entity.Post) error {
	return pu.pr.UpdatePost(ctx, post)
}

// DeletePostは投稿を削除します
func (pu *PostUsecase) DeletePost(ctx context.Context, id string) error {
	return pu.pr.DeletePost(ctx, id)
}
