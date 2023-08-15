package json

type GithubUserJson struct {
	ID        string `json:"login"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
}
