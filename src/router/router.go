package router

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/mr-platform-digital-twin-server/controller"
	"github.com/kajiLabTeam/mr-platform-digital-twin-server/util"
)

func Init() {
	f, _ := os.Create("../log/server.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	util.InitContentFileBucket()

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!!")
	})
	// ユーザースペース
	r.POST("/api/space/user/content/get", func(c *gin.Context) {
		controller.UpdateSpace(c)
	})
	// パブリックスペース
	r.POST("/api/space/public/create", func(c *gin.Context) {
		controller.CreatePublicSpace(c)
	})
	r.POST("/api/space/public/content/create", func(c *gin.Context) {
		controller.CreatePublicSpaceContent(c)
	})
	r.POST("/api/space/public/content/update", func(c *gin.Context) {
		controller.UpdatePublicSpaceContent(c)
	})

	if err := r.Run("0.0.0.0:8000"); err != nil {
		fmt.Println("サーバーの起動に失敗しました:", err)
	} else {
		fmt.Println("サーバーが正常に起動しました")
	}
}
