package handler

import (
	"net/http"

	"github.com/Uwasaru/Sayka/presentation/json"
	"github.com/Uwasaru/Sayka/usecase"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserHandler struct {
	uc usecase.IUserUsecase
}

func NewUserHandler(uc usecase.IUserUsecase) *UserHandler {
	return &UserHandler{
		uc: uc,
	}
}

func (u *UserHandler) CreateUser(ctx *gin.Context) {
	var uj json.UserJson
	if err := ctx.BindJSON(&uj); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	user, err := u.uc.CreateUser(ctx, json.UserJsonToEntity(&uj))
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	userjson := json.UserEntityToJson(user)
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": userjson},
	)
}

func (u *UserHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	i, _ := strconv.Atoi(id)
	err := u.uc.DeleteUser(ctx, i)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{"data": "success"},
	)
}

func (u *UserHandler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	i, _ := strconv.Atoi(id)
	user, err := u.uc.GetUser(ctx, i)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	userjson := json.UserEntityToJson(user)
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": userjson},
	)
}
