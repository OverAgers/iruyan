package main

import (
	_ "iruyan-api/docs" // Swaggerのドキュメントをインポート
	"iruyan-api/infrastructure"
	"iruyan-api/middleware"
	"iruyan-api/repositories"
	"iruyan-api/routes"
	usecase "iruyan-api/usecases/room"
	"log"
	"os"

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

	db := infrastructure.DB

	// Repositoryの初期化
	userRepo := repositories.NewUserRepository(db)
	roomRepo := repositories.NewRoomRepository(db)
	seatRepo := repositories.NewSeatRepository(db)
	workTimeRepo := repositories.NewWorkTimeRepository(db)

	// Usecaseの初期化
	roomUsecase := usecase.NewRoomUsecase(userRepo, roomRepo, seatRepo, workTimeRepo, db)

	// 部屋の初期化処理
	roomName := os.Getenv("DEFAULT_ROOM_NAME")
	if roomName == "" {
		roomName = "General"
	}
	if err := roomUsecase.CreateRoomWithSeats(roomName, 10); err != nil {
		log.Printf("⚠️  初期ルーム作成失敗: %v", err)
	}
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
	routes.RegisterWorktimeRoutes(router)

	// サーバー起動
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
