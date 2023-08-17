package main

import (
	"github.com/Uwasaru/Sayka/infrastructure/database"
	"github.com/Uwasaru/Sayka/presentation/router"
)

func main() {
	conn, err := database.NewConn()
	if err != nil {
		panic(err)
	}

	r := router.NewRouter()
	r.InitUserRouter(conn)
	r.InitSaykaRouter(conn)
	r.InitFavoriteRouter(conn)
	r.InitCommentRouter(conn)
	r.InitTagRouter(conn)
	r.InitAuthRouter(conn)
	r.Serve()
}
