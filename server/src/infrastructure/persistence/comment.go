package persistence

import (
	"context"

	"github.com/Uwasaru/Sayka/domain/entity"
	"github.com/Uwasaru/Sayka/domain/repository"
	"github.com/Uwasaru/Sayka/infrastructure/database"
	d "github.com/Uwasaru/Sayka/infrastructure/persistence/dto"
)

var _ repository.CommentRepository = &CommentRepository{}

type CommentRepository struct {
	conn *database.Conn
}

func NewCommentRepository(conn *database.Conn) *CommentRepository {
	return &CommentRepository{
		conn: conn,
	}
}

func (cr *CommentRepository) GetByID(ctx context.Context, id string) (*entity.Comment, error) {
	query := `
	SELECT *
	FROM comments
	WHERE id = ?
	`
	var dto d.CommentDto
	err := cr.conn.DB.GetContext(ctx, &dto, query, id)
	if err != nil {
		return nil, err
	}
	return d.CommentDtoToEntity(&dto), nil
}

func (cr *CommentRepository) GetByUserID(ctx context.Context, userID string) (*entity.Comments, error) {
	query := `
	SELECT *
	FROM comments
	WHERE user_id = ?
	`
	rows, err := cr.conn.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	var comments = entity.Comments{}
	for rows.Next() {
		var dto d.CommentDto
		err := rows.Scan(&dto)
		if err != nil {
			return nil, err
		}
		comments = append(comments, d.CommentDtoToEntity(&dto))
	}
	return &comments, nil
}

func (cr *CommentRepository) GetBySaykaID(ctx context.Context, saykaID string) (*entity.Comments, error) {
	query := `
	SELECT *
	FROM comments
	WHERE sayka_id = ?
	`
	rows, err := cr.conn.DB.QueryContext(ctx, query, saykaID)
	if err != nil {
		return nil, err
	}

	var comments = entity.Comments{}
	for rows.Next() {
		var dto d.CommentDto
		err := rows.Scan(&dto.ID, &dto.UserID, &dto.SaykaID, &dto.Content, &dto.Type, &dto.CreatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, d.CommentDtoToEntity(&dto))
	}
	return &comments, nil
}

func (cr *CommentRepository) GetAll(ctx context.Context) (*entity.Comments, error) {
	query := `
	SELECT *
	FROM comments
	`
	rows, err := cr.conn.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var comments = entity.Comments{}
	for rows.Next() {
		var dto d.CommentDto
		err := rows.Scan(&dto.ID, &dto.UserID, &dto.SaykaID, &dto.Content, &dto.Type, &dto.CreatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, d.CommentDtoToEntity(&dto))
	}
	return &comments, nil
}

func (cr *CommentRepository) CreateComment(ctx context.Context, comment *entity.Comment) error {
	query := `
	INSERT INTO comments (id, user_id, sayka_id, content, type, created_at)
	VALUES (:id, :user_id, :sayka_id, :content, :type, :created_at)
	`
	dto := d.CommentEntityToDto(comment)

	_, err := cr.conn.DB.NamedExecContext(ctx, query, &dto)
	if err != nil {
		return err
	}
	return nil
}

func (cr *CommentRepository) DeleteComment(ctx context.Context, id string) error {
	query := `
	DELETE FROM comments
	WHERE id = :id
	`

	_, err := cr.conn.DB.NamedExecContext(ctx, query, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	return nil
}
