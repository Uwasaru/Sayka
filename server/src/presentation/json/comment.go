package json

import (
	"time"

	"github.com/Uwasaru/Sayka/domain/entity"
)

type CommentJson struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	SaykaID   string    `json:"sayka_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentsJson []CommentJson

func CommentJsonToEntity(commentJson *CommentJson) *entity.Comment {
	return &entity.Comment{
		ID:        commentJson.ID,
		UserID:    commentJson.UserID,
		SaykaID:   commentJson.SaykaID,
		Content:   commentJson.Content,
		CreatedAt: commentJson.CreatedAt,
	}
}

func CommentEntityToJson(comment *entity.Comment) *CommentJson {
	return &CommentJson{
		ID:        comment.ID,
		UserID:    comment.UserID,
		SaykaID:   comment.SaykaID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
	}
}

func CommentsEntityToJson(comments entity.Comments) *CommentsJson {
	var commentsJson CommentsJson
	for _, comment := range comments {
		commentsJson = append(commentsJson, *CommentEntityToJson(comment))
	}

	return &commentsJson
}
