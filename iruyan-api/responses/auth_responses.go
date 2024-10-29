package responses

// LoginSuccessResponse ログイン成功時のレスポンス
type LoginSuccessResponse struct {
	Message string   `json:"message"`
	User    UserInfo `json:"user"`
}

// RegisterSuccessResponse 新規登録成功時のレスポンス
type RegisterSuccessResponse struct {
	Message string   `json:"message"`
	User    UserInfo `json:"user"`
}

// UserInfo ユーザー情報のレスポンス
type UserInfo struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Task     string `json:"task,omitempty"`
	Email    string `json:"email"`
}

// ErrorResponse エラーレスポンス
type ErrorResponse struct {
	Message string `json:"message"`
}
