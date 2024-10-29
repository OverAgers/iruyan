package main

import (
	"iruyan-api/infrastructure"
	"iruyan-api/middleware"
	"iruyan-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// データベース初期化
	infrastructure.InitDB()

	// Ginのルータを作成
	router := gin.Default()

	// ミドルウェアの登録
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.RecoveryMiddleware())

	// ルートにアクセスしたときに "Hello! IRUYAN" を表示
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello! IRUYAN",
		})
	})

	// ルーティングを登録
	routes.RegisterAuthRoutes(router)
	routes.RegisterUserRoutes(router)
	routes.RegisterRoomRoutes(router)

	// サーバー起動
	router.Run(":8080")
}
