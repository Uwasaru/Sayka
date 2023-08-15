package persistence

import (
	"context"
	// "fmt"

	"github.com/Uwasaru/Sayka/domain/entity"
	"github.com/Uwasaru/Sayka/domain/repository"
	"github.com/Uwasaru/Sayka/infrastructure/database"
	d "github.com/Uwasaru/Sayka/infrastructure/persistence/dto"
)

var _ repository.TagRepository = &TagRepository{}

// TagRepositoryはTagRepositoryの実装です
type TagRepository struct {
	conn *database.Conn
}

// NewTagRepositoryは新しいTagRepositoryを初期化し構造体のポインタを返します
func NewTagRepository(conn *database.Conn) *TagRepository {
	return &TagRepository{
		conn: conn,
	}
}

// GetByIDはIDを指定してタグを取得します
func (tr *TagRepository) GetByID(ctx context.Context, id string) (*entity.Tag, error) {
	query := `
	SELECT *
	FROM tags
	WHERE id = ?
	`
	var dto d.TagDto
	err := tr.conn.DB.Get(&dto, query, id)
	if err != nil {
		return nil, err
	}
	return d.TagDtoToEntity(&dto), nil
}

// GetByNameはNameを指定してタグを取得します
func (tr *TagRepository) GetByName(ctx context.Context, name string) (*entity.Tag, error) {
	query := `
	SELECT *
	FROM tags
	WHERE name = ?
	`
	var dto d.TagDto
	err := tr.conn.DB.Get(&dto, query, name)
	if err != nil {
		return nil, err
	}
	return d.TagDtoToEntity(&dto), nil
}

// GetAllは全てのタグを取得します
func (tr *TagRepository) GetAll(ctx context.Context) (*entity.Tags, error) {
	query := `
	SELECT *
	FROM tags
	`

	rows, err := tr.conn.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var tags entity.Tags
	for rows.Next() {
		var dto d.TagDto
		err := rows.Scan(&dto.ID, &dto.PostID, &dto.Name)
		if err != nil {
			return nil, err
		}
		tags = append(tags, d.TagDtoToEntity(&dto))
	}
	return &tags, nil
}

// GetByPostIDはPostIDを指定してタグを取得します
func (tr *TagRepository) GetByPostID(ctx context.Context, postID string) (*entity.Tags, error) {
	query := `
	SELECT *
	FROM tags
	WHERE post_id = ?
	`
	rows, err := tr.conn.DB.Query(query, postID)
	if err != nil {
		return nil, err
	}

	var tags entity.Tags
	for rows.Next() {
		var dto d.TagDto
		err := rows.Scan(&dto.ID, &dto.PostID, &dto.Name)
		if err != nil {
			return nil, err
		}
		tags = append(tags, d.TagDtoToEntity(&dto))
	}
	return &tags, nil
}

// Createはタグを作成します
func (tr *TagRepository) CreateTag(ctx context.Context, tag *entity.Tag) error {
	query := `
	INSERT INTO tags
	(id, name)
	VALUES (?, ?)
	`
	_, err := tr.conn.DB.Exec(query, tag.ID, tag.Name)
	if err != nil {
		return err
	}
	return nil
}

// Deleteはタグを削除します
func (tr *TagRepository) DeleteTag(ctx context.Context, id string) error {
	query := `
	DELETE FROM tags
	WHERE id = ?
	`
	_, err := tr.conn.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
