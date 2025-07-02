package room

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"iruyan-api/infrastructure"
	"iruyan-api/models"
	"iruyan-api/pkg/errdefs"
	"iruyan-api/responses"
	"iruyan-api/usecases/room/mocks"
)

// テスト用ログ整形
func logStep(t *testing.T, label, msg string) {
	t.Helper()
	t.Logf("[%-10s] %s", label, msg)
}

// テスト用ルーター
func setupTestRouterWithMock(mockRoomUsecase *mocks.RoomUsecase) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	roomHandler := NewRoomHandler(mockRoomUsecase)
	r.GET("/rooms", roomHandler.GetRoomsHandler)
	r.GET("//rooms/:roomId", roomHandler.GetRoomHandler)
	r.POST("/rooms", roomHandler.CreateRoomHandler)
	r.POST("/rooms/:roomId/users/:iruyanId/enter", roomHandler.EnterRoomHandler)
	r.POST("/rooms/:roomId/leave", roomHandler.LeaveRoomHandler)
	r.PUT("/rooms/:roomId/seats/:seatNumber/take", roomHandler.TakeSeatHandler)
	r.PUT("/rooms/:roomId/seats/:seatNumber/leave", roomHandler.LeaveSeatHandler)
	r.GET("/rooms/:roomId/seats/status", roomHandler.GetSeatedUsersInRoomHandler)

	return r
}

func TestMain(m *testing.M) {
	infrastructure.InitTestDB(nil)
	m.Run()
}

