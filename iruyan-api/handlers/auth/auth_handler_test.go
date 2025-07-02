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
	"iruyan-api/repositories"
	usecases "iruyan-api/usecases/auth"
)

// logStep はテスト内ログを整えるヘルパー（ASCII のみ）
func logStep(t *testing.T, label, msg string) {
	t.Helper()
	t.Logf("[%-10s] %s", label, msg)
}

// テスト用のGinルーターをセットアップ
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// 本番のDBではなく、テスト用DBまたはモックを使いたい
	testRepo := repositories.NewUserRepository(infrastructure.DB) // ここをMockにしてもOK
	testUsecase := usecases.NewAuthUsecase(testRepo)
	authHandler := NewAuthHandler(testUsecase)

	// ルーティングには構造体のメソッドを渡す
	r.POST("/login", authHandler.LoginHandler)
	r.POST("/register", authHandler.RegisterHandler)

	return r
}

// TestMain は全テスト実行前の共通初期化処理
func TestMain(m *testing.M) {
	infrastructure.InitTestDB(nil) // テスト用DB初期化
	m.Run()
}

// --- [Login] 正常系 ---
func TestLoginHandler_Success(t *testing.T) {
	router := setupTestRouter()
	t.Logf("\n=== [INFO] Start: 正常なログイン成功テスト ===")
	t.Cleanup(func() {
		t.Logf("--- [INFO] End: 正常なログイン成功テスト ---\n")
	})

	err := repositories.CreateTestUser("testuser", "pass1234")
	assert.NoError(t, err)

	logStep(t, "REQUEST", "Sending login request with valid credentials")

	data := url.Values{}
	data.Set("iruyanId", "testuser")
	data.Set("password", "pass1234")

	req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"Login successful"`)

	logStep(t, "RESULT", "Received 200 OK and success message")
}

// --- [Login] パラメータ不足 ---
func TestLoginHandler_MissingParams(t *testing.T) {
	router := setupTestRouter()
	t.Logf("\n=== [INFO] Start: パラメータ不足テスト ===")
	t.Cleanup(func() {
		t.Logf("--- [INFO] End: パラメータ不足テスト ---\n")
	})

	req, _ := http.NewRequest(http.MethodPost, "/login", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `iruyanId and password are required`)

	logStep(t, "RESULT", "Received 400 Bad Request with expected error message")
}

// --- [Login] 存在しないユーザー ---
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

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "user not found")

	logStep(t, "RESULT", "Received 401 Unauthorized for non-existent user")
}

// --- [Login] パスワード間違い ---
func TestLoginHandler_WrongPassword(t *testing.T) {
	router := setupTestRouter()
	t.Logf("\n=== [INFO] Start: パスワード不一致テスト ===")
	t.Cleanup(func() {
		t.Logf("--- [INFO] End: パスワード不一致テスト ---\n")
	})

	err := repositories.CreateTestUser("testuser", "pass1234")
	assert.NoError(t, err)

	data := url.Values{}
	data.Set("iruyanId", "testuser")
	data.Set("password", "incorrect")

	req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "invalid password")

	logStep(t, "RESULT", "Received 401 Unauthorized for incorrect password")
}

// --- [Register] 登録成功 ---
func TestRegisterHandler_Success(t *testing.T) {
	router := setupTestRouter()
	t.Logf("\n=== [INFO] Start: 登録成功テスト ===")
	t.Cleanup(func() {
		t.Logf("--- [INFO] End: 登録成功テスト ---\n")
	})

	data := url.Values{}
	data.Set("iruyanId", "newuser")
	data.Set("password", "securepass1234")
	data.Set("userName", "New User")
	data.Set("email", "newuser@example.com")

	req, _ := http.NewRequest(http.MethodPost, "/register", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"Registration successful"`)

	logStep(t, "RESULT", "Received 200 OK with success message")
}

