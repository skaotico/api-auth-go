package auth

type RefreshData struct {
	UserId    string `json:"userId"`
	IP        string `json:"ip"`
	UserAgent string `json:"ua"`
	CreatedAt int64  `json:"createdAt"`
}
