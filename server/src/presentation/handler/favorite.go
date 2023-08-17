package handler

import (
	"net/http"

	"github.com/Uwasaru/Sayka/presentation/json"
	"github.com/Uwasaru/Sayka/usecase"
	"github.com/Uwasaru/Sayka/utils"
	"github.com/gin-gonic/gin"
)

// Handlerはハンドラーを表します
type FavoriteHandler struct {
	fu usecase.IFavoriteUsecase
}

// NewFavoriteHandlerは新しいFavoriteHandlerを初期化し構造体のポインタを返します
func NewFavoriteHandler(fu usecase.IFavoriteUsecase) *FavoriteHandler {
	return &FavoriteHandler{
		fu: fu,
	}
}

// GetByIDはIDを指定して投稿を取得します
func (fh *FavoriteHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	favorite, err := fh.fu.GetByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, favorite)
}

// GetByUserIDはUserIDを指定して投稿を取得します
func (fh *FavoriteHandler) GetByUserID(ctx *gin.Context) {
	userID := ctx.Param("userID")
	favorites, err := fh.fu.GetByUserID(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, favorites)
}

// GetBySaykaIDはSaykaIDを指定して投稿を取得します
func (fh *FavoriteHandler) GetBySaykaID(ctx *gin.Context) {
	saykaID := ctx.Param("saykaID")
	favorites, err := fh.fu.GetBySaykaID(ctx, saykaID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, favorites)
}

// GetAllは全ての投稿を取得します
func (fh *FavoriteHandler) GetAll(ctx *gin.Context) {
	favorites, err := fh.fu.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, favorites)
}

// CreateFavoriteは投稿を作成します
func (fh *FavoriteHandler) CreateFavorite(ctx *gin.Context) {
	favorite_json := &json.FavoriteJson{}
	favorite_json.ID = utils.Ulid()
	if err := ctx.BindJSON(favorite_json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	favorite := json.FavoriteJsonToEntity(favorite_json)
	if err := fh.fu.CreateFavorite(ctx, favorite); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, favorite)
}

// DeleteFavoriteは投稿を削除します
func (fh *FavoriteHandler) DeleteFavorite(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := fh.fu.DeleteFavorite(ctx, id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
