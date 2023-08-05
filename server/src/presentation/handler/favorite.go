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
func (fh *FavoriteHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	favorite, err := fh.fu.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, favorite)
}

// GetByUserIDはUserIDを指定して投稿を取得します
func (fh *FavoriteHandler) GetByUserID(c *gin.Context) {
	userID := c.Param("userID")
	favorites, err := fh.fu.GetByUserID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, favorites)
}

// GetByPostIDはPostIDを指定して投稿を取得します
func (fh *FavoriteHandler) GetByPostID(c *gin.Context) {
	postID := c.Param("postID")
	favorites, err := fh.fu.GetByPostID(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, favorites)
}

// GetAllは全ての投稿を取得します
func (fh *FavoriteHandler) GetAll(c *gin.Context) {
	favorites, err := fh.fu.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, favorites)
}

// CreateFavoriteは投稿を作成します
func (fh *FavoriteHandler) CreateFavorite(c *gin.Context) {
	favorite_json := &json.FavoriteJson{}
	favorite_json.ID = utils.Ulid()
	if err := c.BindJSON(favorite_json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	favorite := json.FavoriteJsonToEntity(favorite_json)
	if err := fh.fu.CreateFavorite(favorite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, favorite)
}

// DeleteFavoriteは投稿を削除します
func (fh *FavoriteHandler) DeleteFavorite(c *gin.Context) {
	id := c.Param("id")
	if err := fh.fu.DeleteFavorite(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
