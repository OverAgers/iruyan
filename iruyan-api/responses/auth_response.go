package responses

// LoginSuccessResponse ログイン成功時のレスポンス
type LoginSuccessResponse struct {
	Message string   `json:"message" example:"Login successful"`
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
	IruyanID string `json:"iruyanId" example:"johndoe"`
	Name     string `json:"name" example:"John Doe"`
	Email    string `json:"email" example:"john@example.com"`
}
