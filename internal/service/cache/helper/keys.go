// ============================================================
// @file: keys.go
// @author: Yosemar Andrade
// @date: 2025-11-26
// @lastModified: 2025-11-26
// @description: Helper para generación de claves Redis.
// ============================================================

package helper

import "fmt"

const (
	prefixJwt     = "auth:jwt:"
	prefixRefresh = "auth:refresh:"
	prefixUser    = "auth:user:"
)

// GetJwtKey genera la clave para almacenar el JWT.
func GetJwtKey(token string) string {
	return fmt.Sprintf("%s%s", prefixJwt, token)
}

// GetRefreshKey genera la clave para almacenar el Refresh Token.
func GetRefreshKey(token string) string {
	return fmt.Sprintf("%s%s", prefixRefresh, token)
}

// GetUserKey genera la clave para almacenar el índice de usuario.
func GetUserKey(userId string) string {
	return fmt.Sprintf("%s%s", prefixUser, userId)
}
