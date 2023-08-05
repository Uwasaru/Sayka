package handler

import (
	"net/http"

	"github.com/Uwasaru/Sayka/presentation/json"
	"github.com/Uwasaru/Sayka/usecase"
	"github.com/Uwasaru/Sayka/utils"
	"github.com/gin-gonic/gin"
)

// Handlerはハンドラーを表します
type PostHandler struct {
	pu usecase.IPostUsecase
}

// NewPostHandlerは新しいPostHandlerを初期化し構造体のポインタを返します
func NewPostHandler(pu usecase.IPostUsecase) *PostHandler {
	return &PostHandler{
		pu: pu,
	}
}

// GetByIDはIDを指定して投稿を取得します
func (ph *PostHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	post, err := ph.pu.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

// GetByUserIDはUserIDを指定して投稿を取得します
func (ph *PostHandler) GetByUserID(c *gin.Context) {
	userID := c.Param("userID")
	posts, err := ph.pu.GetByUserID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// GetAllは全ての投稿を取得します
func (ph *PostHandler) GetAll(c *gin.Context) {
	posts, err := ph.pu.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// CreatePostは投稿を作成します
func (ph *PostHandler) CreatePost(c *gin.Context) {
	post_json := &json.PostJson{}
	if err := c.BindJSON(post_json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post := json.PostJsonToEntity(post_json)
	post.ID = utils.Ulid()
	post.CreatedAt = utils.GetCurrentTimeStamps()
	if err := ph.pu.CreatePost(post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

// UpdatePostは投稿を更新します
func (ph *PostHandler) UpdatePost(c *gin.Context) {
	id := c.Param("id")
	post, err := ph.pu.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post_json := &json.PostJson{}
	if err := c.BindJSON(post_json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.UserID = post_json.UserID
	post.AppUrl = post_json.AppUrl
	post.Description = post_json.Description
	post.GithubUrl = post_json.GithubUrl
	post.SlideUrl = post_json.SlideUrl
	post.Title = post_json.Title
	if err := ph.pu.UpdatePost(post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

// DeletePostは投稿を削除します
func (ph *PostHandler) DeletePost(c *gin.Context) {
	id := c.Param("id")
	if err := ph.pu.DeletePost(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
