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
func (th *TagHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	tag, err := th.tu.GetByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tag)
}

// GetByNameはNameを指定して投稿を取得します
func (th *TagHandler) GetByName(ctx *gin.Context) {
	name := ctx.Param("name")
	tag, err := th.tu.GetByName(ctx, name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tag)
}

// GetAllは全ての投稿を取得します
func (th *TagHandler) GetAll(ctx *gin.Context) {
	tags, err := th.tu.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tags)
}

// CreateTagは投稿を作成します
func (th *TagHandler) CreateTag(ctx *gin.Context) {
	tag_json := &json.TagJson{}
	tag_json.ID = utils.Ulid()
	if err := ctx.BindJSON(tag_json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tag := json.TagEntityToJson(tag_json)
	if err := th.tu.CreateTag(ctx, tag); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tag)
}
