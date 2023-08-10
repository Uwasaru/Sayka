package router

import (
	"github.com/Uwasaru/Sayka/infrastructure/database"
	"github.com/Uwasaru/Sayka/infrastructure/persistence"
	"github.com/Uwasaru/Sayka/presentation/handler"
	"github.com/Uwasaru/Sayka/usecase"
)

func (r Router) InitTagRouter(conn *database.Conn) {
	tr := persistence.NewTagRepository(conn)
	tu := usecase.NewTagUsecase(tr)
	th := handler.NewTagHandler(tu)

	g := r.Engine.Group("/tag")
	g.GET("/:id", th.GetByID)
	g.GET("/name/:name", th.GetByName)
	g.GET("/", th.GetAll)
	g.POST("/", th.CreateTag)
}
