package auth

type JwtData struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	// Roles     []string `json:"roles"`
	CreatedAt int64 `json:"createdAt"`
}
