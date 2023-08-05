package persistence

import (
	"database/sql"
	"github.com/Uwasaru/Sayka/domain/entity"
	"github.com/Uwasaru/Sayka/domain/repository"
)

var _ repository.TagRepository = &TagRepository{}

// TagRepositoryはTagRepositoryの実装です
type TagRepository struct {
	db *sql.DB
}

// NewTagRepositoryは新しいTagRepositoryを初期化し構造体のポインタを返します
func NewTagRepository(db *sql.DB) *TagRepository {
	return &TagRepository{
		db: db,
	}
}

// GetByIDはIDを指定してタグを取得します
func (tr *TagRepository) GetByID(id string) (*entity.Tag, error) {
	tag := &entity.Tag{}
	err := tr.db.QueryRow("SELECT * FROM tags WHERE id = ?", id).Scan(&tag.ID, &tag.Name)
	if err != nil {
		return nil, err
	}
	return tag, nil
}

// GetByNameはNameを指定してタグを取得します
func (tr *TagRepository) GetByName(name string) (*entity.Tag, error) {
	tag := &entity.Tag{}
	err := tr.db.QueryRow("SELECT * FROM tags WHERE name = ?", name).Scan(&tag.ID, &tag.Name)
	if err != nil {
		return nil, err
	}
	return tag, nil
}

// GetAllは全てのタグを取得します
func (tr *TagRepository) GetAll() (*entity.Tags, error) {
	rows, err := tr.db.Query("SELECT * FROM tags")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags := &entity.Tags{}
	for rows.Next() {
		tag := &entity.Tag{}
		err := rows.Scan(&tag.ID, &tag.Name)
		if err != nil {
			return nil, err
		}
		*tags = append(*tags, tag)
	}
	return tags, nil
}

// Createはタグを作成します
func (tr *TagRepository) CreateTag(tag *entity.Tag) error {
	_, err := tr.db.Exec("INSERT INTO tags (id, name) VALUES (?, ?)", tag.ID, tag.Name)
	if err != nil {
		return err
	}
	return nil
}

// Deleteはタグを削除します
func (tr *TagRepository) DeleteTag(id string) error {
	_, err := tr.db.Exec("DELETE FROM tags WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}