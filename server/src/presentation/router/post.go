package router

import (
	"github.com/Uwasaru/Sayka/infrastructure/database"
	"github.com/Uwasaru/Sayka/infrastructure/persistence"
	"github.com/Uwasaru/Sayka/presentation/handler"
	"github.com/Uwasaru/Sayka/usecase"
)

func (r Router) InitPostRouter(conn *database.Conn) {
	pr := persistence.NewPostRepository(conn)
	pu := usecase.NewPostUsecase(pr)
	ph := handler.NewPostHandler(pu)

	g := r.Engine.Group("/post")
	g.GET("/:id", ph.GetByID)
	g.GET("/user/:userID", ph.GetByUserID)
	g.GET("/", ph.GetAll)
	g.POST("/", ph.CreatePost)
	g.PUT("/:id", ph.UpdatePost)
	g.DELETE("/:id", ph.DeletePost)
}
