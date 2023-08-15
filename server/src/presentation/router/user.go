package router

import (
	"os"

	"github.com/Uwasaru/Sayka/infrastructure/database"
	"github.com/Uwasaru/Sayka/infrastructure/github"
	"github.com/Uwasaru/Sayka/infrastructure/persistence"
	"github.com/Uwasaru/Sayka/presentation/handler"
	auth_middleware "github.com/Uwasaru/Sayka/presentation/middleware"
	"github.com/Uwasaru/Sayka/usecase"
)

func (r Router) InitUserRouter(conn *database.Conn) {
	repoauth := persistence.NewAuthRepository(conn)
	repodiscord := github.NewClient(os.Getenv("GITHUB_CALLBACK_API"))
	repouser := persistence.NewUserRepository(conn)

	uc := usecase.NewUserUsecase(repouser)
	ac := usecase.NewAuthUsecase(repoauth, repodiscord, repouser)

	m := auth_middleware.NewAuth(ac)
	h := handler.NewUserHandler(uc)

	//認証middleware
	g := r.Engine.Group("/user")
	g.POST("/", m.Authenticate(), h.CreateUser)
	g.GET("/:id", h.GetUser)
	g.DELETE("/:id", m.Authenticate(), h.DeleteUser)
}
