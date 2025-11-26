// ============================================================
// @file: jwtData.go
// @author: Yosemar Andrade
// @date: 2025-11-26
// @lastModified: 2025-11-26
// @description: Define la estructura de datos contenida en el token JWT.
// ============================================================

package auth

// JwtData representa los datos payload del token JWT.
type JwtData struct {
	TokenID  string `json:"token_id"`
	UserId   string `json:"userId"`
	Username string `json:"username"`
	// Roles     []string `json:"roles"`
	CreatedAt int64 `json:"createdAt"`
}
