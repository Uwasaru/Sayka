package handler

import (
	"net/http"

	"github.com/Uwasaru/Sayka/presentation/json"
	"github.com/Uwasaru/Sayka/usecase"
	"github.com/Uwasaru/Sayka/utils"
	"github.com/gin-gonic/gin"
)

// Handlerはハンドラーを表します
type TagHandler struct {
	tu usecase.ITagUsecase
}

// NewTagHandlerは新しいTagHandlerを初期化し構造体のポインタを返します
func NewTagHandler(tu usecase.ITagUsecase) *TagHandler {
	return &TagHandler{
		tu: tu,
	}
}

// GetByIDはIDを指定して投稿を取得します
func (th *TagHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	tag, err := th.tu.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tag)
}

// GetByNameはNameを指定して投稿を取得します
func (th *TagHandler) GetByName(c *gin.Context) {
	name := c.Param("name")
	tag, err := th.tu.GetByName(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tag)
}

// GetAllは全ての投稿を取得します
func (th *TagHandler) GetAll(c *gin.Context) {
	tags, err := th.tu.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tags)
}

// CreateTagは投稿を作成します
func (th *TagHandler) CreateTag(c *gin.Context) {
	tag_json := &json.TagJson{}
	tag_json.ID = utils.Ulid()
	if err := c.BindJSON(tag_json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tag := json.TagEntityToJson(tag_json)
	if err := th.tu.CreateTag(tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tag)
}