package cache

import (
	"context"
	"time"

	domain "api-auth/internal/domain/auth"
)

// TokenRepository define las operaciones necesarias para almacenar y recuperar
// JWTs, refresh tokens e índices de usuario desde Redis.
type TokenRepository interface {

	// SaveTokens persiste:
	// - JWT asociado al usuario
	// - Refresh Token asociado al usuario
	// - Índice auth:user:<idUsuario> para invalidez masiva
	SaveTokens(
		ctx context.Context,
		jwt string,
		refresh string,
		jwtData domain.JwtData,
		refreshData domain.RefreshData,
		jwtTTL time.Duration,
		refreshTTL time.Duration,
	) error

	// GetUserByJwt obtiene información del usuario desde la key auth:jwt:<jwt>
	GetUserByJwt(
		ctx context.Context,
		jwt string,
	) (*domain.JwtData, error)

	// GetUserByRefresh obtiene información desde auth:refresh:<refresh>
	GetUserByRefresh(
		ctx context.Context,
		refresh string,
	) (*domain.RefreshData, error)
}
