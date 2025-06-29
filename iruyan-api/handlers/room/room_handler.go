// Package room provides HTTP handlers for managing rooms,
// including creation, listing, entry/exit operations, and seat assignments.
package room

import (
	"fmt"
	"iruyan-api/handlers/seat"
	"iruyan-api/infrastructure"
	"iruyan-api/models"
	"iruyan-api/responses"

	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateRoomHandler Room作成
// @Summary Create a new room
// @Description Creates a room with the provided name
// @Tags room
// @Accept x-www-form-urlencoded
// @Produce json
// @Param roomName formData string true "Room Name" default(RoomA)
// @Success 200 {object} responses.RoomCreateResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /rooms [post]
func CreateRoomHandler(c *gin.Context) {
	roomName := c.PostForm("roomName")

	room := models.Room{Name: roomName}

	if err := room.ValidateName(); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: err.Error()})
		return
	}

	// トランザクションを開始
	tx := infrastructure.DB.Begin()
	if err := tx.Create(&room).Error; err != nil {
		tx.Rollback() // エラーが発生したらロールバック
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "Failed to create room: " + err.Error(),
		})
		return
	}

	// 部屋の作成と同時にシートを自動的に10個作る
	for i := 0; i < 10; i++ {
		_, err := seat.CreateSeat(tx, room.ID)
		if err != nil {
			tx.Rollback() // シート作成に失敗したらロールバック
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Message: "Failed to create seat: " + err.Error(),
			})
			return
		}
	}

	tx.Commit() // 全て（ルームの作成・シートの作成）成功したらコミット

	c.JSON(http.StatusOK, responses.RoomCreateResponse{
		Message:  "Room created successfully",
		RoomID:   room.ID,
		RoomName: room.Name,
	})

}

// GetRoomsHandler ルーム一覧を取得
// @Summary Get list of rooms with seats
// @Description Retrieves a list of rooms, each with associated seat information
// @Tags rooms
// @Produce json
// @Success 200 {object} responses.RoomListResponse
// @Router /rooms [get]
func GetRoomsHandler(c *gin.Context) {
	var rooms []models.Room

	// ルーム情報を関連するシート情報と共に取得
	if err := infrastructure.DB.Preload("Seats").Find(&rooms).Error; err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "Failed to retrieve rooms: " + err.Error(),
		})
		return
	}

	// レスポンスデータの準備
	roomDetails := make([]responses.RoomDetail, len(rooms))
	for i, room := range rooms {
		seats := make([]responses.SeatDetail, len(room.Seats))
		for j, seat := range room.Seats {
			seats[j] = responses.SeatDetail{
				SeatID:     seat.ID.String(),
				RoomID:     seat.RoomID.String(),
				SeatNumber: seat.SeatNumber,
			}
		}
		roomDetails[i] = responses.RoomDetail{
			RoomID:   room.ID.String(),
			RoomName: room.Name,
			Seats:    seats,
		}
	}

	// レスポンスを送信
	c.JSON(http.StatusOK, responses.RoomListResponse{
		Message: "Rooms retrieved successfully",
		Rooms:   roomDetails,
	})
}

// GetRoomHandler Room表示
// @Summary Get room details
// @Description Retrieves the details of a specific room
// @Tags room
// @Produce json
// @Param roomId path string true "Room ID"
// @Success 200 {object} responses.RoomDetailResponse
// @Failure 404 {object} responses.ErrorResponse
// @Router /rooms/{roomId} [get]
func GetRoomHandler(c *gin.Context) {
	roomID := c.Param("roomId")

	// Roomモデルを定義
	var room models.Room

	// 指定された room_id の Roomをデータベースから取得
	if err := infrastructure.DB.Preload("Seats").Where("id = ?", roomID).First(&room).Error; err != nil {
		// Roomが見つからない場合、404エラーレスポンスを返す
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Message: "Room not found",
		})
		return
	}

	// Room情報とSeat情報をレスポンス形式に整形
	seats := make([]responses.SeatDetail, len(room.Seats))
	for i, seat := range room.Seats {
		seats[i] = responses.SeatDetail{
			SeatID:     seat.ID.String(),
			RoomID:     seat.RoomID.String(),
			SeatNumber: seat.SeatNumber,
		}
	}

	// Roomが見つかった場合の成功レスポンス
	c.JSON(http.StatusOK, responses.RoomResponse{
		Message: "Room details retrieved successfully",
		RoomID:  room.ID.String(),
		Room: responses.RoomDetail{
			RoomID:   room.ID.String(),
			RoomName: room.Name,
			Seats:    seats,
		},
	})
}

