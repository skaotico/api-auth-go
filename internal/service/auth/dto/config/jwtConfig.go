// ============================================================
// @file: jwtConfig.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: Configuración para JWT.
// ============================================================

// Package config contiene las estructuras de configuración.
package config

import "time"

// JWTConfig define la configuración para la generación y validación de tokens JWT.
type JWTConfig struct {
	Secret     string
	Expiration time.Duration
	RefreshTTL time.Duration
}
