package persistence

import (
	"context"

	"github.com/Uwasaru/Sayka/domain/entity"
	"github.com/Uwasaru/Sayka/domain/repository"
	"github.com/Uwasaru/Sayka/infrastructure/database"
	d "github.com/Uwasaru/Sayka/infrastructure/persistence/dto"
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
func (pr *PostRepository) GetByID(ctx context.Context, id string) (*entity.Post, error) {
	query := `
	SELECT *
	FROM posts
	WHERE id = ?
	`
	var dto d.PostDto
	err := pr.conn.DB.GetContext(ctx, &dto, query, id)
	if err != nil {
		return nil, err
	}
	query = `
	SELECT name
	FROM tags
	WHERE post_id = ?
	`
	rows, err := pr.conn.DB.QueryContext(ctx, query, dto.ID)
	if err != nil {
		return nil, err
	}
	var tags []string
	for rows.Next() {
		var tag string
		err := rows.Scan(&tag)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	post := d.PostDtoToEntity(&dto)
	post.Tags = tags
	return post, nil
}

// GetByUserIDはUserIDを指定して投稿を取得します
func (pr *PostRepository) GetByUserID(ctx context.Context, userID string) (*entity.Posts, error) {
	query := `
	SELECT *
	FROM posts
	WHERE user_id = ?
	`
	rows, err := pr.conn.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	var posts entity.Posts
	for rows.Next() {
		var dto d.PostDto
		err := rows.Scan(&dto.ID, &dto.Title, &dto.UserID, &dto.GithubUrl, &dto.AppUrl, &dto.SlideUrl, &dto.Description, &dto.CreatedAt)
		if err != nil {
			return nil, err
		}
		query := `
		SELECT name
		FROM tags
		WHERE post_id = ?
		`
		rows1, err := pr.conn.DB.QueryContext(ctx, query, dto.ID)
		if err != nil {
			return nil, err
		}
		var tags []string
		for rows1.Next() {
			var tag string
			err := rows1.Scan(&tag)
			if err != nil {
				return nil, err
			}
			tags = append(tags, tag)
		}
		post := d.PostDtoToEntity(&dto)
		post.Tags = tags
		posts = append(posts, d.PostDtoToEntity(&dto))
	}
	return &posts, nil
}

// GetAllは全ての投稿を取得します
func (pr *PostRepository) GetAll(ctx context.Context) (*entity.Posts, error) {
	query := `
	SELECT *
	FROM posts
	`
	rows, err := pr.conn.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var posts entity.Posts
	for rows.Next() {
		var dto d.PostDto
		err := rows.Scan(&dto.ID, &dto.Title, &dto.UserID, &dto.GithubUrl, &dto.AppUrl, &dto.SlideUrl, &dto.Description, &dto.CreatedAt)
		if err != nil {
			return nil, err
		}
		query := `
		SELECT name
		FROM tags
		WHERE post_id = ?
		`
		rows, err := pr.conn.DB.QueryContext(ctx, query, dto.ID)
		if err != nil {
			return nil, err
		}
		var tags []string
		for rows.Next() {
			var tag string
			err := rows.Scan(&tag)
			if err != nil {
				return nil, err
			}
			tags = append(tags, tag)
		}
		post := d.PostDtoToEntity(&dto)
		post.Tags = tags
		posts = append(posts, post)
	}
	return &posts, nil
}

// CreatePostは投稿を作成します
func (pr *PostRepository) CreatePost(ctx context.Context, post *entity.Post) error {
	query := `
	INSERT INTO posts (id, user_id, title, github_url, app_url, slide_url, description)
	VALUES (:id,:user_id,:title,:github_url,:app_url,:slide_url,:description)
	`
	dto := d.PostEntityToDto(post)

	_, err := pr.conn.DB.NamedExecContext(ctx, query, &dto)
	if err != nil {
		return err
	}

	tags := post.Tags
	for _, tag := range tags {
		query := `
		INSERT INTO tags (post_id, name)
		VALUES (:post_id, :name)
		`
		dto1 := d.TagEntityToDto(&entity.Tag{
			PostID: dto.ID,
			Name:  tag,
		})
		_, err := pr.conn.DB.NamedExecContext(ctx, query, &dto1)
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdatePostは投稿を更新します
func (pr *PostRepository) UpdatePost(ctx context.Context, post *entity.Post) error {
	query := `
	UPDATE posts
	SET title = :title, github_url = :github_url, app_url = :app_url, slide_url = :slide_url, description = :description
	WHERE id = :id
	`
	dto := d.PostEntityToDto(post)

	_, err := pr.conn.DB.NamedExecContext(ctx, query, &dto)
	if err != nil {
		return err
	}
	return nil
}

// DeletePostは投稿を削除します
func (pr *PostRepository) DeletePost(ctx context.Context, id string) error {
	query := `
	DELETE FROM posts
	WHERE id = :id
	`

	_, err := pr.conn.DB.NamedExecContext(ctx, query, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	return nil
}