// EnterRoomHandler 部屋への入室処理
// @Summary Enter a room
// @Description Records the entry time of a user entering a specific room and stores the entry information in WorkTime.
// @Tags room
// @Accept json
// @Produce json
// @Param roomId path string true "Room ID"
// @Param iruyanId path string true "Iruyan ID" default(johndoe)
// @Success 200 {object} responses.RoomActionResponse "Entered the room successfully"
// @Failure 404 {object} responses.ErrorResponse "Room not found"
// @Failure 500 {object} responses.ErrorResponse "Failed to record entry to the room"
// @Router /rooms/{roomId}/enter/{iruyanId} [post]
func EnterRoomHandler(c *gin.Context) {
	roomID := c.Param("roomId")
	iruyanID := c.Param("iruyanId")
	task := c.PostForm("task")

	// ユーザーをデータベースから取得
	var user models.User
	if err := user.FindByIruyanID(infrastructure.DB, iruyanID); err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Message: "User not found",
		})
		return
	}

	var room models.Room
	// ルームが存在するか確認
	if err := room.FindByID(infrastructure.DB, roomID); err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Message: "Room not found",
		})
		return
	}

	parsedRoomID, err := uuid.Parse(roomID)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "invalid room ID format"})
		return
	}

	// WorkTimeに入室情報を記録
	workTime, err := models.RecordEntry(infrastructure.DB, iruyanID, parsedRoomID, task)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, responses.ErrorResponse{
				Message: "User not found",
			})
		} else if err.Error() == "user is already in the room" {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Message: "User is already in the room",
			})
		} else {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Message: "Failed to record entry to the room",
			})
		}
		return
	}

	// 成功時のレスポンスとしてWorkTime情報を返す
	c.JSON(http.StatusOK, responses.EnterRoomResponse{
		Message:   "Room entry recorded successfully",
		RoomID:    workTime.RoomID.String(),
		RoomName:  room.Name,
		EntryTime: workTime.EntryTime,
		Task:      task,
	})
}

// LeaveRoomHandler Room退室処理
// @Summary Leave a room
// @Description Allows a user to leave a specific room and records the leaving time
// @Tags room
// @Produce json
// @Param roomId path string true "Room ID"
// @Param iruyanId formData string true "Iruyan ID" default(johndoe)
// @Param duration formData string true "Duration (e.g., 1h2m3s)" default(1h2m3s)
// @Success 200 {object} responses.LeaveRoomResponseSwagger "Left the room successfully"
// @Failure 400 {object} responses.ErrorResponse "Invalid duration format or other validation errors"
// @Failure 404 {object} responses.ErrorResponse "Room or user not found"
// @Failure 500 {object} responses.ErrorResponse "Failed to record leaving time"
// @Router /rooms/{roomId}/leave [post]
func LeaveRoomHandler(c *gin.Context) {
	roomID := c.Param("roomId")
	iruyanID := c.PostForm("iruyanId")
	durationStr := c.PostForm("duration") // DurationをPostFormで取得

	// Durationをtime.Duration型に変換
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Invalid duration format. Use format like '1h2m3s'.",
		})
		return
	}

	// ルームが存在するか確認
	var room models.Room
	if err := room.FindByID(infrastructure.DB, roomID); err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Message: "Room not found",
		})
		return
	}

	// ユーザーをデータベースから取得
	var user models.User
	if err := user.FindByIruyanID(infrastructure.DB, iruyanID); err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Message: "User not found",
		})
		return
	}

	parsedRoomID, err := uuid.Parse(roomID)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "invalid room ID format"})
		return
	}

	// WorkTimeテーブルから最新の入室記録を取得（LeavingTimeがNULLのレコードを取得）
	workTime, err := models.GetLatestEntry(infrastructure.DB, iruyanID, parsedRoomID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Message: "User is not currently in the room",
			})
		} else {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Message: "Failed to find entry record",
			})
		}
		return
	}

	// LeavingTimeとDurationを設定
	leavingTime := time.Now()
	workTime.LeavingTime = leavingTime
	workTime.Duration = duration // クライアントから受け取ったDurationを使用

	// Durationがエントリー時間から計算されるDurationより短ければエラーを返す
	if duration < leavingTime.Sub(workTime.EntryTime) {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Invalid duration: Duration is shorter than time spent in room",
		})
		return
	}

	// 退室記録をデータベースに保存
	if err := infrastructure.DB.Save(workTime).Error; err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "Failed to record leaving time",
		})
		return
	}

	// 成功時のレスポンス
	c.JSON(http.StatusOK, responses.LeaveRoomResponse{
		Message:     "Left the room successfully",
		RoomID:      roomID,
		RoomName:    room.Name,
		EntryTime:   workTime.EntryTime,
		LeavingTime: leavingTime,
		Duration:    workTime.Duration,
	})
}