// --- [Register] パラメータ不足 ---
func TestRegisterHandler_MissingParams(t *testing.T) {
	router := setupTestRouter()
	t.Logf("\n=== [INFO] Start: パラメータ不足による登録失敗テスト ===")
	t.Cleanup(func() {
		t.Logf("--- [INFO] End: パラメータ不足による登録失敗テスト ---\n")
	})

	req, _ := http.NewRequest(http.MethodPost, "/register", nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Missing required parameter(s)") // 具体的なバリデーションエラーに応じて修正

	logStep(t, "RESULT", "Received 400 Bad Request due to missing params")
}

// --- [Register] 重複登録（iruyanIDのユニーク制約エラー） ---
func TestRegisterHandler_DuplicateUser(t *testing.T) {
	router := setupTestRouter()
	t.Logf("\n=== [INFO] Start: 重複登録エラーテスト ===")
	t.Cleanup(func() {
		t.Logf("--- [INFO] End: 重複登録エラーテスト ---\n")
	})

	// 最初の登録
	_ = repositories.CreateTestUser("duplicateuser", "samepass1234")

	data := url.Values{}
	data.Set("iruyanId", "duplicateuser")
	data.Set("password", "samepass1234")
	data.Set("userName", "Duplicate User")
	data.Set("email", "duplicate@example.com")

	req, _ := http.NewRequest(http.MethodPost, "/register", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Contains(t, w.Body.String(), "iruyan_id is already taken")

	logStep(t, "RESULT", "Received 500 Internal Server Error for duplicate user")
}

// --- [Register] 重複登録（パスワード制約エラー） ---
func TestRegisterHandler_ShortPassword(t *testing.T) {
	router := setupTestRouter()
	t.Log("\n=== [INFO] Start: パスワードが短すぎるテスト ===")
	t.Cleanup(func() { t.Log("--- [INFO] End ---\n") })

	data := url.Values{}
	data.Set("iruyanId", "shortpwuser")
	data.Set("password", "a1b2") // ← 6文字未満
	data.Set("userName", "Short PW")
	data.Set("email", "shortpw@example.com")

	req, _ := http.NewRequest(http.MethodPost, "/register", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "password must be at least 6 characters long")
}

// --- [Register] 重複登録（パスワード制約エラー） ---
func TestRegisterHandler_WeakPassword_NoNumber(t *testing.T) {
	router := setupTestRouter()
	t.Log("\n=== [INFO] Start: パスワードが数字なし ===")
	t.Cleanup(func() { t.Log("--- [INFO] End ---\n") })

	data := url.Values{}
	data.Set("iruyanId", "nonumberuser")
	data.Set("password", "abcdefg") // ← 数字なし
	data.Set("userName", "No Number")
	data.Set("email", "nonumber@example.com")

	req, _ := http.NewRequest(http.MethodPost, "/register", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "password must contain at least one letter and one number")
}

// --- [Register] 重複登録（パスワード制約エラー） ---
func TestRegisterHandler_WeakPassword_NoLetter(t *testing.T) {
	router := setupTestRouter()
	t.Log("\n=== [INFO] Start: パスワードが英字なし ===")
	t.Cleanup(func() { t.Log("--- [INFO] End ---\n") })

	data := url.Values{}
	data.Set("iruyanId", "noletteruser")
	data.Set("password", "1234567") // ← 英字なし
	data.Set("userName", "No Letter")
	data.Set("email", "noletter@example.com")

	req, _ := http.NewRequest(http.MethodPost, "/register", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "password must contain at least one letter and one number")
}

// --- [Register] 重複登録（メールアドレス制約エラー） ---
func TestRegisterHandler_InvalidEmail(t *testing.T) {
	router := setupTestRouter()
	t.Log("\n=== [INFO] Start: メール形式が不正 ===")
	t.Cleanup(func() { t.Log("--- [INFO] End ---\n") })

	data := url.Values{}
	data.Set("iruyanId", "invalidemailuser")
	data.Set("password", "valid1Pass")
	data.Set("userName", "Invalid Email")
	data.Set("email", "invalid-email") // ← 不正な形式

	req, _ := http.NewRequest(http.MethodPost, "/register", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status 400 BadRequest")
	assert.Contains(t, w.Body.String(), "invalid email format", "Expected error message about invalid email format")
}
