package router

import (
	"github.com/Uwasaru/Sayka/infrastructure/database"
	"github.com/Uwasaru/Sayka/infrastructure/persistence"
	"github.com/Uwasaru/Sayka/presentation/handler"
	"github.com/Uwasaru/Sayka/usecase"
)

func (r Router) InitSaykaRouter(conn *database.Conn) {
	pr := persistence.NewSaykaRepository(conn)
	fr := persistence.NewFavoriteRepository(conn)
	cr := persistence.NewCommentRepository(conn)
	ar := persistence.NewAuthRepository(conn)
	ur := persistence.NewUserRepository(conn)
	pu := usecase.NewSaykaUsecase(pr, fr, cr, ur)
	ju := usecase.NewJwtUsecase(ar)
	ph := handler.NewSaykaHandler(pu, ju)

	g := r.Engine.Group("/sayka")
	g.GET("/:id", ph.GetByID)
	g.GET("/user/:userID", ph.GetByUserID)
	g.GET("/", ph.GetAll)
	g.GET("/timeline", ph.GetTimeLine)
	g.GET("/favorite", ph.GetAllFavoritedSayka)
	g.GET("/comment", ph.GetAllCommentedSayka)
	g.POST("/", ph.CreateSayka)
	g.PUT("/:id", ph.UpdateSayka)
	g.DELETE("/:id", ph.DeleteSayka)
}
