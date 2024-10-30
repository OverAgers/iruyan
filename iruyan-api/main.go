package main

import (
	_ "iruyan-api/docs" // Swaggerのドキュメントをインポート
	"iruyan-api/infrastructure"
	"iruyan-api/middleware"
	"iruyan-api/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title IRUYAN API
// @version 1.0
// @description This is a server for IRUYAN.
// @host localhost:8080
// @BasePath /

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

func main() {
	// データベース初期化
	infrastructure.InitDB()

	// Ginのルータを作成
	router := gin.Default()

	// ミドルウェアの登録
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.RecoveryMiddleware())

	// Swaggerのエンドポイント
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
	routes.RegisterSeatRoutes(router)

	// サーバー起動
	router.Run(":8080")
}
