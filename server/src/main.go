package main

import (
	"github.com/Uwasaru/Sayka/infrastructure/database"
	"github.com/Uwasaru/Sayka/presentation/router"
)

func main() {
	db, err := database.NewDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	r := router.NewRouter()
	r.InitUserRouter(db)
	r.InitPostRouter(db)
	r.InitFavoriteRouter(db)
	r.InitCommentRouter(db)
	r.InitTagRouter(db)
	r.InitSessionRouter(db)
	r.Serve()
}
