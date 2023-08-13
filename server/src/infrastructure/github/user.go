package github

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Uwasaru/Sayka/domain/entity"
	gj "github.com/Uwasaru/Sayka/presentation/json"
	mycontext "github.com/Uwasaru/Sayka/utils/context"
)

func (c *Client) GetMe(ctx context.Context) (*entity.User, error) {
	token, ok := mycontext.GetToken(ctx)
	if !ok {
		return nil, fmt.Errorf("token not found")
	}
	// tokenを使用して、clientを返す
	client := c.auth.Config.Client(ctx, token)
	resp, err := client.Get(os.Getenv("GITHUB_GETME"))
	if err != nil {
		return nil, fmt.Errorf("GITHUB API Get: %w", err)
	}
	defer resp.Body.Close()
	var j gj.GithubUserJson
	if err := json.NewDecoder(resp.Body).Decode(&j); err != nil {
		return nil, fmt.Errorf("decode: %w", err)
	}
	fmt.Println("gitから取得してきたデータ", j)
	user := &entity.User{
		ID:   j.ID,
		Name: j.Name,
		Img:  j.AvatarUrl,
	}
	return user, nil
}
