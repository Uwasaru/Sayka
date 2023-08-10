package persistence

import (
	"github.com/Uwasaru/Sayka/domain/entity"
	"github.com/Uwasaru/Sayka/domain/repository"
	"github.com/Uwasaru/Sayka/infrastructure/database"
)

var _ repository.PostRepository = &PostRepository{}

// PostRepositoryは投稿に関する永続化を担当します
type PostRepository struct {
	conn *database.Conn
}

// NewPostRepositoryは新しいPostRepositoryを初期化し構造体のポインタを返します
func NewPostRepository(conn *database.Conn) *PostRepository {
	return &PostRepository{
		conn: conn,
	}
}

// GetByIDはIDを指定して投稿を取得します
func (pr *PostRepository) GetByID(id string) (*entity.Post, error) {
	post := &entity.Post{}
	err := pr.db.QueryRow("SELECT * FROM posts WHERE id = ?", id).Scan(&post.ID, &post.UserID, &post.Title, &post.GithubUrl, &post.AppUrl, &post.SlideUrl, &post.Description, &post.CreatedAt)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// GetByUserIDはUserIDを指定して投稿を取得します
func (pr *PostRepository) GetByUserID(userID string) (*entity.Posts, error) {
	rows, err := pr.db.Query("SELECT * FROM posts WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := &entity.Posts{}
	for rows.Next() {
		post := &entity.Post{}
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.GithubUrl, &post.AppUrl, &post.SlideUrl, &post.Description, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		*posts = append(*posts, post)
	}
	return posts, nil
}

// GetAllは全ての投稿を取得します
func (pr *PostRepository) GetAll() (*entity.Posts, error) {
	rows, err := pr.db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := &entity.Posts{}
	for rows.Next() {
		post := &entity.Post{}
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.GithubUrl, &post.AppUrl, &post.SlideUrl, &post.Description, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		*posts = append(*posts, post)
	}
	return posts, nil
}

// CreatePostは投稿を作成します
func (pr *PostRepository) CreatePost(post *entity.Post) error {
	_, err := pr.db.Exec("INSERT INTO posts (id, user_id, title, github_url, app_url, slide_url, description, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", post.ID, post.UserID, post.Title, post.GithubUrl, post.AppUrl, post.SlideUrl, post.Description, post.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

// UpdatePostは投稿を更新します
func (pr *PostRepository) UpdatePost(post *entity.Post) error {
	_, err := pr.db.Exec("UPDATE posts SET title = ?, github_url = ?, app_url = ?, slide_url = ?, description = ? WHERE id = ?", post.Title, post.GithubUrl, post.AppUrl, post.SlideUrl, post.Description, post.ID)
	if err != nil {
		return err
	}
	return nil
}

// DeletePostは投稿を削除します
func (pr *PostRepository) DeletePost(id string) error {
	_, err := pr.db.Exec("DELETE FROM posts WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
