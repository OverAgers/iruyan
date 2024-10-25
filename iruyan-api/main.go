package main

import (
	"iruyan-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Ginのルータを作成
	router := gin.Default()

	// ルートにアクセスしたときに "Hello! IRUYAN" を表示
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello! IRUYAN",
		})
	})

	// ユーザー関連のルーティングを登録
	routes.RegisterUserRoutes(router)

	// ルーム関連のルーティングは今後追加
	// routes.RegisterRoomRoutes(router)

	// サーバー起動
	router.Run(":8080")
}
