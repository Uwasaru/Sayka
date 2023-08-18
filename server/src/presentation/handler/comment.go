package handler

import (
	"fmt"
	"net/http"

	"github.com/Uwasaru/Sayka/presentation/json"
	"github.com/Uwasaru/Sayka/usecase"
	"github.com/Uwasaru/Sayka/utils"
	"github.com/gin-gonic/gin"
)

// Handlerはハンドラーを表します
type CommentHandler struct {
	cu usecase.ICommentUsecase
	ju usecase.IJwtUsecase
}

// NewCommentHandlerは新しいCommentHandlerを初期化し構造体のポインタを返します
func NewCommentHandler(cu usecase.ICommentUsecase, ju usecase.IJwtUsecase) *CommentHandler {
	return &CommentHandler{
		cu: cu,
		ju: ju,
	}
}

// GetByIDはIDを指定して投稿を取得します
func (ch *CommentHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	comment, err := ch.cu.GetByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	commentJson := json.CommentEntityToJson(comment)
	ctx.JSON(http.StatusOK, gin.H{"data": commentJson})
}

// GetByUserIDはUserIDを指定して投稿を取得します
func (ch *CommentHandler) GetByUserID(ctx *gin.Context) {
	userID := ctx.Param("userID")
	comments, err := ch.cu.GetByUserID(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	commentsJson := json.CommentsEntityToJson(*comments)
	ctx.JSON(http.StatusOK, gin.H{"data": commentsJson})
}

// GetBySaykaIDはSaykaIDを指定して投稿を取得します
func (ch *CommentHandler) GetBySaykaID(ctx *gin.Context) {
	saykaID := ctx.Param("saykaID")
	comments, err := ch.cu.GetBySaykaID(ctx, saykaID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	commentsJson := json.CommentsEntityToJson(*comments)
	ctx.JSON(http.StatusOK, gin.H{"data": commentsJson})
}

// GetAllは全ての投稿を取得します
func (ch *CommentHandler) GetAll(ctx *gin.Context) {
	comments, err := ch.cu.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	commentsJson := json.CommentsEntityToJson(*comments)
	ctx.JSON(http.StatusOK, gin.H{"data": commentsJson})
}

// CreateCommentは投稿を作成します
func (ch *CommentHandler) CreateComment(ctx *gin.Context) {
	comment_json := &json.CommentJson{}
	comment_json.ID = utils.Ulid()
	if err := ctx.BindJSON(comment_json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, err := ch.ju.GetUserIdFromJwtToken(ctx)
	fmt.Println(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comment_json.UserID = userId
	comment := json.CommentJsonToEntity(comment_json)
	if err := ch.cu.CreateComment(ctx, comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	commentJson := json.CommentEntityToJson(comment)
	ctx.JSON(http.StatusOK, gin.H{"data": commentJson})
}

// DeleteCommentは投稿を削除します
func (ch *CommentHandler) DeleteComment(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := ch.cu.DeleteComment(ctx, id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
