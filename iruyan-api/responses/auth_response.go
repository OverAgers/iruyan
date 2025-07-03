package responses

// LoginSuccessResponse ログイン成功時のレスポンス
type LoginSuccessResponse struct {
	Message string   `json:"message" example:"Login successful"`
	User    UserInfo `json:"user"`
	Token   string   `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// RegisterSuccessResponse 新規登録成功時のレスポンス
type RegisterSuccessResponse struct {
	Message string   `json:"message" example:"Registration successful"`
	User    UserInfo `json:"user"`
	Token   string   `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// LogoutSuccessResponse ログアウト成功時のレスポンス
type LogoutSuccessResponse struct {
	Message string   `json:"message" example:"Logout successful"`
	User    UserInfo `json:"user"`
}

// UserInfo ユーザー情報のレスポンス
type UserInfo struct {
	IruyanID string `json:"iruyanId" example:"johndoe"`
	Name     string `json:"userName" example:"John Doe"`
	Email    string `json:"email" example:"john@example.com"`
}
