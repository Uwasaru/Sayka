package persistence

import (
	"context"

	"github.com/Uwasaru/Sayka/domain/entity"
	"github.com/Uwasaru/Sayka/domain/repository"
	"github.com/Uwasaru/Sayka/infrastructure/database"
	d "github.com/Uwasaru/Sayka/infrastructure/persistence/dto"
)

var _ repository.UserRepository = &UserRepository{}

type UserRepository struct {
	conn *database.Conn
}

func NewUserRepository(conn *database.Conn) repository.UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	query := `
	INSERT INTO users (id, name,img)
	VALUES (:id,:name,:img)
	`
	dto := d.UserEntityToDto(user)

	_, err := ur.conn.DB.NamedExecContext(ctx, query, &dto)
	if err != nil {
		return nil, err
	}
	return d.UserDtoToEntity(&dto), nil
}

func (ur *UserRepository) DeleteUser(ctx context.Context, id int) error {
	query := `
	DELETE FROM users
	WHERE id = :id
	`

	_, err := ur.conn.DB.NamedExecContext(ctx, query, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetUser(ctx context.Context, id int) (*entity.User, error) {
	query := `
	SELECT *
	FROM users
	WHERE id = ?
	`
	var dto d.UserDto
	err := ur.conn.DB.GetContext(ctx, &dto, query, id)
	if err != nil {
		return nil, err
	}
	return d.UserDtoToEntity(&dto), nil
}
