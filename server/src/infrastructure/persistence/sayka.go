package persistence

import (
	"context"

	"github.com/Uwasaru/Sayka/domain/entity"
	"github.com/Uwasaru/Sayka/domain/repository"
	"github.com/Uwasaru/Sayka/infrastructure/database"
	d "github.com/Uwasaru/Sayka/infrastructure/persistence/dto"
)

var _ repository.SaykaRepository = &SaykaRepository{}

// SaykaRepositoryは投稿に関する永続化を担当します
type SaykaRepository struct {
	conn *database.Conn
}

// NewSaykaRepositoryは新しいSaykaRepositoryを初期化し構造体のポインタを返します
func NewSaykaRepository(conn *database.Conn) *SaykaRepository {
	return &SaykaRepository{
		conn: conn,
	}
}

// GetByIDはIDを指定して投稿を取得します
func (pr *SaykaRepository) GetByID(ctx context.Context, id string) (*entity.Sayka, error) {
	query := `
	SELECT *
	FROM saykas
	WHERE id = ?
	`
	var dto d.SaykaDto
	err := pr.conn.DB.GetContext(ctx, &dto, query, id)
	if err != nil {
		return nil, err
	}
	query = `
	SELECT name
	FROM tags
	WHERE sayka_id = ?
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
	sayka := d.SaykaDtoToEntity(&dto)
	sayka.Tags = tags
	return sayka, nil
}

// GetByUserIDはUserIDを指定して投稿を取得します
func (pr *SaykaRepository) GetByUserID(ctx context.Context, userID string) (*entity.Saykas, error) {
	query := `
	SELECT *
	FROM saykas
	WHERE user_id = ?
	`
	rows, err := pr.conn.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	var saykas = entity.Saykas{}
	for rows.Next() {
		var dto d.SaykaDto
		err := rows.Scan(&dto.ID, &dto.Title, &dto.UserID, &dto.GithubUrl, &dto.AppUrl, &dto.SlideUrl, &dto.ArticleUrl, &dto.FigmaUrl, &dto.Description, &dto.CreatedAt, &dto.UpdatedAt)
		if err != nil {
			return nil, err
		}
		query := `
		SELECT name
		FROM tags
		WHERE sayka_id = ?
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
		sayka := d.SaykaDtoToEntity(&dto)
		sayka.Tags = tags
		saykas = append(saykas, d.SaykaDtoToEntity(&dto))
	}
	return &saykas, nil
}

// GetAllは全ての投稿を取得します
func (pr *SaykaRepository) GetAll(ctx context.Context) (*entity.Saykas, error) {
	query := `
	SELECT *
	FROM saykas
	ORDER BY id DESC
	`
	rows, err := pr.conn.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var saykas = entity.Saykas{}
	for rows.Next() {
		var dto d.SaykaDto
		err := rows.Scan(&dto.ID, &dto.Title, &dto.UserID, &dto.GithubUrl, &dto.AppUrl, &dto.SlideUrl, &dto.ArticleUrl, &dto.FigmaUrl, &dto.Description, &dto.CreatedAt, &dto.UpdatedAt)
		if err != nil {
			return nil, err
		}
		query := `
		SELECT name
		FROM tags
		WHERE sayka_id = ?
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
		sayka := d.SaykaDtoToEntity(&dto)
		sayka.Tags = tags
		saykas = append(saykas, sayka)
	}
	return &saykas, nil
}

// GetTimeLineはタイムラインを取得します
func (pr *SaykaRepository) GetTimeLine(ctx context.Context, id string, tag string) (*entity.Saykas, error) {
	limit := 10
	var query string
	var args []interface{}

	if tag != "" {
		query = `
		SELECT p.*
		FROM saykas p
		JOIN tags t ON p.id = t.sayka_id
		WHERE p.id < ? AND t.name LIKE ?
		ORDER BY p.id DESC
		LIMIT ?
		`
		args = append(args, id, tag, limit)
	} else {
		query = `
		SELECT *
		FROM saykas
		WHERE id < ?
		ORDER BY id DESC
		LIMIT ?
		`
		args = append(args, id, limit)
	}

	rows, err := pr.conn.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var saykas = entity.Saykas{}
	for rows.Next() {
		var dto d.SaykaDto
		err := rows.Scan(&dto.ID, &dto.Title, &dto.UserID, &dto.GithubUrl, &dto.AppUrl, &dto.SlideUrl, &dto.ArticleUrl, &dto.FigmaUrl, &dto.Description, &dto.CreatedAt, &dto.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tagQuery := `
        SELECT name
        FROM tags
        WHERE sayka_id = ?
        `
		tagRows, err := pr.conn.DB.QueryContext(ctx, tagQuery, dto.ID)
		if err != nil {
			return nil, err
		}

		var tags []string
		for tagRows.Next() {
			var tag string
			err := tagRows.Scan(&tag)
			if err != nil {
				return nil, err
			}
			tags = append(tags, tag)
		}
		tagRows.Close()

		sayka := d.SaykaDtoToEntity(&dto)
		sayka.Tags = tags
		saykas = append(saykas, sayka)
	}
	rows.Close()

	return &saykas, nil
}

// GetAllFavoritedSaykaはいいねした全ての投稿を取得します
func (pr *SaykaRepository) GetAllFavoritedSayka(ctx context.Context, myId string) (*entity.Saykas, error) {
	limit := 10
	query := `
	SELECT s.*
	FROM saykas s
	JOIN favorites f ON s.id = f.sayka_id
	WHERE f.user_id = ?
	ORDER BY s.id DESC
	LIMIT ?
	`
	rows, err := pr.conn.DB.QueryContext(ctx, query, myId, limit)
	if err != nil {
		return nil, err
	}


	var saykas = entity.Saykas{}
	for rows.Next() {
		var dto d.SaykaDto
		err := rows.Scan(&dto.ID, &dto.Title, &dto.UserID, &dto.GithubUrl, &dto.AppUrl, &dto.SlideUrl, &dto.ArticleUrl, &dto.FigmaUrl, &dto.Description, &dto.CreatedAt, &dto.UpdatedAt)
		if err != nil {
			return nil, err
		}
		query := `
		SELECT name
		FROM tags
		WHERE sayka_id = ?
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
		sayka := d.SaykaDtoToEntity(&dto)
		sayka.Tags = tags
		saykas = append(saykas, sayka)
	}
	return &saykas, nil
}

