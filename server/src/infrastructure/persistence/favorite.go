package persistence

import (
	"context"

	"github.com/Uwasaru/Sayka/domain/entity"
	"github.com/Uwasaru/Sayka/domain/repository"
	"github.com/Uwasaru/Sayka/infrastructure/database"
	d "github.com/Uwasaru/Sayka/infrastructure/persistence/dto"
)

var _ repository.FavoriteRepository = &FavoriteRepository{}

// FavoriteRepositoryはFavoriteRepositoryの実装です
type FavoriteRepository struct {
	conn *database.Conn
}

// NewFavoriteRepositoryは新しいFavoriteRepositoryを初期化し構造体のポインタを返します

func NewFavoriteRepository(conn *database.Conn) *FavoriteRepository {
	return &FavoriteRepository{
		conn: conn,
	}
}

// GetByIDはIDを指定して投稿を取得します
func (fr *FavoriteRepository) GetByID(ctx context.Context, id string) (*entity.Favorite, error) {
	query := `
	SELECT *
	FROM favorites
	WHERE id = ?
	`
	var dto d.FavoriteDto
	err := fr.conn.DB.GetContext(ctx, &dto, query, id)
	if err != nil {
		return nil, err
	}
	return d.FavoriteDtoToEntity(&dto), nil
}

// GetByUserIDはUserIDを指定して投稿を取得します
func (fr *FavoriteRepository) GetByUserID(ctx context.Context, userID string) (*entity.Favorites, error) {
	query := `
	SELECT *
	FROM favorites
	WHERE user_id = ?
	`
	rows, err := fr.conn.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	var favorites entity.Favorites
	for rows.Next() {
		var dto d.FavoriteDto
		err := rows.Scan(&dto.ID, &dto.UserID, &dto.PostID)
		if err != nil {
			return nil, err
		}
		favorites = append(favorites, d.FavoriteDtoToEntity(&dto))
	}
	return &favorites, nil
}

// GetByPostIDはPostIDを指定して投稿を取得します
func (fr *FavoriteRepository) GetByPostID(ctx context.Context, postID string) (entity.FavoriteUsers, error) {
	query := `
	SELECT user_id
	FROM favorites
	WHERE post_id = ?
	`
	rows, err := fr.conn.DB.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}

	var favorites = make(entity.FavoriteUsers, 0)
	for rows.Next() {
		var dto d.FavoriteUserDto
		err := rows.Scan(&dto.UserID)
		if err != nil {
			return nil, err
		}
		favorites = append(favorites, d.FavoriteUserDtoToEntity(&dto))
	}
	return favorites, nil
}

// GetAllは全ての投稿を取得します
func (fr *FavoriteRepository) GetAll(ctx context.Context) (*entity.Favorites, error) {
	query := `
	SELECT *
	FROM favorites
	`
	rows, err := fr.conn.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var favorites = entity.Favorites{}
	for rows.Next() {
		var dto d.FavoriteDto
		err := rows.Scan(&dto.ID, &dto.UserID, &dto.PostID)
		if err != nil {
			return nil, err
		}
		favorites = append(favorites, d.FavoriteDtoToEntity(&dto))
	}
	return &favorites, nil
}

// Createは投稿を作成します
func (fr *FavoriteRepository) CreateFavorite(ctx context.Context, favorite *entity.Favorite) error {
	query := `
	INSERT INTO favorites (
		id,
		user_id,
		post_id
	) VALUES (
		:id,
		:user_id,
		:post_id
	)
	`

	dto := d.FavoriteEntityToDto(favorite)
	_, err := fr.conn.DB.NamedExecContext(ctx, query, &dto)
	if err != nil {
		return err
	}
	return nil
}

// DeleteFavoriteは投稿を削除します
func (fr *FavoriteRepository) DeleteFavorite(ctx context.Context, id string) error {
	query := `
	DELETE FROM favorites
	WHERE id = :id
	`
	_, err := fr.conn.DB.NamedExecContext(ctx, query, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	return nil
}
