package dto

import (
	"time"

	"github.com/Uwasaru/Sayka/domain/entity"
)

type CommentDto struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	SaykaID   string    `db:"sayka_id"`
	Content   string    `db:"content"`
	Type      int       `db:"type"`
	CreatedAt time.Time `db:"created_at"`
}

type CommentSaykaDto struct {
	SaykaID string `db:"sayka_id"`
}

type CommentsDto []*CommentDto

func CommentDtoToEntity(dto *CommentDto) *entity.Comment {
	return &entity.Comment{
		ID:        dto.ID,
		UserID:    dto.UserID,
		SaykaID:   dto.SaykaID,
		Content:   dto.Content,
		Type:      dto.Type,
		CreatedAt: dto.CreatedAt,
	}
}

func CommentSaykaDtoToEntity(dto *CommentSaykaDto) string {
	return dto.SaykaID
}

func CommentEntityToDto(c *entity.Comment) CommentDto {
	return CommentDto{
		ID:        c.ID,
		UserID:    c.UserID,
		SaykaID:   c.SaykaID,
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
