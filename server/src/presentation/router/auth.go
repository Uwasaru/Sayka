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

func (r Router) InitAuthRouter(conn *database.Conn) {
	ar := persistence.NewAuthRepository(conn)
	ur := persistence.NewUserRepository(conn)
	gr := github.NewClient(os.Getenv("GITHUB_CALLBACK_API"))

	ac := usecase.NewAuthUsecase(ar, gr, ur)
	uc := usecase.NewUserUsecase(ur)
	h := handler.NewAuthHandler(ac, uc)
	m := auth_middleware.NewAuth(ac)

	g := r.Engine.Group("/auth")
	g.GET("/login", h.Login)
	g.GET("/callback", h.Callback)
	g.GET("/user", m.Authenticate(), h.User)
	g.GET("/cookie/:user_id", h.SetCookieFromUserId)
}
