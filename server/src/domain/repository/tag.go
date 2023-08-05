package repository

import (
	"github.com/Uwasaru/Sayka/domain/entity"
)


// TagRepositoryはタグに関する永続化を担当します

type TagRepository interface {
	// GetByIDはIDを指定してタグを取得します
	GetByID(id string) (*entity.Tag, error)
	// GetByNameはNameを指定してタグを取得します
	GetByName(name string) (*entity.Tag, error)
	// GetAllは全てのタグを取得します
	GetAll() (*entity.Tags, error)
	// Createはタグを作成します
	CreateTag(tag *entity.Tag) error
	// Deleteはタグを削除します
	DeleteTag(id string) error
}