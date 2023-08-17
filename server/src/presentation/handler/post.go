package handler

import (
	"net/http"

	"github.com/Uwasaru/Sayka/presentation/json"
	"github.com/Uwasaru/Sayka/usecase"
	"github.com/Uwasaru/Sayka/utils"
	"github.com/gin-gonic/gin"
)

// Handlerはハンドラーを表します
type PostHandler struct {
	pu usecase.IPostUsecase
	ju usecase.IJwtUsecase
}

// NewPostHandlerは新しいPostHandlerを初期化し構造体のポインタを返します
func NewPostHandler(pu usecase.IPostUsecase, ju usecase.IJwtUsecase) *PostHandler {
	return &PostHandler{
		pu: pu,
		ju: ju,
	}
}

// GetByIDはIDを指定して投稿を取得します
func (ph *PostHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	post, err := ph.pu.GetByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, post)
}

// GetByUserIDはUserIDを指定して投稿を取得します
func (ph *PostHandler) GetByUserID(ctx *gin.Context) {
	userID := ctx.Param("userID")
	posts, err := ph.pu.GetByUserID(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, posts)
}

// GetAllは全ての投稿を取得します
func (ph *PostHandler) GetAll(ctx *gin.Context) {
	posts, err := ph.pu.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, posts)
}

// GetTimeLineはタイムラインを取得します
func (ph *PostHandler) GetTimeLine(ctx *gin.Context) {
	id := ctx.Query("id")
	tag := ctx.Query("tag")
	posts, err := ph.pu.GetTimeLine(ctx, id, tag)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postsJson := json.PostsEntityToJson(*posts)
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": postsJson},
	)
}

// CreatePostは投稿を作成します
func (ph *PostHandler) CreatePost(ctx *gin.Context) {
	post_json := &json.PostJson{}
	if err := ctx.BindJSON(post_json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post := json.PostJsonToEntity(post_json)
	post.ID = utils.Ulid()
	userID, err := ph.ju.GetUserIdFromJwtToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.UserID = userID
	if err := ph.pu.CreatePost(ctx, post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	postJson := json.PostEntityToJson(post)
	ctx.JSON(http.StatusOK, gin.H{"data": postJson})
}

// UpdatePostは投稿を更新します
func (ph *PostHandler) UpdatePost(ctx *gin.Context) {
	id := ctx.Param("id")
	post, err := ph.pu.GetByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post_json := &json.PostJson{}
	if err := ctx.BindJSON(post_json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.UserID = post_json.UserID
	post.AppUrl = post_json.AppUrl
	post.Description = post_json.Description
	post.GithubUrl = post_json.GithubUrl
	post.SlideUrl = post_json.SlideUrl
	post.Title = post_json.Title
	if err := ph.pu.UpdatePost(ctx, post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, post)
}

// DeletePostは投稿を削除します
func (ph *PostHandler) DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := ph.pu.DeletePost(ctx, id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
