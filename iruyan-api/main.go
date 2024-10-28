package main

import (
	"iruyan-api/infrastructure"
	"iruyan-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// データベース初期化
	infrastructure.InitDB()

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

	// サーバー起動
	router.Run(":8080")
}
