package repository

import (
	"github.com/Uwasaru/Sayka/domain/entity"
)

// CommentRepositoryはコメントに関する永続化を担当します
type CommentRepository interface {
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