// --- [CreateRoom] 正常系 ---
func TestCreateRoomHandler_Success(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)

	// 🔧 ここが重要！"roomtest" に対して成功を返すモックを定義
	mockUsecase.On("CreateRoom", "roomtest").Return(&models.Room{
		ID:   uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name: "roomtest",
	}, nil)

	router := setupTestRouterWithMock(mockUsecase)

	t.Logf("\n=== [INFO] Start: 正常なルーム作成 ===")
	t.Cleanup(func() {
		t.Logf("--- [INFO] End ---\n")
	})

	data := url.Values{}
	data.Set("roomName", "roomtest")

	req, _ := http.NewRequest(http.MethodPost, "/rooms", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"Room created successfully"`)

	mockUsecase.AssertExpectations(t)
}

func TestCreateRoomHandler_InvalidRoomName(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)

	// "invalid" という引数に対して ErrInvalidRoomName を返すモック
	mockUsecase.On("CreateRoom", "invalid").Return(nil, errdefs.ErrInvalidRoomName)

	router := setupTestRouterWithMock(mockUsecase)

	// --- テストリクエストの準備 ---
	data := url.Values{}
	data.Set("roomName", "invalid")

	req, err := http.NewRequest(http.MethodPost, "/rooms", strings.NewReader(data.Encode()))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// --- 実行 ---
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	// --- アサーション ---
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `"invalid room name"`)

	// モックが呼ばれたことを検証（ミス防止に有効）
	mockUsecase.AssertExpectations(t)
}

func TestCreateRoomHandler_CreateRoomFailed(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)

	// モックの期待値設定
	mockUsecase.On("CreateRoom", "roomtest").Return(nil, errdefs.ErrCreateRoomFailed)

	router := setupTestRouterWithMock(mockUsecase)

	// --- リクエスト準備 ---
	data := url.Values{}
	data.Set("roomName", "roomtest")

	req, err := http.NewRequest(http.MethodPost, "/rooms", strings.NewReader(data.Encode()))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// --- 実行 ---
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	// --- 検証 ---
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), `"failed to create room"`)

	// モックの呼び出しを検証
	mockUsecase.AssertExpectations(t)
}

func TestCreateRoomHandler_CreateSeatFailed(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)

	// モック設定：roomName が "roomtest" のとき CreateSeat エラーを返す
	mockUsecase.On("CreateRoom", "roomtest").Return(nil, errdefs.ErrCreateSeatFailed)

	router := setupTestRouterWithMock(mockUsecase)

	// --- リクエスト作成 ---
	data := url.Values{}
	data.Set("roomName", "roomtest")

	req, err := http.NewRequest(http.MethodPost, "/rooms", strings.NewReader(data.Encode()))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// --- 実行 ---
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	// --- アサーション ---
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), `"failed to create seats"`)

	// --- モック呼び出しの検証 ---
	mockUsecase.AssertExpectations(t)
}

func TestCreateRoomHandler_TransactionCommitFailed(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)

	// トランザクションコミット失敗のモック
	mockUsecase.On("CreateRoom", "roomtest").Return(nil, errdefs.ErrTransactionCommit)

	router := setupTestRouterWithMock(mockUsecase)

	// --- リクエスト準備 ---
	data := url.Values{}
	data.Set("roomName", "roomtest")

	req, err := http.NewRequest(http.MethodPost, "/rooms", strings.NewReader(data.Encode()))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// --- 実行 ---
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	// --- 検証 ---
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), `"failed to finalize room creation"`)

	// モック呼び出しの検証
	mockUsecase.AssertExpectations(t)
}

func TestCreateRoomHandler_UnexpectedError(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)

	// "roomtest" に対して予期しないエラーを返すモック
	mockUsecase.On("CreateRoom", "roomtest").Return(nil, errors.New("unexpected"))

	router := setupTestRouterWithMock(mockUsecase)

	// --- リクエスト準備 ---
	data := url.Values{}
	data.Set("roomName", "roomtest")

	req, err := http.NewRequest(http.MethodPost, "/rooms", strings.NewReader(data.Encode()))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// --- 実行 ---
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	// --- アサーション ---
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), `"unexpected error"`)

	// モック呼び出しの検証
	mockUsecase.AssertExpectations(t)
}

func TestGetRoomsHandler_Success(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)

	mockUsecase.On("GetRooms").Return([]responses.RoomDetail{
		{
			RoomID:   "11111111-1111-1111-1111-111111111111",
			RoomName: "RoomA",
			Seats: []responses.SeatDetail{
				{
					SeatID:     "aaaaaaa1-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
					RoomID:     "11111111-1111-1111-1111-111111111111",
					SeatNumber: 1,
				},
				{
					SeatID:     "aaaaaaa2-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
					RoomID:     "11111111-1111-1111-1111-111111111111",
					SeatNumber: 2,
				},
			},
		},
	}, nil)

	router := setupTestRouterWithMock(mockUsecase)

	req, err := http.NewRequest(http.MethodGet, "/rooms", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusOK, w.Code)

	var response responses.RoomListResponse
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Rooms retrieved successfully", response.Message)
	assert.Len(t, response.Rooms, 1)

	room := response.Rooms[0]
	assert.Equal(t, "11111111-1111-1111-1111-111111111111", room.RoomID)
	assert.Equal(t, "RoomA", room.RoomName)
	assert.Len(t, room.Seats, 2)
	assert.Equal(t, 1, room.Seats[0].SeatNumber)
	assert.Equal(t, 2, room.Seats[1].SeatNumber)

	mockUsecase.AssertExpectations(t)
}

func TestGetRoomsHandler_Failure(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)

	mockUsecase.On("GetRooms").Return(nil, errors.New("database error"))

	router := setupTestRouterWithMock(mockUsecase)

	req, err := http.NewRequest(http.MethodGet, "/rooms", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var res responses.ErrorResponse
	err = json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.Contains(t, res.Message, "failed to retrieve rooms: database error")

	mockUsecase.AssertExpectations(t)
}

func TestEnterRoomHandler_Success(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)

	roomID := uuid.New()
	userID := "test-user"
	task := "reading"
	entryTime := time.Now()

	mockUsecase.On("EnterRoom", userID, roomID, task).Return(&models.WorkTime{
		RoomID:    roomID,
		EntryTime: entryTime,
	}, &models.Room{
		ID:   roomID,
		Name: "RoomA",
	}, nil)

	router := setupTestRouterWithMock(mockUsecase)

	reqURL := fmt.Sprintf("/rooms/%s/users/%s/enter", roomID, userID)
	data := url.Values{} // OK!

	data.Set("task", task)

	req, err := http.NewRequest(http.MethodPost, reqURL, strings.NewReader(data.Encode()))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusOK, w.Code)

	var res responses.EnterRoomResponse
	err = json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)

	assert.Equal(t, "Room entry recorded successfully", res.Message)
	assert.Equal(t, roomID.String(), res.RoomID)
	assert.Equal(t, "RoomA", res.RoomName)
	assert.Equal(t, task, res.Task)

	mockUsecase.AssertExpectations(t)
}

func TestEnterRoomHandler_InvalidUUID(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	router := setupTestRouterWithMock(mockUsecase)

	req, _ := http.NewRequest(http.MethodPost, "/rooms/invalid-uuid/users/test-user/enter", nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `"invalid room ID format"`)
}

func TestEnterRoomHandler_UserNotFound(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)

	roomID := uuid.New()
	mockUsecase.On("EnterRoom", "notfound", roomID, "").Return(nil, nil, errdefs.ErrUserNotFound)

	router := setupTestRouterWithMock(mockUsecase)

	url := fmt.Sprintf("/rooms/%s/users/notfound/enter", roomID.String())
	req, _ := http.NewRequest(http.MethodPost, url, nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), `"user not found"`)
}

func TestEnterRoomHandler_RoomNotFound(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)

	roomID := uuid.New()
	mockUsecase.On("EnterRoom", "user1", roomID, "").Return(nil, nil, errdefs.ErrRoomNotFound)

	router := setupTestRouterWithMock(mockUsecase)

	url := fmt.Sprintf("/rooms/%s/users/user1/enter", roomID.String())
	req, _ := http.NewRequest(http.MethodPost, url, nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), `"room not found"`)
}

func TestEnterRoomHandler_AlreadyInRoom(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)

	roomID := uuid.New()
	mockUsecase.On("EnterRoom", "user1", roomID, "").Return(nil, nil, errdefs.ErrAlreadyInRoom)

	router := setupTestRouterWithMock(mockUsecase)

	url := fmt.Sprintf("/rooms/%s/users/user1/enter", roomID.String())
	req, _ := http.NewRequest(http.MethodPost, url, nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `"user is already in the room"`)
}

func TestEnterRoomHandler_UnexpectedError(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)

	roomID := uuid.New()
	mockUsecase.On("EnterRoom", "user1", roomID, "").Return(nil, nil, errors.New("unexpected error"))

	router := setupTestRouterWithMock(mockUsecase)

	url := fmt.Sprintf("/rooms/%s/users/user1/enter", roomID.String())
	req, _ := http.NewRequest(http.MethodPost, url, nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), `"failed to record entry to the room"`)
}

func TestLeaveRoomHandler_Success(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)

	roomID := uuid.New()
	iruyanID := "user123"
	duration := 2 * time.Hour
	entryTime := time.Now().Add(-duration)
	leavingTime := time.Now()

	mockUsecase.On("LeaveRoom", iruyanID, roomID, duration).Return(&responses.LeaveRoomResponse{
		RoomID:      roomID.String(),
		RoomName:    "RoomA",
		EntryTime:   entryTime,
		LeavingTime: leavingTime,
		Duration:    duration,
	}, nil)

	router := setupTestRouterWithMock(mockUsecase)

	form := url.Values{}
	form.Set("iruyanId", iruyanID)
	form.Set("duration", duration.String())

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/rooms/%s/leave", roomID), strings.NewReader(form.Encode()))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusOK, w.Code)

	var res responses.LeaveRoomResponse
	err = json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)

	assert.Equal(t, "Left the room successfully", res.Message)
	assert.Equal(t, roomID.String(), res.RoomID)
	assert.Equal(t, "RoomA", res.RoomName)

	mockUsecase.AssertExpectations(t)
}

func TestLeaveRoomHandler_InvalidDurationFormat(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	roomID := uuid.New()

	form := url.Values{}
	form.Set("iruyanId", "user123")
	form.Set("duration", "invalid-duration")

	router := setupTestRouterWithMock(mockUsecase)

	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/rooms/%s/leave", roomID), strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid duration format")
}

func TestLeaveRoomHandler_InvalidRoomIDFormat(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)

	form := url.Values{}
	form.Set("iruyanId", "user123")
	form.Set("duration", "1h")

	router := setupTestRouterWithMock(mockUsecase)

	req, _ := http.NewRequest(http.MethodPost, "/rooms/invalid-uuid/leave", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid room ID format")
}

func TestLeaveRoomHandler_UserNotFound(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	roomID := uuid.New()

	mockUsecase.On("LeaveRoom", "missing-user", roomID, time.Hour).Return(nil, errdefs.ErrUserNotFound)

	router := setupTestRouterWithMock(mockUsecase)

	form := url.Values{}
	form.Set("iruyanId", "missing-user")
	form.Set("duration", "1h")

	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/rooms/%s/leave", roomID), strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "user not found")
}

func TestLeaveRoomHandler_RoomNotFound(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	roomID := uuid.New()

	mockUsecase.On("LeaveRoom", "user123", roomID, time.Hour).Return(nil, errdefs.ErrRoomNotFound)

	router := setupTestRouterWithMock(mockUsecase)

	form := url.Values{}
	form.Set("iruyanId", "user123")
	form.Set("duration", "1h")

	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/rooms/%s/leave", roomID), strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "room not found")
}

func TestLeaveRoomHandler_NotInRoom(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	roomID := uuid.New()

	mockUsecase.On("LeaveRoom", "user123", roomID, time.Hour).Return(nil, errdefs.ErrNotInRoom)

	router := setupTestRouterWithMock(mockUsecase)

	form := url.Values{}
	form.Set("iruyanId", "user123")
	form.Set("duration", "1h")

	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/rooms/%s/leave", roomID), strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "user is not currently in the room")
}

func TestLeaveRoomHandler_InvalidDuration(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	roomID := uuid.New()

	mockUsecase.On("LeaveRoom", "user123", roomID, time.Second).Return(nil, errdefs.ErrInvalidDuration)

	router := setupTestRouterWithMock(mockUsecase)

	form := url.Values{}
	form.Set("iruyanId", "user123")
	form.Set("duration", "1s")

	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/rooms/%s/leave", roomID), strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid duration")
}

func TestLeaveRoomHandler_UnexpectedError(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	roomID := uuid.New()

	mockUsecase.On("LeaveRoom", "user123", roomID, time.Hour).Return(nil, errors.New("DB failure"))

	router := setupTestRouterWithMock(mockUsecase)

	form := url.Values{}
	form.Set("iruyanId", "user123")
	form.Set("duration", "1h")

	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/rooms/%s/leave", roomID), strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Failed to record leaving time")
}

func TestTakeSeatHandler_Success(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	roomID := uuid.New()
	seatNumber := 5
	iruyanID := "user123"

	mockUsecase.
		On("TakeSeat", iruyanID, roomID, seatNumber).
		Return(nil)

	router := setupTestRouterWithMock(mockUsecase)

	form := url.Values{}
	form.Set("iruyanId", iruyanID)

	req, err := http.NewRequest(http.MethodPut,
		fmt.Sprintf("/rooms/%s/seats/%d/take", roomID, seatNumber),
		strings.NewReader(form.Encode()))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"Seated successfully"`)

	mockUsecase.AssertExpectations(t)
}

