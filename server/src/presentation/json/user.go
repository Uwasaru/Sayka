package json

import (
	"github.com/Uwasaru/Sayka/domain/entity"
	"time"
)

type UserJson struct {
	ID        string    `json:"id"`
	GithubID	string    `json:"github_id"`
	GithubIcon	string    `json:"github_icon"`
	AccessToken	string    `json:"access_token"`
	RefleshToken	string    `json:"reflesh_token"`
	TokenExpire time.Time `json:"token_expire"`
	CreatedAt time.Time `json:"created_at"`
}

func UserJsonToEntity(userJson *UserJson) *entity.User {
	return &entity.User{
		ID:        userJson.ID,
		GithubID:	userJson.GithubID,
		GithubIcon:	userJson.GithubIcon,
		AccessToken:	userJson.AccessToken,
		RefleshToken:	userJson.RefleshToken,
		TokenExpire: userJson.TokenExpire,
		CreatedAt: userJson.CreatedAt,
	}
}

func UserEntityToJson(user *entity.User) *UserJson {
	return &UserJson{
		ID:        user.ID,
		GithubID:	user.GithubID,
		GithubIcon:	user.GithubIcon,
		AccessToken:	user.AccessToken,
		RefleshToken:	user.RefleshToken,
		TokenExpire: user.TokenExpire,
		CreatedAt: user.CreatedAt,
	}
}
