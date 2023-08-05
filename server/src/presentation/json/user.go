package json

import (
	"github.com/Uwasaru/Sayka/domain/entity"
	"time"
)

type UserJson struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	GithubUrl string    `json:"github_url"`
	CreatedAt time.Time `json:"created_at"`
}

func UserJsonToEntity(userJson *UserJson) *entity.User {
	return &entity.User{
		ID:        userJson.ID,
		Name:      userJson.Name,
		Email:     userJson.Email,
		Password:  userJson.Password,
		GithubUrl: userJson.GithubUrl,
		CreatedAt: userJson.CreatedAt,
	}
}

func UserEntityToJson(user *entity.User) *UserJson {
	return &UserJson{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		GithubUrl: user.GithubUrl,
		CreatedAt: user.CreatedAt,
	}
}
