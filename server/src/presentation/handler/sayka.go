package handler

import (
	"net/http"

	"github.com/Uwasaru/Sayka/presentation/json"
	"github.com/Uwasaru/Sayka/usecase"
	"github.com/Uwasaru/Sayka/utils"
	"github.com/gin-gonic/gin"
)

// Handlerはハンドラーを表します
type SaykaHandler struct {
	pu usecase.ISaykaUsecase
	ju usecase.IJwtUsecase
}

// NewSaykaHandlerは新しいSaykaHandlerを初期化し構造体のポインタを返します
func NewSaykaHandler(pu usecase.ISaykaUsecase, ju usecase.IJwtUsecase) *SaykaHandler {
	return &SaykaHandler{
		pu: pu,
		ju: ju,
	}
}

// GetByIDはIDを指定して投稿を取得します
func (ph *SaykaHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := ph.ju.GetUserIdFromJwtToken(ctx)
	if err != nil {
		userId = ""
	}
	sayka, err := ph.pu.GetByID(ctx, id, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	saykaJson := json.SaykaEntityToJson(sayka)
	ctx.JSON(http.StatusOK, gin.H{"data": saykaJson})
}

// GetByUserIDはUserIDを指定して投稿を取得します
func (ph *SaykaHandler) GetByUserID(ctx *gin.Context) {
	userID := ctx.Param("userID")
	saykas, err := ph.pu.GetByUserID(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, saykas)
}

// GetAllは全ての投稿を取得します
func (ph *SaykaHandler) GetAll(ctx *gin.Context) {
	myId, err := ph.ju.GetUserIdFromJwtToken(ctx)
	if err != nil {
		myId = ""
	}
	saykas, err := ph.pu.GetAll(ctx, myId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, saykas)
}

// GetTimeLineはタイムラインを取得します
func (ph *SaykaHandler) GetTimeLine(ctx *gin.Context) {
	id := ctx.Query("id")
	tag := ctx.Query("tag")
	myId, err := ph.ju.GetUserIdFromJwtToken(ctx)
	if err != nil {
		myId = ""
	}
	saykas, err := ph.pu.GetTimeLine(ctx, id, tag, myId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	saykasJson := json.SaykasEntityToJson(*saykas)
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": saykasJson},
	)
}

func (sh *SaykaHandler) GetAllFavoritedSayka(ctx *gin.Context) {
	myId, err := sh.ju.GetUserIdFromJwtToken(ctx)
	if err != nil {
		myId = ""
	}
	saykas, err := sh.pu.GetAllFavoritedSayka(ctx, myId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	saykasJson := json.SaykasEntityToJson(*saykas)
	ctx.JSON(http.StatusOK, gin.H{"data": saykasJson})
}

// CreateSaykaは投稿を作成します
func (ph *SaykaHandler) CreateSayka(ctx *gin.Context) {
	sayka_json := &json.SaykaJson{}
	if err := ctx.BindJSON(sayka_json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sayka := json.SaykaJsonToEntity(sayka_json)
	sayka.ID = utils.Ulid()
	userID, err := ph.ju.GetUserIdFromJwtToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sayka.UserID = userID
	if err := ph.pu.CreateSayka(ctx, sayka); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	saykaJson := json.SaykaEntityToJson(sayka)
	ctx.JSON(http.StatusOK, gin.H{"data": saykaJson})
}

// UpdateSaykaは投稿を更新します
func (ph *SaykaHandler) UpdateSayka(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := ph.ju.GetUserIdFromJwtToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sayka, err := ph.pu.GetByID(ctx, id, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sayka_json := &json.SaykaJson{}
	if err := ctx.BindJSON(sayka_json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sayka.UserID = sayka_json.UserID
	sayka.AppUrl = sayka_json.AppUrl
	sayka.Description = sayka_json.Description
	sayka.GithubUrl = sayka_json.GithubUrl
	sayka.SlideUrl = sayka_json.SlideUrl
	sayka.Title = sayka_json.Title
	if err := ph.pu.UpdateSayka(ctx, sayka); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, sayka)
}

// DeleteSaykaは投稿を削除します
func (ph *SaykaHandler) DeleteSayka(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := ph.pu.DeleteSayka(ctx, id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
