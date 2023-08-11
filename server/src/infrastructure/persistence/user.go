package persistence

import (
	"github.com/Uwasaru/Sayka/domain/entity"
	"github.com/Uwasaru/Sayka/domain/repository"
	"github.com/Uwasaru/Sayka/infrastructure/database"
)

var _ repository.UserRepository = &UserRepository{}

// UserRepositoryはユーザーに関する永続化を担当します
type UserRepository struct {
	conn *database.Conn
}

// NewUserRepositoryは新しいUserRepositoryを初期化し構造体のポインタを返します
func NewUserRepository(conn *database.Conn) *UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

// GetByIDはIDを指定してユーザーを取得します
func (ur *UserRepository) GetByID(id string) (*entity.User, error) {
	user := &entity.User{}
	err := ur.db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.GithubID, &user.GithubIcon, &user.AccessToken, &user.RefleshToken, &user.TokenExpire, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetByGithubIdはEGithubIDを指定してユーザーを取得します
func (ur *UserRepository) GetByGithubId(GithubID string) (*entity.User, error) {
	user := &entity.User{}
	err := ur.db.QueryRow("SELECT * FROM users WHERE github_id = ?", GithubID).Scan(&user.ID, &user.GithubID, &user.GithubIcon, &user.AccessToken, &user.RefleshToken, &user.TokenExpire, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUserはユーザーを作成します
func (ur *UserRepository) CreateUser(user *entity.User) error {
	_, err := ur.db.Exec("INSERT INTO users (id, github_id, github_icon, access_token, reflesh_token, token_expire) VALUES (?, ?, ?, ?, ?, ?)", user.ID, user.GithubID, user.GithubIcon, user.AccessToken, user.RefleshToken, user.TokenExpire)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUserはユーザーを更新します
func (ur *UserRepository) UpdateUser(user *entity.User) error {
	_, err := ur.db.Exec("UPDATE users SET github_id = ?, github_icon = ?, access_token = ?, reflesh_token = ?, token_expire = ? WHERE id = ?", user.GithubID, user.GithubIcon, user.AccessToken, user.RefleshToken, user.TokenExpire, user.ID)
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
