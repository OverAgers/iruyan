package user

import (
	"iruyan-api/infrastructure"
	"iruyan-api/models"
	"iruyan-api/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ユーザー画面表示
func UserPageHandler(c *gin.Context) {
	userID := c.Param("user_id")
	c.JSON(http.StatusOK, gin.H{
		"message": "User page accessed successfully",
		"user_id": userID,
	})
}

// ユーザーの削除
func UserDeleteHandler(c *gin.Context) {
	// user_idをuint型に変換してuserID変数に保存
	userIDParam := c.Param("user_id")
	userIDUint64, err := strconv.ParseUint(userIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Invalid user ID format",
		})
		return 
	}
	userID := uint(userIDUint64)


	var user models.User

	// ユーザIDがDBに存在するか確認する
	// ユーザIDが存在しない場合は、エラーメッセージを返す
	if err = user.FindByID(infrastructure.DB, userID); err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Message: "User not found",
		})
		return 
	}

	// ユーザIDが存在する場合は、その行を削除する
	if err := infrastructure.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "Failed to delete user",
		})
		return 
	}

	// 削除成功のレスポンスを返す
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
		"user_id": userID,
	})
}

// 1週間の作業日取得
func WorkInfoHandler(c *gin.Context) {
	userID := c.Param("user_id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Work info accessed successfully",
		"user_id": userID,
	})
}

// 一緒に居た時間
func TogetherTimeHandler(c *gin.Context) {
	userID := c.Param("user_id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Together time accessed successfully",
		"user_id": userID,
	})
}

// 集中ランキング取得
func RankingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Ranking page accessed successfully",
	})
}
