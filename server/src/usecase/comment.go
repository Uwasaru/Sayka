package usecase

import (
	"context"
	"github.com/Uwasaru/Sayka/domain/entity"
	"github.com/Uwasaru/Sayka/domain/repository"
)

var _ ICommentUsecase = &CommentUsecase{}

// CommentUsecaseはコメントに関するユースケースを担当します
type CommentUsecase struct {
	cr repository.CommentRepository
}

// ICommentUsecaseはコメントに関するユースケースを担当します
type ICommentUsecase interface {
	GetByID(ctx context.Context, id string) (*entity.Comment, error)
	GetByUserID(ctx context.Context, userID string) (*entity.Comments, error)
	GetByPostID(ctx context.Context, postID string) (*entity.Comments, error)
	GetAll(ctx context.Context, ) (*entity.Comments, error)
	CreateComment(ctx context.Context, comment *entity.Comment) error
	DeleteComment(ctx context.Context, id string) error
}

// NewCommentUsecaseは新しいCommentUsecaseを初期化し構造体のポインタを返します
func NewCommentUsecase(cr repository.CommentRepository) ICommentUsecase {
	return &CommentUsecase{
		cr: cr,
	}
}

// GetByIDはIDを指定してコメントを取得します
func (cu *CommentUsecase) GetByID(ctx context.Context, id string) (*entity.Comment, error) {
	return cu.cr.GetByID(ctx, id)
}

// GetByUserIDはUserIDを指定してコメントを取得します
func (cu *CommentUsecase) GetByUserID(ctx context.Context,userID string) (*entity.Comments, error) {
	return cu.cr.GetByUserID(ctx, userID)
}

// GetByPostIDはPostIDを指定してコメントを取得します
func (cu *CommentUsecase) GetByPostID(ctx context.Context,postID string) (*entity.Comments, error) {
	return cu.cr.GetByPostID(ctx, postID)
}

// GetAllは全てのコメントを取得します
func (cu *CommentUsecase) GetAll(ctx context.Context) ( *entity.Comments, error) {
	return cu.cr.GetAll(ctx)
}

// Createはコメントを作成します
func (cu *CommentUsecase) CreateComment(ctx context.Context, comment *entity.Comment) error {
	return cu.cr.CreateComment(ctx, comment)
}

// DeleteCommentはコメントを削除します
func (cu *CommentUsecase) DeleteComment(ctx context.Context, id string) error {
	return cu.cr.DeleteComment(ctx, id)
}
