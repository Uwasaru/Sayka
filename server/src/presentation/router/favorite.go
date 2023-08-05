package router

import (
	"database/sql"

	"github.com/Uwasaru/Sayka/infrastructure/persistence"
	"github.com/Uwasaru/Sayka/presentation/handler"
	"github.com/Uwasaru/Sayka/usecase"
)

func (r *Router) InitFavoriteRouter(db *sql.DB) {
	fr := persistence.NewFavoriteRepository(db)
	fu := usecase.NewFavoriteUsecase(fr)
	fh := handler.NewFavoriteHandler(fu)

	g := r.Engine.Group("/favorite")
	g.GET("/:id", fh.GetByID)
	g.GET("/user/:userID", fh.GetByUserID)
	g.GET("/post/:postID", fh.GetByPostID)
	g.GET("/", fh.GetAll)
	g.POST("/", fh.CreateFavorite)
	g.DELETE("/:id", fh.DeleteFavorite)
}
