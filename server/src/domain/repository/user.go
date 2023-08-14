package repository

import (
	"context"

	"github.com/Uwasaru/Sayka/domain/entity"
)

// UserRepositoryはユーザーに関する永続化を担当します
type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetUser(ctx context.Context, id int) (*entity.User, error)
}
