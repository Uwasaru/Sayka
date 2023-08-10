package router

import (
	"github.com/Uwasaru/Sayka/infrastructure/database"
	"github.com/Uwasaru/Sayka/infrastructure/persistence"
	"github.com/Uwasaru/Sayka/presentation/handler"
	"github.com/Uwasaru/Sayka/usecase"
)

func (r Router) InitUserRouter(conn *database.Conn) {
	ur := persistence.NewUserRepository(conn)
	uu := usecase.NewUserUsecase(ur)
	uh := handler.NewUserHandler(uu)

	g := r.Engine.Group("/user")
	g.GET("/:id", uh.GetByID)
	g.GET("/github_id/:githubId", uh.GetByGithubId)
	g.POST("/", uh.CreateUser)
	g.POST("/login", uh.LoginUser)
	g.PUT("/:id", uh.UpdateUser)
	g.DELETE("/:id", uh.DeleteUser)
}
