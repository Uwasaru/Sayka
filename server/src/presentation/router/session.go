package router

import (
	"github.com/Uwasaru/Sayka/infrastructure/database"
	"github.com/Uwasaru/Sayka/infrastructure/persistence"
	"github.com/Uwasaru/Sayka/presentation/handler"
	"github.com/Uwasaru/Sayka/usecase"
)

func (r *Router) InitSessionRouter(conn *database.Conn) {
	sr := persistence.NewSessionRepository(conn)
	su := usecase.NewSessionUsecase(sr)
	sh := handler.NewSessionHandler(su)

	g := r.Engine.Group("/session")
	g.GET("/:id", sh.GetByID)
	g.GET("/user/:userID", sh.GetByUserID)
	g.GET("/", sh.GetAll)
	g.POST("/", sh.CreateSession)
	g.DELETE("/:id", sh.DeleteSession)
}
