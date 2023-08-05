package persistence

import (
	"database/sql"
	"github.com/Uwasaru/Sayka/domain/entity"
	"github.com/Uwasaru/Sayka/domain/repository"
)

var _ repository.UserRepository = &UserRepository{}

// UserRepositoryはユーザーに関する永続化を担当します
type UserRepository struct {
	db *sql.DB
}

// NewUserRepositoryは新しいUserRepositoryを初期化し構造体のポインタを返します
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// GetByIDはIDを指定してユーザーを取得します
func (ur *UserRepository) GetByID(id string) (*entity.User, error) {
	user := &entity.User{}
	err := ur.db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.GithubUrl, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetByEmailはEmailを指定してユーザーを取得します
func (ur *UserRepository) GetByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	err := ur.db.QueryRow("SELECT * FROM users WHERE email = ?", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.GithubUrl, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUserはユーザーを作成します
func (ur *UserRepository) CreateUser(user *entity.User) error {
	_, err := ur.db.Exec("INSERT INTO users (id, name, email, password, github_url, created_at) VALUES (?, ?, ?, ?, ?, ?)", user.ID, user.Name, user.Email, user.Password, user.GithubUrl, user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUserはユーザーを更新します
func (ur *UserRepository) UpdateUser(user *entity.User) error {
	_, err := ur.db.Exec("UPDATE users SET name = ?, email = ?, password = ?, github_url = ? WHERE id = ?", user.Name, user.Email, user.Password, user.GithubUrl, user.ID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUserはユーザーを削除します
func (ur *UserRepository) DeleteUser(id string) error {
	_, err := ur.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
