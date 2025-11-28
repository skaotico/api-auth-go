// ============================================================
// @file: tokenRepository.go
// @author: Yosemar Andrade
// @date: 2025-11-26
// @lastModified: 2025-11-26
// @description: Define la interfaz para el repositorio de tokens en caché.
// ============================================================

package cache

import (
	"context"
	"time"

	domain "api-auth/internal/domain/auth"
)

// TokenRepository define las operaciones necesarias para almacenar y recuperar
// JWTs, refresh tokens e índices de usuario desde Redis.
type TokenRepository interface {

	// SaveTokens persiste los tokens y el índice de usuario en caché.
	//
	// Parámetros:
	//   - ctx: contexto de la operación.
	//   - jwt: token JWT.
	//   - refresh: refresh token.
	//   - jwtData: datos asociados al JWT.
	//   - refreshData: datos asociados al refresh token.
	//   - jwtTTL: tiempo de vida del JWT.
	//   - refreshTTL: tiempo de vida del refresh token.
	//
	// Retorna:
	//   - error: error si falla la persistencia en caché.
	SaveTokens(
		ctx context.Context,
		jwt string,
		refresh string,
		jwtData domain.JwtData,
		refreshData domain.RefreshData,
		jwtTTL time.Duration,
		refreshTTL time.Duration,
	) error

	// GetUserByJwt obtiene información del usuario desde la key auth:jwt:<jwt>.
	//
	// Parámetros:
	//   - ctx: contexto de la operación.
	//   - jwt: token JWT.
	//
	// Retorna:
	//   - *domain.JwtData: datos del JWT.
	//   - error: error si no se encuentra o falla la lectura.
	GetUserByJwt(
		ctx context.Context,
		jwt string,
	) (*domain.JwtData, error)

	// GetUserByRefresh obtiene información desde auth:refresh:<refresh>.
	//
	// Parámetros:
	//   - ctx: contexto de la operación.
	//   - refresh: refresh token.
	//
	// Retorna:
	//   - *domain.RefreshData: datos del refresh token.
	//   - error: error si no se encuentra o falla la lectura.
	GetUserByRefresh(
		ctx context.Context,
		refresh string,
	) (*domain.RefreshData, error)
}
