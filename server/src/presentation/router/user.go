package router

import (
	"database/sql"

	"github.com/Uwasaru/Sayka/infrastructure/persistence"
	"github.com/Uwasaru/Sayka/presentation/handler"
	"github.com/Uwasaru/Sayka/usecase"
)

func (r *Router) InitUserRouter(db *sql.DB) {
	ur := persistence.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)
	uh := handler.NewUserHandler(uu)

	g := r.Engine.Group("/user")
	g.GET("/:id", uh.GetByID)
	g.GET("/user/:userID", uh.GetByGithubId)
	g.POST("/", uh.CreateUser)
	g.POST("/login", uh.LoginUser)
	g.PUT("/:id", uh.UpdateUser)
	g.DELETE("/:id", uh.DeleteUser)
}
