package json

type GithubUserJson struct {
	ID        int `json:"id"`
	Name      string `json:"login"`
	AvatarUrl string `json:"avatar_url"`
}
