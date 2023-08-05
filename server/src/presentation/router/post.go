package router

import (
	"database/sql"

	"github.com/Uwasaru/Sayka/infrastructure/persistence"
	"github.com/Uwasaru/Sayka/presentation/handler"
	"github.com/Uwasaru/Sayka/usecase"
)

func (r *Router) InitPostRouter(db *sql.DB) {
	pr := persistence.NewPostRepository(db)
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
