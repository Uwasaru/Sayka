package router

import (
	"github.com/Uwasaru/Sayka/infrastructure/database"
	"github.com/Uwasaru/Sayka/infrastructure/persistence"
	"github.com/Uwasaru/Sayka/presentation/handler"
	"github.com/Uwasaru/Sayka/usecase"
)

func (r Router) InitFavoriteRouter(conn *database.Conn) {
	fr := persistence.NewFavoriteRepository(conn)
	ar := persistence.NewAuthRepository(conn)
	fu := usecase.NewFavoriteUsecase(fr)
	ju := usecase.NewJwtUsecase(ar)
	fh := handler.NewFavoriteHandler(fu, ju)

	g := r.Engine.Group("/favorite")
	g.GET("/:id", fh.GetByID)
	g.GET("/user/:userID", fh.GetByUserID)
	g.GET("/sayka/:saykaID", fh.GetBySaykaID)
	g.GET("/", fh.GetAll)
	// g.POST("/", fh.CreateFavorite)
	g.POST("/", fh.FavoriteAction)
	g.DELETE("/:id", fh.DeleteFavorite)
}
