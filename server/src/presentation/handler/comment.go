package handler

import (
	"net/http"

	"github.com/Uwasaru/Sayka/presentation/json"
	"github.com/Uwasaru/Sayka/usecase"
	"github.com/Uwasaru/Sayka/utils"
	"github.com/gin-gonic/gin"
)

// Handlerはハンドラーを表します
type CommentHandler struct {
	cu usecase.ICommentUsecase
}

// NewCommentHandlerは新しいCommentHandlerを初期化し構造体のポインタを返します
func NewCommentHandler(cu usecase.ICommentUsecase) *CommentHandler {
	return &CommentHandler{
		cu: cu,
	}
}

// GetByIDはIDを指定して投稿を取得します
func (ch *CommentHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	comment, err := ch.cu.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comment)
}

// GetByUserIDはUserIDを指定して投稿を取得します
func (ch *CommentHandler) GetByUserID(c *gin.Context) {
	userID := c.Param("userID")
	comments, err := ch.cu.GetByUserID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}

// GetByPostIDはPostIDを指定して投稿を取得します
func (ch *CommentHandler) GetByPostID(c *gin.Context) {
	postID := c.Param("postID")
	comments, err := ch.cu.GetByPostID(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}

// GetAllは全ての投稿を取得します
func (ch *CommentHandler) GetAll(c *gin.Context) {
	comments, err := ch.cu.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}

// CreateCommentは投稿を作成します
func (ch *CommentHandler) CreateComment(c *gin.Context) {
	comment_json := &json.CommentJson{}
	comment_json.ID = utils.Ulid()
	comment_json.CreatedAt = utils.GetCurrentTimeStamps()
	if err := c.BindJSON(comment_json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comment := json.CommentJsonToEntity(comment_json)
	if err := ch.cu.CreateComment(comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comment)
}

// DeleteCommentは投稿を削除します
func (ch *CommentHandler) DeleteComment(c *gin.Context) {
	id := c.Param("id")
	if err := ch.cu.DeleteComment(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
