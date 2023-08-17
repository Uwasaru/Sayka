package router

import (
	"github.com/Uwasaru/Sayka/infrastructure/database"
	"github.com/Uwasaru/Sayka/infrastructure/persistence"
	"github.com/Uwasaru/Sayka/presentation/handler"
	"github.com/Uwasaru/Sayka/usecase"
)

func (r Router) InitPostRouter(conn *database.Conn) {
	pr := persistence.NewPostRepository(conn)
	fr := persistence.NewFavoriteRepository(conn)
	cr := persistence.NewCommentRepository(conn)
	ar := persistence.NewAuthRepository(conn)
	pu := usecase.NewPostUsecase(pr, fr, cr)
	ju := usecase.NewJwtUsecase(ar)
	ph := handler.NewPostHandler(pu, ju)

	g := r.Engine.Group("/post")
	g.GET("/:id", ph.GetByID)
	g.GET("/user/:userID", ph.GetByUserID)
	g.GET("/", ph.GetAll)
	g.GET("/timeline", ph.GetTimeLine)
	g.POST("/", ph.CreatePost)
	g.PUT("/:id", ph.UpdatePost)
	g.DELETE("/:id", ph.DeletePost)
}
