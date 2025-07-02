package auth

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"iruyan-api/infrastructure"
	"iruyan-api/models"
	"iruyan-api/pkg/errdefs"
	"iruyan-api/usecases/auth/mocks"
)

// logStep はテスト内ログを整えるヘルパー（ASCII のみ）
func logStep(t *testing.T, label, msg string) {
	t.Helper()
	t.Logf("[%-10s] %s", label, msg)
}

// テスト用のGinルーターをセットアップ
func setupTestRouterWithMock(mockAuthUsecase *mocks.AuthUsecase) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	authHandler := NewAuthHandler(mockAuthUsecase)
	r.POST("/login", authHandler.LoginHandler)
	r.POST("/register", authHandler.RegisterHandler)
	r.POST("/logout", authHandler.LogoutHandler)

	return r
}

// TestMain は全テスト実行前の共通初期化処理
func TestMain(m *testing.M) {
	infrastructure.InitTestDB(nil) // テスト用DB初期化
	m.Run()
}

// --- [Login] 正常系 ---
func TestLoginHandler_Success(t *testing.T) {
	t.Logf("\n=== [INFO] Start: 正常なログイン成功テスト（Mock使用） ===")
	t.Cleanup(func() {
		t.Logf("--- [INFO] End: 正常なログイン成功テスト ---\n")
	})

	mockUsecase := new(mocks.AuthUsecase)

	mockUsecase.On("LoginUser", "testuser", "pass1234").Return(&models.User{
		ID:       1,
		IruyanID: "testuser",
		Name:     "テストユーザー",
		Email:    "testuser@example.com",
	}, nil)

	router := setupTestRouterWithMock(mockUsecase)

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

	mockUsecase.AssertExpectations(t)
}

// --- [Login] パラメータ不足 ---
func TestLoginHandler_MissingParams(t *testing.T) {
	t.Logf("\n=== [INFO] Start: パラメータ不足テスト（Mock使用） ===")
	t.Cleanup(func() {
		t.Logf("--- [INFO] End: パラメータ不足テスト ---\n")
	})

	mockUsecase := new(mocks.AuthUsecase)
	router := setupTestRouterWithMock(mockUsecase)

	// パラメータなしのリクエスト
	req, _ := http.NewRequest(http.MethodPost, "/login", nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `iruyanId and password are required`)

	logStep(t, "RESULT", "Received 400 Bad Request with expected error message")

	// AuthUsecase.Login は呼び出されていないことを確認（呼び出しがない前提）
	mockUsecase.AssertExpectations(t)
}

// --- [Login] 存在しないユーザー ---
func TestLoginHandler_UserNotFound(t *testing.T) {
	t.Log("\n=== [INFO] Start: 存在しないユーザーのテスト（Mock使用） ===")
	t.Cleanup(func() { t.Log("--- [INFO] End: 存在しないユーザーのテスト ---\n") })

	mockUsecase := new(mocks.AuthUsecase)
	mockUsecase.On("LoginUser", "nonexistent", "wrongpass").
		Return((*models.User)(nil), errdefs.ErrUserNotFound)

	router := setupTestRouterWithMock(mockUsecase)

	data := url.Values{}
	data.Set("iruyanId", "nonexistent")
	data.Set("password", "wrongpass")

	req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "user not found")

	mockUsecase.AssertExpectations(t)
}

// --- [Login] パスワード間違い ---
func TestLoginHandler_WrongPassword(t *testing.T) {
	t.Logf("\n=== [INFO] Start: パスワード不一致テスト（Mock使用） ===")
	t.Cleanup(func() {
		t.Logf("--- [INFO] End: パスワード不一致テスト ---\n")
	})

	mockUsecase := new(mocks.AuthUsecase)
	router := setupTestRouterWithMock(mockUsecase)

	// パスワード不一致時に返されるエラーを設定
	mockUsecase.
		On("LoginUser", "testuser", "incorrect").
		Return(nil, errdefs.ErrInvalidPassword)

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

	mockUsecase.AssertExpectations(t)
}

// --- [Register] 登録成功 ---
func TestRegisterHandler_Success(t *testing.T) {
	t.Logf("\n=== [INFO] Start: 登録成功テスト（Mock使用） ===")
	t.Cleanup(func() {
		t.Logf("--- [INFO] End: 登録成功テスト ---\n")
	})

	mockUsecase := new(mocks.AuthUsecase)
	router := setupTestRouterWithMock(mockUsecase)

	// Registerが正常に完了する場合はnilを返す
	mockUsecase.
		On("RegisterUser", "New User", "newuser", "securepass1234", "newuser@example.com").
		Return(&models.User{
			Name:     "New User",
			IruyanID: "newuser",
			Email:    "newuser@example.com",
		}, nil)

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

	mockUsecase.AssertExpectations(t)
}

