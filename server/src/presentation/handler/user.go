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

// GetByEmailはEmailを指定してユーザーを取得します
func (uh *UserHandler) GetByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := uh.uu.GetByEmail(email)
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
	user_json.Password, _ = utils.HashPassword(user_json.Password)
	user_json.CreatedAt = utils.GetCurrentTimeStamps()
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
	if user_json.Password != "" {
		user.Password, _ = utils.HashPassword(user_json.Password)
	}
	if user_json.Name != "" {
		user.Name = user_json.Name
	}
	if user_json.Email != "" {
		user.Email = user_json.Email
	}
	if user_json.GithubUrl != "" {
		user.GithubUrl = user_json.GithubUrl
	}
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
	user, err := uh.uu.GetByEmail(user_json.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = utils.ComparePassword(user.Password, user_json.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password is wrong"})
		return
	}
	c.JSON(http.StatusOK, user)
}
