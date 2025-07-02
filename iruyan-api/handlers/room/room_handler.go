// Package room provides HTTP handlers for managing rooms,
// including creation, listing, entry/exit operations, and seat assignments.
package room

import (
	"errors"
	errorhandler "iruyan-api/handlers/error"
	"iruyan-api/pkg/errdefs"
	"iruyan-api/presenters"
	"iruyan-api/responses"
	usecase "iruyan-api/usecases/room"

	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	var roomUsecase usecase.RoomUsecase
	room, err := roomUsecase.CreateRoom(roomName)
	if err != nil {
		switch {
		case errors.Is(err, errdefs.ErrInvalidRoomName):
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "invalid room name"})
		case errors.Is(err, errdefs.ErrRoomAlreadyExists):
			c.JSON(http.StatusConflict, responses.ErrorResponse{Message: "room already exists"})
		default:
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "unexpected error: " + err.Error()})
		}
		return
	}

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
	var roomUsecase usecase.RoomUsecase
	rooms, err := roomUsecase.GetRooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "failed to retrieve rooms: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.RoomListResponse{
		Message: "Rooms retrieved successfully",
		Rooms:   rooms,
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
	roomIDStr := c.Param("roomId")

	// UUIDのパース
	roomID, err := uuid.Parse(roomIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "invalid room ID format"})
		return
	}

	var roomUsecase usecase.RoomUsecase
	room, err := roomUsecase.GetRoomByID(roomID)
	if err != nil {
		switch {
		case errors.Is(err, errdefs.ErrRoomNotFound):
			c.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "room not found"})
		default:
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "unexpected error"})
		}
		return
	}

	c.JSON(http.StatusOK, responses.RoomResponse{
		Message: "Room details retrieved successfully",
		RoomID:  room.ID.String(),
		Room:    presenters.ToRoomDetail(room),
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
	roomIDStr := c.Param("roomId")
	iruyanID := c.Param("iruyanId")
	task := c.PostForm("task")

	// UUIDのパース
	roomID, err := uuid.Parse(roomIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "invalid room ID format"})
		return
	}

	var roomUsecase usecase.RoomUsecase
	workTime, room, err := roomUsecase.EnterRoom(iruyanID, roomID, task)
	if err != nil {
		switch {
		case errors.Is(err, errdefs.ErrUserNotFound):
			c.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "user not found"})
		case errors.Is(err, errdefs.ErrRoomNotFound):
			c.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "room not found"})
		case errors.Is(err, errdefs.ErrAlreadyInRoom):
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "user is already in the room"})
		default:
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "failed to record entry to the room"})
		}
		return
	}

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
	durationStr := c.PostForm("duration")

	errorHandler := errorhandler.ErrorHandler{}

	// durationのバリデーション
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		errorHandler.BadRequest(c, "Invalid duration format. Use format like '1h2m3s'.")
		return
	}

	// UUIDチェック
	parsedRoomID, err := uuid.Parse(roomID)
	if err != nil {
		errorHandler.BadRequest(c, "Invalid room ID format")
		return
	}

	var roomUsecase usecase.RoomUsecase
	leaveInfo, err := roomUsecase.LeaveRoom(iruyanID, parsedRoomID, duration)
	if err != nil {
		switch {
		case errors.Is(err, errdefs.ErrUserNotFound):
			errorHandler.NotFoundError(c, "user not found")
		case errors.Is(err, errdefs.ErrRoomNotFound):
			errorHandler.NotFoundError(c, "room not found")
		case errors.Is(err, errdefs.ErrNotInRoom):
			errorHandler.BadRequest(c, "user is not currently in the room")
		case errors.Is(err, errdefs.ErrInvalidDuration):
			errorHandler.BadRequest(c, "invalid duration: shorter than time spent")
		default:
			errorHandler.InternalServerError(c, "Failed to record leaving time")
		}
		return
	}

	c.JSON(http.StatusOK, responses.LeaveRoomResponse{
		Message:     "Left the room successfully",
		RoomID:      leaveInfo.RoomID,
		RoomName:    leaveInfo.RoomName,
		EntryTime:   leaveInfo.EntryTime,
		LeavingTime: leaveInfo.LeavingTime,
		Duration:    leaveInfo.Duration,
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

	roomID, err := uuid.Parse(roomIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "Invalid roomID format"})
		return
	}

	seatNumber, err := strconv.Atoi(seatNumberParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "Invalid seatNumber format"})
		return
	}

	var roomUsecase usecase.RoomUsecase
	if err := roomUsecase.TakeSeat(iruyanID, roomID, seatNumber); err != nil {
		switch {
		case errors.Is(err, errdefs.ErrUserNotFound):
			c.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "user not found"})
		case errors.Is(err, errdefs.ErrSeatNotFound):
			c.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "seat not found in this room"})
		case errors.Is(err, errdefs.ErrSeatAlreadyTaken):
			c.JSON(http.StatusConflict, responses.ErrorResponse{Message: "seat is already taken"})
		case errors.Is(err, errdefs.ErrNotInRoom):
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "user is not currently in the room"})
		default:
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "unexpected error"})
		}
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

	roomID, err := uuid.Parse(roomIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "invalid room ID format"})
		return
	}

	seatNumber, err := strconv.Atoi(seatNumberParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "invalid seat number format"})
		return
	}

	var roomUsecase usecase.RoomUsecase
	if err := roomUsecase.LeaveSeat(iruyanID, roomID, seatNumber); err != nil {
		switch {
		case errors.Is(err, errdefs.ErrUserNotFound):
			c.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "user not found"})
		case errors.Is(err, errdefs.ErrNotSeated):
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "the user is not currently seated"})
		case errors.Is(err, errdefs.ErrSeatMismatch):
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "unexpected error"})
		}
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
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "invalid room ID format"})
		return
	}

	var roomUsecase usecase.RoomUsecase
	seatedUsers, err := roomUsecase.GetSeatedUsers(roomID)
	if err != nil {
		switch {
		case errors.Is(err, errdefs.ErrRoomNotFound):
			c.JSON(http.StatusNotFound, responses.ErrorResponse{Message: "room not found"})
		case errors.Is(err, errdefs.ErrDataRetrievalFailed):
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "failed to retrieve seat status"})
		default:
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "unexpected error"})
		}
		return
	}

	c.JSON(http.StatusOK, seatedUsers)

}
