package dto

import (
	"github.com/Uwasaru/Sayka/domain/entity"
	"time"
)

type CommentDto struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	PostID    string    `db:"post_id"`
	Content   string    `db:"content"`
	Type      int       `db:"type"`
	CreatedAt time.Time `db:"created_at"`
}

type CommentPostDto struct {
	PostID string `db:"post_id"`
}

type CommentsDto []*CommentDto

func CommentDtoToEntity(dto *CommentDto) *entity.Comment {
	return &entity.Comment{
		ID:        dto.ID,
		UserID:    dto.UserID,
		PostID:    dto.PostID,
		Content:   dto.Content,
		Type:      dto.Type,
		CreatedAt: dto.CreatedAt,
	}
}

func CommentPostDtoToEntity(dto *CommentPostDto) string {
	return dto.PostID
}

func CommentEntityToDto(c *entity.Comment) CommentDto {
	return CommentDto{
		ID:        c.ID,
		UserID:    c.UserID,
		PostID:    c.PostID,
		Content:   c.Content,
		Type:      c.Type,
		CreatedAt: c.CreatedAt,
	}
}

func CommentsDtoToEntity(dtos *CommentsDto) *entity.Comments {
	comments := &entity.Comments{}
	for _, dto := range *dtos {
		comment := CommentDtoToEntity(dto)
		*comments = append(*comments, comment)
	}
	return comments
}
