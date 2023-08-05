package persistence

import (
	"database/sql"
	"github.com/Uwasaru/Sayka/domain/entity"
	"github.com/Uwasaru/Sayka/domain/repository"
)

var _ repository.FavoriteRepository = &FavoriteRepository{}

// FavoriteRepositoryはFavoriteRepositoryの実装です
type FavoriteRepository struct {
	db *sql.DB
}

// NewFavoriteRepositoryは新しいFavoriteRepositoryを初期化し構造体のポインタを返します

func NewFavoriteRepository(db *sql.DB) *FavoriteRepository {
	return &FavoriteRepository{
		db: db,
	}
}

// GetByIDはIDを指定して投稿を取得します
func (fr *FavoriteRepository) GetByID(id string) (*entity.Favorite, error) {
	favorite := &entity.Favorite{}
	err := fr.db.QueryRow("SELECT * FROM favorites WHERE id = ?", id).Scan(&favorite.ID, &favorite.UserID, &favorite.PostID)
	if err != nil {
		return nil, err
	}
	return favorite, nil
}

// GetByUserIDはUserIDを指定して投稿を取得します
func (fr *FavoriteRepository) GetByUserID(userID string) (*entity.Favorites, error) {
	rows, err := fr.db.Query("SELECT * FROM favorites WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	favorites := &entity.Favorites{}
	for rows.Next() {
		favorite := &entity.Favorite{}
		err := rows.Scan(&favorite.ID, &favorite.UserID, &favorite.PostID)
		if err != nil {
			return nil, err
		}
		*favorites = append(*favorites, favorite)
	}
	return favorites, nil
}

// GetByPostIDはPostIDを指定して投稿を取得します
func (fr *FavoriteRepository) GetByPostID(postID string) (*entity.Favorites, error) {
	rows, err := fr.db.Query("SELECT * FROM favorites WHERE post_id = ?", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	favorites := &entity.Favorites{}
	for rows.Next() {
		favorite := &entity.Favorite{}
		err := rows.Scan(&favorite.ID, &favorite.UserID, &favorite.PostID)
		if err != nil {
			return nil, err
		}
		*favorites = append(*favorites, favorite)
	}
	return favorites, nil
}

// GetAllは全ての投稿を取得します
func (fr *FavoriteRepository) GetAll() (*entity.Favorites, error) {
	rows, err := fr.db.Query("SELECT * FROM favorites")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	favorites := &entity.Favorites{}
	for rows.Next() {
		favorite := &entity.Favorite{}
		err := rows.Scan(&favorite.ID, &favorite.UserID, &favorite.PostID)
		if err != nil {
			return nil, err
		}
		*favorites = append(*favorites, favorite)
	}
	return favorites, nil
}

// Createは投稿を作成します
func (fr *FavoriteRepository) CreateFavorite(favorite *entity.Favorite) error {
	_, err := fr.db.Exec("INSERT INTO favorites (id, user_id, post_id) VALUES (?, ?, ?)", favorite.ID, favorite.UserID, favorite.PostID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteFavoriteは投稿を削除します
func (fr *FavoriteRepository) DeleteFavorite(id string) error {
	_, err := fr.db.Exec("DELETE FROM favorites WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