// TakeSeatHandler godoc
// @Summary ユーザーが座席に着席する
// @Description 指定された部屋・座席番号にユーザーを着席させます。
// @Tags seat
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param roomId path string true "Room UUID"
// @Param seatNumber path int true "Seat Number"
// @Param iruyanId formData string true "Iruyan ID（ユーザー識別子）"
// @Success 200 {object} map[string]interface{} "着席成功時のレスポンス"
// @Failure 400 {object} responses.ErrorResponse "リクエスト形式や座席不整合時"
// @Failure 404 {object} responses.ErrorResponse "ユーザーまたは座席が存在しない場合"
// @Failure 409 {object} responses.ErrorResponse "座席がすでに他ユーザーに使用されている場合"
// @Failure 500 {object} responses.ErrorResponse "サーバ内部エラー"
// @Router /rooms/{roomId}/seat/{seatNumber}/take [put]
func TakeSeatHandler(c *gin.Context) {
	roomIDParam := c.Param("roomId")
	seatNumberParam := c.Param("seatNumber")
	iruyanID := c.PostForm("iruyanId")

	// room_idをUUID型に変換してroomID変数に保存
	roomID, err := uuid.Parse(roomIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Invalid roomID format",
		})
		return
	}

	// seat_idをUUID型に変換してseatID変数に保存
	seatNumber, err := strconv.Atoi(seatNumberParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Invalid seatID format",
		})
		return
	}

	// ユーザーが存在するか確認
	var user models.User
	if err := user.FindByIruyanID(infrastructure.DB, iruyanID); err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Message: "User not found",
		})
		return
	}

	// 指定された座席が部屋に存在するかを確認
	// 存在しない場合、エラー（無効な座席番号）を返す
	var seat models.Seat
	if err := seat.IsSeatExistsInRoom(infrastructure.DB, roomID, seatNumber); err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Message: "Seat is not found in this room",
		})
		return
	}

	// 指定された座席が空いているかを確認
	// 空いていない場合、エラーを返す
	var work models.WorkTime
	if err := work.IsSeatTakenInRoom(infrastructure.DB, roomID, seatNumber); err != nil {
		c.JSON(http.StatusConflict, responses.ErrorResponse{
			Message: "Seat is already taken",
		})
		return
	}

	// WorkTimeテーブルからユーザの最新の入室記録を取得
	workTime, err := models.GetLatestEntry(infrastructure.DB, iruyanID, roomID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Message: "User is not currently in the room",
			})
		} else {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Message: "Failed to find entry record",
			})
		}
		return
	}

	// 座席情報を更新する
	workTime.SeatNumber = seatNumber

	// 座席情報をデータベースに更新
	if err := infrastructure.DB.Model(&workTime).Update("seat_number", seatNumber).Error; err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "Failed to record seat number",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Seated successfully",
		"roomId":  roomID,
		"seatId":  seatNumber,
	})
}

