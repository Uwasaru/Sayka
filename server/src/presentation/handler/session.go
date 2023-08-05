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
func (sh *SessionHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	session, err := sh.su.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, session)
}

// GetByUserIDはUserIDを指定してセッションを取得します
func (sh *SessionHandler) GetByUserID(c *gin.Context) {
	userID := c.Param("userID")
	session, err := sh.su.GetByUserID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, session)
}

// GetAllは全てのセッションを取得します
func (sh *SessionHandler) GetAll(c *gin.Context) {
	sessions, err := sh.su.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sessions)
}

// CreateSessionはセッションを作成します
func (sh *SessionHandler) CreateSession(c *gin.Context) {
	session_json := &json.SessionJson{}
	if err := c.BindJSON(session_json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session_json.ID = utils.Ulid()
	session := json.SessionJsonToEntity(session_json)
	if err := sh.su.CreateSession(session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, session)
}

// UpdateSessionはセッションを更新します
func (sh *SessionHandler) UpdateSession(c *gin.Context) {
	id := c.Param("id")
	session, err := sh.su.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session_json := &json.SessionJson{}
	if err := c.BindJSON(session_json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session_json.ID = session.ID
	session = json.SessionJsonToEntity(session_json)
	if err := sh.su.UpdateSession(session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, session)
}

// DeleteSessionはセッションを削除します
func (sh *SessionHandler) DeleteSession(c *gin.Context) {
	id := c.Param("id")
	if err := sh.su.DeleteSession(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}