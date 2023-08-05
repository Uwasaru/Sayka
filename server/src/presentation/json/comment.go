package json

import (
	"github.com/Uwasaru/Sayka/domain/entity"
	"time"
)

type CommentJson struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	PostID    string    `json:"post_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func CommentJsonToEntity(commentJson *CommentJson) *entity.Comment {
	return &entity.Comment{
		ID:        commentJson.ID,
		UserID:    commentJson.UserID,
		PostID:    commentJson.PostID,
		Content:   commentJson.Content,
		CreatedAt: commentJson.CreatedAt,
	}
}

func CommentEntityToJson(comment *entity.Comment) *CommentJson {
	return &CommentJson{
		ID:        comment.ID,
		UserID:    comment.UserID,
		PostID:    comment.PostID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
	}
}