func TestTakeSeatHandler_InvalidRoomID(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)

	router := setupTestRouterWithMock(mockUsecase)

	form := url.Values{}
	form.Set("iruyanId", "user123")

	req, _ := http.NewRequest(http.MethodPut,
		"/rooms/invalid-uuid/seats/1/take",
		strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid roomID format")
}

func TestTakeSeatHandler_InvalidSeatNumber(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	roomID := uuid.New()

	router := setupTestRouterWithMock(mockUsecase)

	form := url.Values{}
	form.Set("iruyanId", "user123")

	req, _ := http.NewRequest(http.MethodPut,
		fmt.Sprintf("/rooms/%s/seats/abc/take", roomID),
		strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid seatNumber format")
}

func TestTakeSeatHandler_UserNotFound(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	roomID := uuid.New()

	mockUsecase.On("TakeSeat", "user123", roomID, 1).Return(errdefs.ErrUserNotFound)

	router := setupTestRouterWithMock(mockUsecase)

	form := url.Values{}
	form.Set("iruyanId", "user123")

	req, _ := http.NewRequest(http.MethodPut,
		fmt.Sprintf("/rooms/%s/seats/1/take", roomID),
		strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "user not found")
}

func TestTakeSeatHandler_SeatNotFound(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	roomID := uuid.New()

	mockUsecase.On("TakeSeat", "user123", roomID, 99).Return(errdefs.ErrSeatNotFound)

	router := setupTestRouterWithMock(mockUsecase)

	form := url.Values{}
	form.Set("iruyanId", "user123")

	req, _ := http.NewRequest(http.MethodPut,
		fmt.Sprintf("/rooms/%s/seats/99/take", roomID),
		strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "seat not found")
}

func TestTakeSeatHandler_SeatAlreadyTaken(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	roomID := uuid.New()

	mockUsecase.On("TakeSeat", "user123", roomID, 1).Return(errdefs.ErrSeatAlreadyTaken)

	router := setupTestRouterWithMock(mockUsecase)

	form := url.Values{}
	form.Set("iruyanId", "user123")

	req, _ := http.NewRequest(http.MethodPut,
		fmt.Sprintf("/rooms/%s/seats/1/take", roomID),
		strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Contains(t, w.Body.String(), "seat is already taken")
}

func TestTakeSeatHandler_NotInRoom(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	roomID := uuid.New()

	mockUsecase.On("TakeSeat", "user123", roomID, 1).Return(errdefs.ErrNotInRoom)

	router := setupTestRouterWithMock(mockUsecase)

	form := url.Values{}
	form.Set("iruyanId", "user123")

	req, _ := http.NewRequest(http.MethodPut,
		fmt.Sprintf("/rooms/%s/seats/1/take", roomID),
		strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "user is not currently in the room")
}

func TestTakeSeatHandler_UnexpectedError(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	roomID := uuid.New()

	mockUsecase.On("TakeSeat", "user123", roomID, 1).Return(errors.New("db failure"))

	router := setupTestRouterWithMock(mockUsecase)

	form := url.Values{}
	form.Set("iruyanId", "user123")

	req, _ := http.NewRequest(http.MethodPut,
		fmt.Sprintf("/rooms/%s/seats/1/take", roomID),
		strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "unexpected error")
}

func TestLeaveSeatHandler_Success(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	roomID := uuid.New()
	seatNumber := 1
	iruyanID := "user123"

	mockUsecase.
		On("LeaveSeat", iruyanID, roomID, seatNumber).
		Return(nil)

	router := setupTestRouterWithMock(mockUsecase)

	form := url.Values{}
	form.Set("iruyanId", iruyanID)

	req, err := http.NewRequest(http.MethodPut,
		fmt.Sprintf("/rooms/%s/seats/%d/leave", roomID, seatNumber),
		strings.NewReader(form.Encode()))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"Left the seat successfully"`)

	mockUsecase.AssertExpectations(t)
}

func TestLeaveSeatHandler_InvalidRoomID(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)

	router := setupTestRouterWithMock(mockUsecase)

	form := url.Values{}
	form.Set("iruyanId", "user123")

	req, _ := http.NewRequest(http.MethodPut,
		"/rooms/invalid-uuid/seats/1/leave",
		strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid room ID format")
}

func TestLeaveSeatHandler_InvalidSeatNumber(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	roomID := uuid.New()

	router := setupTestRouterWithMock(mockUsecase)

	form := url.Values{}
	form.Set("iruyanId", "user123")

	req, _ := http.NewRequest(http.MethodPut,
		fmt.Sprintf("/rooms/%s/seats/abc/leave", roomID),
		strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid seat number format")
}

func TestLeaveSeatHandler_UserNotFound(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	roomID := uuid.New()

	mockUsecase.
		On("LeaveSeat", "missing-user", roomID, 1).
		Return(errdefs.ErrUserNotFound)

	router := setupTestRouterWithMock(mockUsecase)

	form := url.Values{}
	form.Set("iruyanId", "missing-user")

	req, _ := http.NewRequest(http.MethodPut,
		fmt.Sprintf("/rooms/%s/seats/1/leave", roomID),
		strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "user not found")
}

func TestLeaveSeatHandler_NotSeated(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	roomID := uuid.New()

	mockUsecase.
		On("LeaveSeat", "user123", roomID, 1).
		Return(errdefs.ErrNotSeated)

	router := setupTestRouterWithMock(mockUsecase)

	form := url.Values{}
	form.Set("iruyanId", "user123")

	req, _ := http.NewRequest(http.MethodPut,
		fmt.Sprintf("/rooms/%s/seats/1/leave", roomID),
		strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "the user is not currently seated")
}

func TestLeaveSeatHandler_SeatMismatch(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	roomID := uuid.New()

	mockUsecase.
		On("LeaveSeat", "user123", roomID, 1).
		Return(errdefs.ErrSeatMismatch)

	router := setupTestRouterWithMock(mockUsecase)

	form := url.Values{}
	form.Set("iruyanId", "user123")

	req, _ := http.NewRequest(http.MethodPut,
		fmt.Sprintf("/rooms/%s/seats/1/leave", roomID),
		strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "the seat number does not match the user's current seat")
}

func TestLeaveSeatHandler_UnexpectedError(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	roomID := uuid.New()

	mockUsecase.
		On("LeaveSeat", "user123", roomID, 1).
		Return(errors.New("db failure"))

	router := setupTestRouterWithMock(mockUsecase)

	form := url.Values{}
	form.Set("iruyanId", "user123")

	req, _ := http.NewRequest(http.MethodPut,
		fmt.Sprintf("/rooms/%s/seats/1/leave", roomID),
		strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "unexpected error")
}

func TestGetSeatedUsersInRoomHandler_Success(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	roomID := uuid.New()

	expected := []responses.SeatStatusResponse{
		{
			SeatNumber: 1,
			IruyanID:   "user123",
			UserName:   "User One",
			Email:      "user1@example.com",
			AvatarUrl:  "https://example.com/avatar1.png",
			Task:       "Coding",
			Note:       "Important task",
			Status:     "seated",
			StartTime:  1720000000,
		},
	}

	mockUsecase.On("GetSeatedUsers", roomID).Return(expected, nil)

	router := setupTestRouterWithMock(mockUsecase)
	req, _ := http.NewRequest(http.MethodGet, "/rooms/"+roomID.String()+"/seats/status", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	t.Logf("Status: %d", w.Code)
	t.Logf("Body: %s", w.Body.String())

	assert.Equal(t, http.StatusOK, w.Code)

	var actual []responses.SeatStatusResponse
	err := json.Unmarshal(w.Body.Bytes(), &actual)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)

	mockUsecase.AssertExpectations(t)
}

func TestGetSeatedUsersInRoomHandler_InvalidUUID(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)

	router := setupTestRouterWithMock(mockUsecase)
	req, _ := http.NewRequest(http.MethodGet, "/rooms/invalid-uuid/seats/status", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid room ID format")
}

func TestGetSeatedUsersInRoomHandler_RoomNotFound(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	roomID := uuid.New()

	mockUsecase.On("GetSeatedUsers", roomID).Return(nil, errdefs.ErrRoomNotFound)

	router := setupTestRouterWithMock(mockUsecase)
	req, _ := http.NewRequest(http.MethodGet, "/rooms/"+roomID.String()+"/seats/status", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "room not found")
	mockUsecase.AssertExpectations(t)
}

func TestGetSeatedUsersInRoomHandler_DataRetrievalFailed(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	roomID := uuid.New()

	mockUsecase.On("GetSeatedUsers", roomID).Return(nil, errdefs.ErrDataRetrievalFailed)

	router := setupTestRouterWithMock(mockUsecase)
	req, _ := http.NewRequest(http.MethodGet, "/rooms/"+roomID.String()+"/seats/status", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "failed to retrieve seat status")
	mockUsecase.AssertExpectations(t)
}

func TestGetRoomHandler_Success(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	roomID := uuid.New()

	mockRoom := &models.Room{
		ID:   roomID,
		Name: "RoomA",
		Seats: []models.Seat{
			{ID: uuid.New(), RoomID: roomID, SeatNumber: 1},
			{ID: uuid.New(), RoomID: roomID, SeatNumber: 2},
		},
	}

	mockUsecase.On("GetRoomByID", roomID).Return(mockRoom, nil)

	router := setupTestRouterWithMock(mockUsecase)
	req, _ := http.NewRequest(http.MethodGet, "/rooms/"+roomID.String(), nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	t.Logf("Status: %d", w.Code)
	t.Logf("Body: %s", w.Body.String())

	assert.Equal(t, http.StatusOK, w.Code)

	var res responses.RoomResponse
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.Equal(t, "Room details retrieved successfully", res.Message)
	assert.Equal(t, roomID.String(), res.RoomID)
	assert.Equal(t, "RoomA", res.Room.RoomName)
	assert.Len(t, res.Room.Seats, 2)

	mockUsecase.AssertExpectations(t)
}

func TestGetRoomHandler_InvalidUUID(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)

	router := setupTestRouterWithMock(mockUsecase)
	req, _ := http.NewRequest(http.MethodGet, "/rooms/invalid-uuid", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid room ID format")
}

func TestGetRoomHandler_RoomNotFound(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	roomID := uuid.New()

	mockUsecase.On("GetRoomByID", roomID).Return(nil, errdefs.ErrRoomNotFound)

	router := setupTestRouterWithMock(mockUsecase)
	req, _ := http.NewRequest(http.MethodGet, "/rooms/"+roomID.String(), nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "room not found")

	mockUsecase.AssertExpectations(t)
}

func TestGetRoomHandler_UnexpectedError(t *testing.T) {
	mockUsecase := new(mocks.RoomUsecase)
	roomID := uuid.New()

	mockUsecase.On("GetRoomByID", roomID).Return(nil, errors.New("DB failure"))

	router := setupTestRouterWithMock(mockUsecase)
	req, _ := http.NewRequest(http.MethodGet, "/rooms/"+roomID.String(), nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "unexpected error")

	mockUsecase.AssertExpectations(t)
}
