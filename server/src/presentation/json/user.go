package json

import (
	"github.com/Uwasaru/Sayka/domain/entity"
)

type UserJson struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Img  string `json:"img"`
}

func UserEntityToJson(c *entity.User) *UserJson {
	return &UserJson{
		ID:   c.ID,
		Name: c.Name,
		Img:  c.Img,
	}
}

func UserJsonToEntity(j *UserJson) *entity.User {
	return &entity.User{
		ID:   j.ID,
		Name: j.Name,
		Img:  j.Img,
	}
}