// LeaveSeatHandler godoc
// @Summary ユーザーが座席から離席する
// @Description ユーザーが現在着席中の座席から離れます。
// @Tags seat
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param roomId path string true "Room UUID"
// @Param seatNumber path int true "Seat Number"
// @Param iruyanId formData string true "Iruyan ID（ユーザー識別子）"
// @Success 200 {object} map[string]interface{} "離席成功時のレスポンス"
// @Failure 400 {object} responses.ErrorResponse "ユーザーが着席していない、または指定された座席と一致しない場合"
// @Failure 404 {object} responses.ErrorResponse "ユーザーが存在しない場合"
// @Failure 500 {object} responses.ErrorResponse "サーバ内部エラー"
// @Router /rooms/{roomId}/seat/{seatNumber}/leave [put]
func LeaveSeatHandler(c *gin.Context) {
	roomIDParam := c.Param("roomId")
	seatNumberParam := c.Param("seatNumber")
	iruyanID := c.PostForm("iruyanId")

	// room_idをUUID型に変換してroomID変数に保存
	roomID, err := uuid.Parse(roomIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Invalid roomID format",
		})
		return
	}

	// seat_idをUUID型に変換してseatID変数に保存
	seatNumber, err := strconv.Atoi(seatNumberParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Invalid seatID format",
		})
		return
	}

	// ユーザーが存在するか確認
	var user models.User
	if err := user.FindByIruyanID(infrastructure.DB, iruyanID); err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Message: "User not found",
		})
		return
	}

	// WorkTimeテーブルからユーザの最新の入室記録を取得
	workTime, err := models.GetLatestEntry(infrastructure.DB, iruyanID, roomID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{
				Message: "User is not currently in the room",
			})
		} else {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
				Message: "Failed to find entry record",
			})
		}
		return
	}

	// ユーザのDB上の座席情報が指定された座席情報と一致しているか確認する
	// 着席していなかった場合
	if workTime.SeatNumber == 0 {
		// 座席番号が0であれば、ユーザーは着席していない
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "The user is not currently seated.",
		})
		return
	}

	// 着席しているが、番号が一致していなかった場合
	if workTime.SeatNumber != seatNumber {
		// 座席番号が一致しない場合のエラーメッセージ
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: fmt.Sprintf("The seat number does not match the user's record. Current seat number is %d", workTime.SeatNumber),
		})
		return
	}

	// 座席情報を更新する
	newSeatNumber := 0
	workTime.SeatNumber = seatNumber

	// 座席情報をデータベースに更新
	if err := infrastructure.DB.Model(&workTime).Update("seat_number", newSeatNumber).Error; err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "Failed to record seat number",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Left the seat successfully",
		"roomId":     roomID,
		"seatNumber": seatNumber,
	})
}

// GetSeatedUsersInRoomHandler godoc
// @Summary Get seated users in a specific room
// @Description Returns a list of users currently seated in the given room
// @Tags room
// @Produce json
// @Param roomId path string true "Room ID"
// @Success 200 {array} responses.SeatStatusResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /rooms/{roomId}/seats/status [get]
func GetSeatedUsersInRoomHandler(c *gin.Context) {
	roomIDParam := c.Param("roomId")

	roomID, err := uuid.Parse(roomIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Invalid roomID format",
		})
		return
	}

	var workTimes []models.WorkTime
	if err := infrastructure.DB.
		Preload("User").
		Where("room_id = ? AND seat_number > 0", roomID).
		Find(&workTimes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "Failed to retrieve seat status",
		})
		return
	}

	// 誰も座っていない場合は空配列を返す
	if len(workTimes) == 0 {
		c.JSON(http.StatusOK, []responses.SeatStatusResponse{})
		return
	}

	// 整形して返す
	response := make([]responses.SeatStatusResponse, 0, len(workTimes))
	for _, wt := range workTimes {
		response = append(response, responses.SeatStatusResponse{
			SeatNumber: wt.SeatNumber,
			IruyanID:   wt.User.IruyanID,
			UserName:   wt.User.Name,
			Email:      wt.User.Email,
			AvatarUrl:  "",                  // GORMでUserにAvatarUrlフィールドがある場合
			Task:       wt.Task,             // WorkTimeにTaskがある想定
			Note:       "",                  // WorkTimeにNoteがある想定
			Status:     "",                  // enumなら文字列に変換
			StartTime:  wt.EntryTime.Unix(), // time.TimeならUnix秒に変換
		})
	}

	c.JSON(http.StatusOK, response)
}
