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
	ur repository.UserRepository
}

// ICommentUsecaseはコメントに関するユースケースを担当します
type ICommentUsecase interface {
	GetByID(ctx context.Context, id string) (*entity.Comment, error)
	GetByUserID(ctx context.Context, userID string) (*entity.Comments, error)
	GetBySaykaID(ctx context.Context, saykaID string) (*entity.Comments, error)
	GetAll(ctx context.Context) (*entity.Comments, error)
	CreateComment(ctx context.Context, comment *entity.Comment) error
	DeleteComment(ctx context.Context, id string) error
}

// NewCommentUsecaseは新しいCommentUsecaseを初期化し構造体のポインタを返します
func NewCommentUsecase(cr repository.CommentRepository, ur repository.UserRepository) ICommentUsecase {
	return &CommentUsecase{
		cr: cr,
		ur: ur,
	}
}

// GetByIDはIDを指定してコメントを取得します
func (cu *CommentUsecase) GetByID(ctx context.Context, id string) (*entity.Comment, error) {
	return cu.cr.GetByID(ctx, id)
}

// GetByUserIDはUserIDを指定してコメントを取得します
func (cu *CommentUsecase) GetByUserID(ctx context.Context, userID string) (*entity.Comments, error) {
	comments, err := cu.cr.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	for _, comment := range *comments {
		user, err := cu.ur.GetUser(ctx, comment.UserID)
		if err != nil {
			return nil, err
		}
		comment.User = user
	}
	return comments, nil
}

// GetBySaykaIDはSaykaIDを指定してコメントを取得します
func (cu *CommentUsecase) GetBySaykaID(ctx context.Context, saykaID string) (*entity.Comments, error) {
	comments, err := cu.cr.GetBySaykaID(ctx, saykaID)
	if err != nil {
		return nil, err
	}
	for _, comment := range *comments {
		user, err := cu.ur.GetUser(ctx, comment.UserID)
		if err != nil {
			return nil, err
		}
		comment.User = user
	}
	return comments, nil
}

// GetAllは全てのコメントを取得します
func (cu *CommentUsecase) GetAll(ctx context.Context) (*entity.Comments, error) {
	comments, err := cu.cr.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, comment := range *comments {
		user, err := cu.ur.GetUser(ctx, comment.UserID)
		if err != nil {
			return nil, err
		}
		comment.User = user
	}
	return comments, nil
}

// Createはコメントを作成します
func (cu *CommentUsecase) CreateComment(ctx context.Context, comment *entity.Comment) error {
	err := cu.cr.CreateComment(ctx, comment)
	user, err := cu.ur.GetUser(ctx, comment.UserID)
	if err != nil {
		return err
	}
	comment.User = user
	return err
}

// DeleteCommentはコメントを削除します
func (cu *CommentUsecase) DeleteComment(ctx context.Context, id string) error {
	return cu.cr.DeleteComment(ctx, id)
}
