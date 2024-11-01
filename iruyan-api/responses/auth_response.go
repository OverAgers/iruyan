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

// LogoutSuccessResponse ログアウト成功時のレスポンス
type LogoutSuccessResponse struct {
	Message string   `json:"message"`
	User    UserInfo `json:"user"`
}

// UserInfo ユーザー情報のレスポンス
type UserInfo struct {
	IruyanID string `json:"iruyan_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}
