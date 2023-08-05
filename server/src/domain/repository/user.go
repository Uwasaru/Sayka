package repository

import (
	"github.com/Uwasaru/Sayka/domain/entity"
)

// UserRepositoryはユーザーに関する永続化を担当します
type UserRepository interface {
	// GetByIDはIDを指定してユーザーを取得します
	GetByID(id string) (*entity.User, error)
	// GetByEmailはEmailを指定してユーザーを取得します
	GetByEmail(email string) (*entity.User, error)
	// Createはユーザーを作成します
	CreateUser(user *entity.User) error
	// UpdateUserはユーザーを更新します
	UpdateUser(user *entity.User) error
	// DeleteUserはユーザーを削除します
	DeleteUser(id string) error
}
