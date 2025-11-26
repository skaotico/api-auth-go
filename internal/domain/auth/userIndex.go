// ============================================================
// @file: userIndex.go
// @author: Yosemar Andrade
// @date: 2025-11-26
// @lastModified: 2025-11-26
// @description: Define la estructura de índice de usuario para caché.
// ============================================================

package auth

// UserIndex representa el índice de sesión del usuario en caché.
type UserIndex struct {
	ActiveJwt     string `json:"activeJwt"`
	ActiveRefresh string `json:"activeRefresh"`
	LastLogin     int64  `json:"lastLogin"`
}
