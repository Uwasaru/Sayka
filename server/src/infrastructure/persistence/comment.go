package persistence

import (
	"database/sql"
	"github.com/Uwasaru/Sayka/domain/entity"
	"github.com/Uwasaru/Sayka/domain/repository"
)

var _ repository.CommentRepository = &CommentRepository{}

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (cr *CommentRepository) GetByID(id string) (*entity.Comment, error) {
	comment := &entity.Comment{}
	err := cr.db.QueryRow("SELECT * FROM comments WHERE id = ?", id).Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Content, &comment.Type, &comment.CreatedAt)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (cr *CommentRepository) GetByUserID(userID string) (*entity.Comments, error) {
	rows, err := cr.db.Query("SELECT * FROM comments WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := &entity.Comments{}
	for rows.Next() {
		comment := &entity.Comment{}
		err := rows.Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Content, &comment.Type, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}
		*comments = append(*comments, comment)
	}
	return comments, nil
}

func (cr *CommentRepository) GetByPostID(postID string) (*entity.Comments, error) {
	rows, err := cr.db.Query("SELECT * FROM comments WHERE post_id = ?", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := &entity.Comments{}
	for rows.Next() {
		comment := &entity.Comment{}
		err := rows.Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Content, &comment.Type, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}
		*comments = append(*comments, comment)
	}
	return comments, nil
}

func (cr *CommentRepository) GetAll() (*entity.Comments, error) {
	rows, err := cr.db.Query("SELECT * FROM comments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := &entity.Comments{}
	for rows.Next() {
		comment := &entity.Comment{}
		err := rows.Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Content, &comment.Type, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}
		*comments = append(*comments, comment)
	}
	return comments, nil
}

func (cr *CommentRepository) CreateComment(comment *entity.Comment) error {
	_, err := cr.db.Exec("INSERT INTO comments (id, user_id, post_id, content, type, created_at) VALUES (?, ?, ?, ?, ?, ?)", comment.ID, comment.UserID, comment.PostID, comment.Content, comment.Type, comment.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (cr *CommentRepository) DeleteComment(id string) error {
	_, err := cr.db.Exec("DELETE FROM comments WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
