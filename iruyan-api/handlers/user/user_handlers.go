package user

import (
	"net/http"
	"iruyan-api/infrastructure"
	"iruyan-api/models"
	"iruyan-api/responses"

	"github.com/gin-gonic/gin"
	"strconv"
)

// ユーザー画面表示
func UserPageHandler(c *gin.Context) {
	userIDParam := c.Param("user_id")
	// user_idをuint型に変換してuserID変数に保存
	userIDUint64, err := strconv.ParseUint(userIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Invalid user id format",
		})
		return 
	}
	userID := uint(userIDUint64)

	var user models.User
	// ユーザIDがDBに存在するか確認する
	// ユーザが見つからない場合はエラーを返す
	if err = user.FindByID(infrastructure.DB, userID); err != nil{
		if err.Error() == "user not found" {
			// ユーザがデータベースに存在しない場合
			c.JSON(http.StatusNotFound, responses.ErrorResponse {
				Message: "User not found",
			})
			return 
		} else {
			// その他のサーバ側のエラーの場合
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse {
				Message: "Database error",
			})
			return  
		}
	}

	// ユーザが存在する場合は200を返す
	c.JSON(http.StatusOK, gin.H{
		"message": "User page accessed successfully",
		"user_id": userID,
	})
}

// ユーザーの削除
func UserDeleteHandler(c *gin.Context) {
	userID := c.Param("user_id")
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
