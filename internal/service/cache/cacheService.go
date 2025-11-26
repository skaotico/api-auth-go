package cache

import (
	authDomain "api-auth/internal/domain/auth"
	"context"
	"time"
)

// CacheService define las operaciones de almacenamiento en caché
// para JWT, Refresh Tokens y el índice del usuario.
// Esta capa contiene lógica de aplicación y abstrae el uso de Redis.
type CacheService interface {

	// SaveTokens guarda el JWT, el Refresh Token y el índice del usuario,
	// aplicando sus TTL respectivos.
	//
	// jwt: valor del JWT.
	// refresh: valor del refresh token.
	// jwtData: datos asociados al JWT.
	// refreshData: datos asociados al refresh token.
	// jwtTTL: tiempo de vida del JWT.
	// refreshTTL: tiempo de vida del refresh token.
	SaveTokens(
		ctx context.Context,
		jwt string,
		refresh string,
		jwtData *authDomain.JwtData,
		refreshData *authDomain.RefreshData,
		jwtTTL time.Duration,
		refreshTTL time.Duration,
	) error

	// GetJwtData obtiene los datos del JWT desde la caché.
	GetJwtData(ctx context.Context, jwt string) (*authDomain.JwtData, error)

	// GetRefreshData obtiene los datos del Refresh Token desde la caché.
	GetRefreshData(ctx context.Context, refresh string) (*authDomain.RefreshData, error)

	// GetUserIndex obtiene el índice del usuario (último login, tokens activos).
	GetUserIndex(ctx context.Context, userId string) (*authDomain.UserIndex, error)

	// DeleteAll elimina JWT, Refresh y UserIndex asociados a un usuario.
	DeleteAll(ctx context.Context, userId string, jwt string, refresh string) error
}
