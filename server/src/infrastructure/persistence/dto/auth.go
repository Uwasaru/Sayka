package dto

import (
	"time"

	"github.com/Uwasaru/Sayka/domain/entity"
)

type LoginSessionsDto struct {
	ID     string    `db:"id"`
	UserID string    `db:"user_id"`
	Expiry time.Time `db:"expiry"`
}

type AuthStatesDto struct {
	State       string `db:"state"`
	RedirectURL string `db:"redirect_url"`
}

type GithubAuthDto struct {
	UserID       string    `db:"user_id"`
	AccessToken  string    `db:"access_token"`
	RefreshToken string    `db:"refresh_token"`
	Expiry       time.Time `db:"expiry"`
}

func AuthStatesEntityToDto(u *entity.AuthStates) AuthStatesDto {
	return AuthStatesDto{
		State:       u.State,
		RedirectURL: u.RedirectURL,
	}
}

func AuthStatesDtoToEntity(dto *AuthStatesDto) *entity.AuthStates {
	return &entity.AuthStates{
		State:       dto.State,
		RedirectURL: dto.RedirectURL,
	}
}
