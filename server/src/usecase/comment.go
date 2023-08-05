package usecase

import (
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
	// GetByIDはIDを指定してコメントを取得します
	GetByID(id string) (*entity.Comment, error)
	// GetByUserIDはUserIDを指定してコメントを取得します
	GetByUserID(userID string) (*entity.Comments, error)
	// GetByPostIDはPostIDを指定してコメントを取得します
	GetByPostID(postID string) (*entity.Comments, error)
	// GetAllは全てのコメントを取得します
	GetAll() (*entity.Comments, error)
	// Createはコメントを作成します
	CreateComment(comment *entity.Comment) error
	// DeleteCommentはコメントを削除します
	DeleteComment(id string) error
}

// NewCommentUsecaseは新しいCommentUsecaseを初期化し構造体のポインタを返します
func NewCommentUsecase(cr repository.CommentRepository) ICommentUsecase {
	return &CommentUsecase{
		cr: cr,
	}
}

// GetByIDはIDを指定してコメントを取得します
func (cu *CommentUsecase) GetByID(id string) (*entity.Comment, error) {
	return cu.cr.GetByID(id)
}

// GetByUserIDはUserIDを指定してコメントを取得します
func (cu *CommentUsecase) GetByUserID(userID string) (*entity.Comments, error) {
	return cu.cr.GetByUserID(userID)
}

// GetByPostIDはPostIDを指定してコメントを取得します
func (cu *CommentUsecase) GetByPostID(postID string) (*entity.Comments, error) {
	return cu.cr.GetByPostID(postID)
}

// GetAllは全てのコメントを取得します
func (cu *CommentUsecase) GetAll() (*entity.Comments, error) {
	return cu.cr.GetAll()
}

// Createはコメントを作成します
func (cu *CommentUsecase) CreateComment(comment *entity.Comment) error {
	return cu.cr.CreateComment(comment)
}

// DeleteCommentはコメントを削除します
func (cu *CommentUsecase) DeleteComment(id string) error {
	return cu.cr.DeleteComment(id)
}
