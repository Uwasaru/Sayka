package router

import (
	"github.com/Uwasaru/Sayka/infrastructure/database"
	"github.com/Uwasaru/Sayka/infrastructure/persistence"
	"github.com/Uwasaru/Sayka/presentation/handler"
	"github.com/Uwasaru/Sayka/usecase"
)

func (r Router) InitCommentRouter(conn *database.Conn) {
	cr := persistence.NewCommentRepository(conn)
	ur := persistence.NewUserRepository(conn)
	ar := persistence.NewAuthRepository(conn)
	cu := usecase.NewCommentUsecase(cr, ur)
	ju := usecase.NewJwtUsecase(ar)
	ch := handler.NewCommentHandler(cu, ju)

	g := r.Engine.Group("/comment")
	g.GET("/:id", ch.GetByID)
	g.GET("/user/:userID", ch.GetByUserID)
	g.GET("/sayka/:saykaID", ch.GetBySaykaID)
	g.GET("/", ch.GetAll)
	g.POST("/", ch.CreateComment)
	g.DELETE("/:id", ch.DeleteComment)
}
