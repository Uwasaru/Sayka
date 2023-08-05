package router

import (
	"database/sql"

	"github.com/Uwasaru/Sayka/infrastructure/persistence"
	"github.com/Uwasaru/Sayka/presentation/handler"
	"github.com/Uwasaru/Sayka/usecase"
)

func (r *Router) InitSessionRouter(db *sql.DB) {
	sr := persistence.NewSessionRepository(db)
	su := usecase.NewSessionUsecase(sr)
	sh := handler.NewSessionHandler(su)

	g := r.Engine.Group("/session")
	g.GET("/:id", sh.GetByID)
	g.GET("/user/:userID", sh.GetByUserID)
	g.GET("/", sh.GetAll)
	g.POST("/", sh.CreateSession)
	g.DELETE("/:id", sh.DeleteSession)
}