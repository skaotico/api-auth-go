// ============================================================
// @file: refreshData.go
// @author: Yosemar Andrade
// @date: 2025-11-26
// @lastModified: 2025-11-26
// @description: Define la estructura de datos para el refresh token.
// ============================================================

package auth

// RefreshData representa los datos asociados a un refresh token.
type RefreshData struct {
	UserId    string `json:"userId"`
	IP        string `json:"ip"`
	UserAgent string `json:"ua"`
	CreatedAt int64  `json:"createdAt"`
}