// GetAllCommentedSaykaはコメントした全ての投稿を取得します
func (pr *SaykaRepository) GetAllCommentedSayka(ctx context.Context, myId string) (*entity.Saykas, error) {
	limit := 10
	query := `
	SELECT s.*
	FROM saykas s
	JOIN comments c ON s.id = c.sayka_id
	WHERE c.user_id = ?
	ORDER BY s.id DESC
	LIMIT ?
	`
	rows, err := pr.conn.DB.QueryContext(ctx, query, myId, limit)
	if err != nil {
		return nil, err
	}

	var saykas = entity.Saykas{}
	for rows.Next() {
		var dto d.SaykaDto
		err := rows.Scan(&dto.ID, &dto.Title, &dto.UserID, &dto.GithubUrl, &dto.AppUrl, &dto.SlideUrl, &dto.ArticleUrl, &dto.FigmaUrl, &dto.Description, &dto.CreatedAt, &dto.UpdatedAt)
		if err != nil {
			return nil, err
		}
		query := `
		SELECT name
		FROM tags
		WHERE sayka_id = ?
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
		sayka := d.SaykaDtoToEntity(&dto)
		sayka.Tags = tags
		saykas = append(saykas, sayka)
	}
	return &saykas, nil
}

// CreateSaykaは投稿を作成します
func (pr *SaykaRepository) CreateSayka(ctx context.Context, sayka *entity.Sayka) error {
	query := `
	INSERT INTO saykas (id, user_id, title, github_url, app_url, slide_url, article_url, figma_url, description)
	VALUES (:id,:user_id,:title,:github_url,:app_url,:slide_url,:article_url,:figma_url,:description)
	`
	dto := d.SaykaEntityToDto(sayka)

	_, err := pr.conn.DB.NamedExecContext(ctx, query, &dto)
	if err != nil {
		return err
	}

	tags := sayka.Tags
	for _, tag := range tags {
		query := `
		INSERT INTO tags (sayka_id, name)
		VALUES (:sayka_id, :name)
		`
		dto1 := d.TagEntityToDto(&entity.Tag{
			SaykaID: dto.ID,
			Name:    tag,
		})
		_, err := pr.conn.DB.NamedExecContext(ctx, query, &dto1)
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdateSaykaは投稿を更新します
func (pr *SaykaRepository) UpdateSayka(ctx context.Context, sayka *entity.Sayka) error {
	query := `
	UPDATE saykas
	SET title = :title, github_url = :github_url, app_url = :app_url, slide_url = :slide_url, description = :description
	WHERE id = :id
	`
	dto := d.SaykaEntityToDto(sayka)

	_, err := pr.conn.DB.NamedExecContext(ctx, query, &dto)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSaykaは投稿を削除します
func (pr *SaykaRepository) DeleteSayka(ctx context.Context, id string) error {
	query := `
	DELETE FROM saykas
	WHERE id = :id
	`

	_, err := pr.conn.DB.NamedExecContext(ctx, query, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	return nil
}
