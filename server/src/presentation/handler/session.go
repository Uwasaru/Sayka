package handler

import (
	"net/http"

	"github.com/Uwasaru/Sayka/presentation/json"
	"github.com/Uwasaru/Sayka/usecase"
	"github.com/Uwasaru/Sayka/utils"
	"github.com/gin-gonic/gin"
)

// Handlerはハンドラーを表します
type SessionHandler struct {
	su usecase.ISessionUsecase
}

// NewSessionHandlerは新しいSessionHandlerを初期化し構造体のポインタを返します
func NewSessionHandler(su usecase.ISessionUsecase) *SessionHandler {
	return &SessionHandler{
		su: su,
	}
}

// GetByIDはIDを指定してセッションを取得します
func (sh *SessionHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	session, err := sh.su.GetByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, session)
}

// GetByUserIDはUserIDを指定してセッションを取得します
func (sh *SessionHandler) GetByUserID(ctx *gin.Context) {
	userID := ctx.Param("userID")
	session, err := sh.su.GetByUserID(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, session)
}

// GetAllは全てのセッションを取得します
func (sh *SessionHandler) GetAll(ctx *gin.Context) {
	sessions, err := sh.su.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, sessions)
}

// CreateSessionはセッションを作成します
func (sh *SessionHandler) CreateSession(ctx *gin.Context) {
	session_json := &json.SessionJson{}
	if err := ctx.BindJSON(session_json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session_json.ID = utils.Ulid()
	session := json.SessionJsonToEntity(session_json)
	if err := sh.su.CreateSession(ctx, session); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, session)
}

// UpdateSessionはセッションを更新します
func (sh *SessionHandler) UpdateSession(ctx *gin.Context) {
	id := ctx.Param("id")
	session, err := sh.su.GetByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session_json := &json.SessionJson{}
	if err := ctx.BindJSON(session_json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session_json.ID = session.ID
	session = json.SessionJsonToEntity(session_json)
	if err := sh.su.UpdateSession(ctx, session); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, session)
}

// DeleteSessionはセッションを削除します
func (sh *SessionHandler) DeleteSession(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := sh.su.DeleteSession(ctx, id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}