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
	ju usecase.IJwtUsecase
}

// NewFavoriteHandlerは新しいFavoriteHandlerを初期化し構造体のポインタを返します
func NewFavoriteHandler(fu usecase.IFavoriteUsecase, ju usecase.IJwtUsecase) *FavoriteHandler {
	return &FavoriteHandler{
		fu: fu,
		ju: ju,
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

func (fh *FavoriteHandler) FavoriteAction(ctx *gin.Context) {
	favorite_json := &json.FavoriteJson{}
	if err := ctx.BindJSON(favorite_json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, err := fh.ju.GetUserIdFromJwtToken(ctx)
	favorite_json.UserID = userId
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if favorite_json.IsFavorite {
		if err := fh.fu.DeleteFavoriteBySaykaIDUserID(ctx, favorite_json.SaykaID, favorite_json.UserID); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		favorite_json.ID = utils.Ulid()
		favorite := json.FavoriteJsonToEntity(favorite_json)
		if err := fh.fu.CreateFavorite(ctx, favorite); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	favoriteCount, err := fh.fu.GetCountBySaykaID(ctx, favorite_json.SaykaID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := &json.FavoriterResponseJson{}
	res.IsFavorite = !favorite_json.IsFavorite
	res.FavoriteCount = favoriteCount
	ctx.JSON(http.StatusOK, gin.H{"data": res})
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
