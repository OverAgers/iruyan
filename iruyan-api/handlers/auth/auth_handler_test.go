package auth

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"iruyan-api/infrastructure"
	"iruyan-api/repository"
)

// 共通レスポンス構造体
type errorResponse struct {
	Message string `json:"message"`
}

// logStep はテスト内ログを整えるヘルパー（ASCII のみ）
func logStep(t *testing.T, label, msg string) {
	t.Helper()
	t.Logf("[%-10s] %s", label, msg)
}

// テスト用のGinルーターをセットアップ
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/login", LoginHandler)
	return r
}

// TestMain は全テスト実行前の共通初期化処理
func TestMain(m *testing.M) {
	infrastructure.InitTestDB(nil) // テスト用DB初期化
	m.Run()
}

// --- 正常系 ---
func TestLoginHandler_Success(t *testing.T) {
	router := setupTestRouter()
	t.Logf("\n=== [INFO] Start: 正常なログイン成功テスト ===")
	t.Cleanup(func() {
		t.Logf("--- [INFO] End: 正常なログイン成功テスト ---\n")
	})

	err := repository.CreateTestUser("testuser", "pass1234")
	assert.NoError(t, err)

	logStep(t, "REQUEST", "Sending login request with valid credentials")

	data := url.Values{}
	data.Set("iruyanId", "testuser")
	data.Set("password", "pass1234")

	req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"Login successful"`)

	logStep(t, "RESULT", "Received 200 OK and success message")
}

// --- パラメータ不足 ---
func TestLoginHandler_MissingParams(t *testing.T) {
	router := setupTestRouter()
	t.Logf("\n=== [INFO] Start: パラメータ不足テスト ===")
	t.Cleanup(func() {
		t.Logf("--- [INFO] End: パラメータ不足テスト ---\n")
	})

	req, _ := http.NewRequest(http.MethodPost, "/login", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `iruyanId and password are required`)

	logStep(t, "RESULT", "Received 400 Bad Request with expected error message")
}

// --- 存在しないユーザー ---
func TestLoginHandler_UserNotFound(t *testing.T) {
	router := setupTestRouter()
	t.Logf("\n=== [INFO] Start: 存在しないユーザーのテスト ===")
	t.Cleanup(func() {
		t.Logf("--- [INFO] End: 存在しないユーザーのテスト ---\n")
	})

	data := url.Values{}
	data.Set("iruyanId", "nonexistent")
	data.Set("password", "anything")

	req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "user not found")

	logStep(t, "RESULT", "Received 401 Unauthorized for non-existent user")
}

// --- パスワード間違い ---
func TestLoginHandler_WrongPassword(t *testing.T) {
	router := setupTestRouter()
	t.Logf("\n=== [INFO] Start: パスワード不一致テスト ===")
	t.Cleanup(func() {
		t.Logf("--- [INFO] End: パスワード不一致テスト ---\n")
	})

	err := repository.CreateTestUser("testuser", "pass1234")
	assert.NoError(t, err)

	data := url.Values{}
	data.Set("iruyanId", "testuser")
	data.Set("password", "incorrect")

	req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "invalid password")

	logStep(t, "RESULT", "Received 401 Unauthorized for incorrect password")
}
