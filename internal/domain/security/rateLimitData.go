// ============================================================
// @file: rateLimit.go
// @author: Yosemar Andrade
// @date: 2025-11-26
// @lastModified: 2025-11-26
// @description: Define la estructura de datos usada para control de Rate Limiting en Redis.
// ============================================================

package security

// RateLimitData representa la información almacenada en Redis
// para aplicar restricciones de peticiones.
type RateLimitData struct {
	Key       string `json:"key"`       // clave en redis (ip|user)
	Limit     int64  `json:"limit"`     // total permitido
	Attempts  int64  `json:"attempts"`  // cuántos intentos se han usado
	ExpiresAt int64  `json:"expiresAt"` // timestamp de expiración
}
