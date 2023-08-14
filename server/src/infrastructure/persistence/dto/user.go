package dto

import "github.com/Uwasaru/Sayka/domain/entity"

type UserDto struct {
	ID   int `db:"id"`
	Name string `db:"name"`
	Img  string `db:"img"`
}

func UserDtoToEntity(dto *UserDto) *entity.User {
	return &entity.User{
		ID:   dto.ID,
		Name: dto.Name,
		Img:  dto.Img,
	}
}

func UserEntityToDto(u *entity.User) UserDto {
	return UserDto{
		ID:   u.ID,
		Name: u.Name,
		Img:  u.Img,
	}
}
