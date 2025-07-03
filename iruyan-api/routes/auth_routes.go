package routes

import (
	handlers "iruyan-api/handlers/auth"
	"iruyan-api/infrastructure"
	"iruyan-api/repositories"
	usecases "iruyan-api/usecases/auth"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	// --- 依存性の注入 ---
	db := infrastructure.DB // *gorm.DB のインスタンス
	userRepo := repositories.NewUserRepository(db)
	authUsecase := usecases.NewAuthUsecase(userRepo)
	authHandler := handlers.NewAuthHandler(authUsecase)

	// ログイン画面表示
	router.GET("/login", authHandler.LoginPageHandler)

	// ログイン処理
	router.POST("/login", authHandler.LoginHandler)

	// 新規登録画面表示
	router.GET("/register", authHandler.RegisterPageHandler)

	// 新規登録処理
	router.POST("/register", authHandler.RegisterHandler)

	// ログアウト
	router.POST("/logout", authHandler.LogoutHandler)
}
