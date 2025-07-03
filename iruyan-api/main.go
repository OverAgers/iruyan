package main

import (
	"iruyan-api/config"
	_ "iruyan-api/docs" // Swaggerのドキュメントをインポート
	"iruyan-api/infrastructure"
	"iruyan-api/middleware"
	"iruyan-api/repositories"
	"iruyan-api/routes"
	usecase "iruyan-api/usecases/room"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title IRUYAN API
// @version 1.0
// @description This is a server for IRUYAN.
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @type http
// @scheme bearer
// @bearerFormat JWT
// @in header
// @name Authorization
// @description JWT形式のアクセストークン。「Bearer <token>」の形式で入力してください
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

func main() {
	// 環境変数の初期化
	config.Init()

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
	defaultRoomName := config.DefaultRoomName
	if defaultRoomName == "" {
		defaultRoomName = "General"
	}
	if err := roomUsecase.CreateRoomWithSeats(defaultRoomName, 10); err != nil {
		log.Printf("⚠️ 初期ルーム作成失敗: %v", err)
	}
	// Ginのルータを作成
	router := gin.Default()

	// ミドルウェアの登録
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.RecoveryMiddleware())

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ルートにアクセスしたときに "Hello! IRUYAN" を表示
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello! IRUYAN",
		})
	})

	// === APIグループ化: /api/v1 ===
	apiV1 := router.Group("/api/v1")

	// 認証ミドルウェアを必要とするAPIグループ（必要に応じて）
	authRequired := apiV1.Group("")
	authRequired.Use(middleware.JWTMiddleware())

	// エンドポイント登録（必要に応じて authRequired or apiV1 に振り分け）
	routes.RegisterAuthRoutes(apiV1)        // 例: /api/v1/login, /api/v1/register
	routes.RegisterUserRoutes(authRequired) // 認証必要: /api/v1/user
	routes.RegisterRoomRoutes(authRequired)
	routes.RegisterSeatRoutes(authRequired)
	routes.RegisterWorktimeRoutes(authRequired)

	// サーバー起動
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
