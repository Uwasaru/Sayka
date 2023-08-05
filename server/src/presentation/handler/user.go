package handler

import (
	"net/http"

	"github.com/Uwasaru/Sayka/presentation/json"
	"github.com/Uwasaru/Sayka/usecase"
	"github.com/Uwasaru/Sayka/utils"
	"github.com/gin-gonic/gin"
)

// Handlerはハンドラーを表します
type UserHandler struct {
	uu usecase.IUserUsecase
}

// NewUserHandlerは新しいUserHandlerを初期化し構造体のポインタを返します
func NewUserHandler(uu usecase.IUserUsecase) *UserHandler {
	return &UserHandler{
		uu: uu,
	}
}

// GetByIDはIDを指定してユーザーを取得します
func (uh *UserHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	user, err := uh.uu.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetByGithubIdはgithubIdを指定してユーザーを取得します
func (uh *UserHandler) GetByGithubId(c *gin.Context) {
	githubId := c.Param("githubId")
	user, err := uh.uu.GetByGithubId(githubId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUserはユーザーを作成します
func (uh *UserHandler) CreateUser(c *gin.Context) {
	user_json := &json.UserJson{}
	if err := c.BindJSON(user_json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user_json.ID = utils.Ulid()
	user_json.AccessToken = utils.EncodeToken(user_json.AccessToken)
	user_json.TokenExpire = utils.GetExpireTimeStamps()
	user := json.UserJsonToEntity(user_json)
	if err := uh.uu.CreateUser(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUserはユーザーを更新します
func (uh *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	user, err := uh.uu.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user_json := &json.UserJson{}
	if err := c.BindJSON(user_json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user_json.ID = user.ID
	user_json.CreatedAt = user.CreatedAt
	user_json.TokenExpire = user.TokenExpire
	user_json.AccessToken = utils.EncodeToken(user_json.AccessToken)
	user = json.UserJsonToEntity(user_json)
	if err := uh.uu.UpdateUser(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUserはユーザーを削除します
func (uh *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := uh.uu.DeleteUser(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// LoginUserはユーザーをログインします
func (uh *UserHandler) LoginUser(c *gin.Context) {
	user_json := &json.UserJson{}
	if err := c.BindJSON(user_json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := uh.uu.GetByGithubId(user_json.GithubID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
