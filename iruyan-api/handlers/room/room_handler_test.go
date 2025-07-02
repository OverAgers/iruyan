package room

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"iruyan-api/infrastructure"
	"iruyan-api/repositories"
	usecases "iruyan-api/usecases/room"
)

// テスト用ログ整形
func logStep(t *testing.T, label, msg string) {
	t.Helper()
	t.Logf("[%-10s] %s", label, msg)
}

// テスト用ルーター
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// 本番のDBではなく、テスト用DBまたはモックを使いたい
	testUserRepo := repositories.NewUserRepository(infrastructure.DB)
	testRoomRepo := repositories.NewRoomRepository(infrastructure.DB)
	testSeatRepo := repositories.NewSeatRepository(infrastructure.DB)
	testWorkTimeRepo := repositories.NewWorkTimeRepository(infrastructure.DB)
	testUsecase := usecases.NewRoomUsecase(testUserRepo, testRoomRepo, testSeatRepo, testWorkTimeRepo, infrastructure.DB)
	roomHandler := NewRoomHandler(testUsecase)
	r.POST("/rooms", roomHandler.CreateRoomHandler)
	return r
}

func TestMain(m *testing.M) {
	infrastructure.InitTestDB(nil)
	m.Run()
}

// --- [CreateRoom] 正常系 ---
func TestCreateRoomHandler_Success(t *testing.T) {
	router := setupTestRouter()

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

	logStep(t, "RESULT", "Room created and response returned")
}