// --- [Register] パラメータ不足 ---
func TestRegisterHandler_MissingParams(t *testing.T) {
	t.Logf("\n=== [INFO] Start: パラメータ不足による登録失敗テスト（Mock使用） ===")
	t.Cleanup(func() {
		t.Logf("--- [INFO] End: パラメータ不足による登録失敗テスト ---\n")
	})

	mockUsecase := new(mocks.AuthUsecase)
	router := setupTestRouterWithMock(mockUsecase)

	req, _ := http.NewRequest(http.MethodPost, "/register", nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Missing required parameter(s)")

	logStep(t, "RESULT", "Received 400 Bad Request due to missing params")

	// このケースでは usecase.Register は呼ばれない想定なので、呼ばれていないことを確認
	mockUsecase.AssertNotCalled(t, "Register", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

// --- [Register] 重複登録（iruyanIDのユニーク制約エラー） ---
func TestRegisterHandler_DuplicateUser(t *testing.T) {
	t.Logf("\n=== [INFO] Start: 重複登録エラーテスト（Mock使用） ===")
	t.Cleanup(func() {
		t.Logf("--- [INFO] End: 重複登録エラーテスト ---\n")
	})

	mockUsecase := new(mocks.AuthUsecase)
	router := setupTestRouterWithMock(mockUsecase)

	// 入力データ
	data := url.Values{}
	data.Set("iruyanId", "duplicateuser")
	data.Set("password", "samepass1234")
	data.Set("userName", "Duplicate User")
	data.Set("email", "duplicate@example.com")

	// モックの期待値設定
	mockUsecase.
		On("RegisterUser", "Duplicate User", "duplicateuser", "samepass1234", "duplicate@example.com").
		Return(nil, errdefs.ErrDuplicateIruyanID)

	req, _ := http.NewRequest(http.MethodPost, "/register", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Contains(t, w.Body.String(), "iruyan_id is already taken")

	logStep(t, "RESULT", "Received 409 Conflict due to duplicate user")

	mockUsecase.AssertExpectations(t)
}

// --- [Register] 重複登録（パスワード制約エラー） ---
func TestRegisterHandler_ShortPassword(t *testing.T) {
	t.Log("\n=== [INFO] Start: パスワードが短すぎるテスト（Mock使用） ===")
	t.Cleanup(func() { t.Log("--- [INFO] End ---\n") })

	mockUsecase := new(mocks.AuthUsecase)
	router := setupTestRouterWithMock(mockUsecase)

	data := url.Values{}
	data.Set("iruyanId", "shortpwuser")
	data.Set("password", "a1b2") // ← 6文字未満
	data.Set("userName", "Short PW")
	data.Set("email", "shortpw@example.com")

	// モックの期待される返り値設定（バリデーションエラー）
	mockUsecase.
		On("RegisterUser", "Short PW", "shortpwuser", "a1b2", "shortpw@example.com").
		Return(nil, errdefs.ErrPasswordTooShort)

	req, _ := http.NewRequest(http.MethodPost, "/register", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "password must be at least 6 characters long")

	logStep(t, "RESULT", "Received 400 Bad Request for short password")
	mockUsecase.AssertExpectations(t)
}

// --- [Register] 重複登録（パスワード制約エラー） ---
func TestRegisterHandler_WeakPassword_NoNumber(t *testing.T) {
	t.Log("\n=== [INFO] Start: パスワードが数字なし ===")
	t.Cleanup(func() { t.Log("--- [INFO] End ---\n") })

	mockUsecase := new(mocks.AuthUsecase)
	router := setupTestRouterWithMock(mockUsecase)

	iruyanId := "nonumberuser"
	password := "abcdefg" // ← 数字なし
	userName := "No Number"
	email := "nonumber@example.com"

	data := url.Values{}
	data.Set("iruyanId", iruyanId)
	data.Set("password", password)
	data.Set("userName", userName)
	data.Set("email", email)

	// モックの挙動を設定：バリデーションエラーを返す
	mockUsecase.
		On("RegisterUser", userName, iruyanId, password, email).
		Return(nil, errdefs.ErrPasswordMissingChars)

	req, _ := http.NewRequest(http.MethodPost, "/register", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "password must contain at least one letter and one number")

	logStep(t, "RESULT", "Received 400 Bad Request for password with no number")
	mockUsecase.AssertExpectations(t)
}

// --- [Register] 重複登録（パスワード制約エラー） ---
func TestRegisterHandler_WeakPassword_NoLetter(t *testing.T) {
	t.Log("\n=== [INFO] Start: パスワードが英字なし ===")
	t.Cleanup(func() { t.Log("--- [INFO] End ---\n") })

	mockUsecase := new(mocks.AuthUsecase)
	router := setupTestRouterWithMock(mockUsecase)

	iruyanId := "noletteruser"
	password := "1234567" // ← 英字なし
	userName := "No Letter"
	email := "noletter@example.com"

	data := url.Values{}
	data.Set("iruyanId", iruyanId)
	data.Set("password", password)
	data.Set("userName", userName)
	data.Set("email", email)

	// モックの挙動：パスワードエラーを返す
	mockUsecase.
		On("RegisterUser", userName, iruyanId, password, email).
		Return(nil, errdefs.ErrPasswordMissingChars)

	req, _ := http.NewRequest(http.MethodPost, "/register", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "password must contain at least one letter and one number")

	logStep(t, "RESULT", "Received 400 Bad Request for password with no letter")
	mockUsecase.AssertExpectations(t)
}

// --- [Register] 重複登録（メールアドレス制約エラー） ---
func TestRegisterHandler_InvalidEmail(t *testing.T) {
	t.Log("\n=== [INFO] Start: メール形式が不正 ===")
	t.Cleanup(func() { t.Log("--- [INFO] End ---\n") })

	mockUsecase := new(mocks.AuthUsecase)
	router := setupTestRouterWithMock(mockUsecase)

	iruyanId := "invalidemailuser"
	password := "valid1Pass"
	userName := "Invalid Email"
	email := "invalid-email" // ← 不正な形式

	data := url.Values{}
	data.Set("iruyanId", iruyanId)
	data.Set("password", password)
	data.Set("userName", userName)
	data.Set("email", email)

	// モックの挙動：不正なメール形式のエラーを返す
	mockUsecase.
		On("RegisterUser", userName, iruyanId, password, email).
		Return(nil, errdefs.ErrInvalidEmail)

	req, _ := http.NewRequest(http.MethodPost, "/register", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status 400 BadRequest")
	assert.Contains(t, w.Body.String(), "invalid email format", "Expected error message about invalid email format")

	logStep(t, "RESULT", "Received 400 Bad Request for invalid email format")
	mockUsecase.AssertExpectations(t)
}

// --- [Logout] ログアウト成功 ---
func TestLogoutHandler_Success(t *testing.T) {
	t.Log("\n=== [INFO] Start: 正常なログアウト処理 ===")
	t.Cleanup(func() { t.Log("--- [INFO] End ---\n") })

	mockUsecase := new(mocks.AuthUsecase)
	router := setupTestRouterWithMock(mockUsecase)

	iruyanId := "logoutuser"

	data := url.Values{}
	data.Set("iruyanId", iruyanId)

	// モックの期待値設定：正常終了（エラーなし）
	mockUsecase.
		On("LogoutUser", iruyanId).
		Return(&models.User{IruyanID: iruyanId}, nil)

	req, _ := http.NewRequest(http.MethodPost, "/logout", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusOK, w.Code, "Expected status 200 OK")
	assert.Contains(t, w.Body.String(), "Logout successful")
	assert.Contains(t, w.Body.String(), `"logoutuser"`)

	logStep(t, "RESULT", "Received 200 OK and logout message")
	mockUsecase.AssertExpectations(t)
}

// --- [Logout] 存在しないユーザー ---
func TestLogoutHandler_UserNotFound(t *testing.T) {
	t.Log("\n=== [INFO] Start: 存在しないユーザーでログアウト ===")
	t.Cleanup(func() { t.Log("--- [INFO] End ---\n") })

	mockUsecase := new(mocks.AuthUsecase)
	router := setupTestRouterWithMock(mockUsecase)

	iruyanId := "nonexistentuser"

	// モックの設定：user not found エラーを返す
	mockUsecase.
		On("LogoutUser", iruyanId).
		Return(nil, errdefs.ErrUserNotFound)

	data := url.Values{}
	data.Set("iruyanId", iruyanId)

	req, _ := http.NewRequest(http.MethodPost, "/logout", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusUnauthorized, w.Code, "Expected status 401 Unauthorized")
	assert.Contains(t, w.Body.String(), "user not found")

	logStep(t, "RESULT", "Received 401 Unauthorized for non-existent user")
	mockUsecase.AssertExpectations(t)
}

// --- [Logout] パラメータ未指定 ---
func TestLogoutHandler_MissingParams(t *testing.T) {
	t.Log("\n=== [INFO] Start: パラメータ未指定のログアウト ===")
	t.Cleanup(func() { t.Log("--- [INFO] End ---\n") })

	mockUsecase := new(mocks.AuthUsecase)
	router := setupTestRouterWithMock(mockUsecase)

	// パラメータなしのPOSTリクエスト
	req, _ := http.NewRequest(http.MethodPost, "/logout", nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("[RESPONSE] StatusCode: %d", w.Code)
	t.Logf("[RESPONSE] Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "iruyanId is required")

	logStep(t, "RESULT", "Received 400 Bad Request for missing iruyanId")
}
