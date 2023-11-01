package main

import (
	"github.com/Uwasaru/Sayka/infrastructure/database"
	"github.com/Uwasaru/Sayka/presentation/router"
	"github.com/gin-gonic/gin"
)

func main() {
	conn, err := database.NewConn()
	if err != nil {
		panic(err)
	}

	r := router.NewRouter()
	// health check
	r.Engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, "ok")
	})
	r.InitUserRouter(conn)
	r.InitSaykaRouter(conn)
	r.InitFavoriteRouter(conn)
	r.InitCommentRouter(conn)
	r.InitTagRouter(conn)
	r.InitAuthRouter(conn)
	r.Serve()
}
