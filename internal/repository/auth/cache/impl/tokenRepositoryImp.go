// ============================================================
// @file: tokenRepositoryImp.go
// @author: Yosemar Andrade
// @date: 2025-11-26
// @lastModified: 2025-11-26
// @description: Implementación del repositorio de tokens en caché usando Redis.
// ============================================================

package impl

import (
	"api-auth/pkg/logger"
	"context"
	"encoding/json"
	"time"

	domain "api-auth/internal/domain/auth"
	"api-auth/internal/repository/auth/cache"
	"api-auth/pkg/platform/redis"

	"go.uber.org/zap"
)

// tokenRepositoryImp implementa la interfaz TokenRepository.
type tokenRepositoryImp struct{}

// NewTokenRepository crea una nueva instancia del repositorio de tokens.
//
// Parámetros:
//   - No recibe parámetros.
//
// Retorna:
//   - cache.TokenRepository: interfaz del repositorio de tokens.
//
// Errores:
//   - No retorna errores.
func NewTokenRepository() cache.TokenRepository {
	return &tokenRepositoryImp{}
}

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
//
// Errores:
//   - Retorna error si falla `json.Marshal` o `redis.Client.Set`.
func (r *tokenRepositoryImp) SaveTokens(
	ctx context.Context,
	jwt string,
	refresh string,
	jwtData domain.JwtData,
	refreshData domain.RefreshData,
	jwtTTL time.Duration,
	refreshTTL time.Duration,
) error {

	logger.Log.Debug("Guardando tokens en caché", zap.String("userId", jwtData.UserId))

	// Convertir structs a JSON
	jBytes, err := json.Marshal(jwtData)
	if err != nil {
		logger.Log.Error("Error al serializar jwtData", zap.Error(err))
		return err
	}

	rBytes, err := json.Marshal(refreshData)
	if err != nil {
		logger.Log.Error("Error al serializar refreshData", zap.Error(err))
		return err
	}

	// ====== Guardar JWT ======
	if err := redis.Client.Set(
		ctx,
		"auth:jwt:"+jwt,
		jBytes,
		jwtTTL,
	).Err(); err != nil {
		logger.Log.Error("Error al guardar JWT en Redis", zap.Error(err))
		return err
	}

	// ====== Guardar Refresh ======
	if err := redis.Client.Set(
		ctx,
		"auth:refresh:"+refresh,
		rBytes,
		refreshTTL,
	).Err(); err != nil {
		logger.Log.Error("Error al guardar Refresh Token en Redis", zap.Error(err))
		return err
	}

	// ====== Guardar índice por usuario ======
	userIndex := domain.UserIndex{
		ActiveJwt:     jwt,
		ActiveRefresh: refresh,
		LastLogin:     time.Now().Unix(),
	}

	uBytes, err := json.Marshal(userIndex)
	if err != nil {
		logger.Log.Error("Error al serializar userIndex", zap.Error(err))
		return err
	}

	if err := redis.Client.Set(
		ctx,
		"auth:user:"+jwtData.UserId,
		uBytes,
		refreshTTL,
	).Err(); err != nil {
		logger.Log.Error("Error al guardar índice de usuario en Redis", zap.Error(err))
		return err
	}

	return nil
}

// GetUserByJwt obtiene información del usuario desde la key auth:jwt:<jwt>.
//
// Parámetros:
//   - ctx: contexto de la operación.
//   - jwt: token JWT.
//
// Retorna:
//   - *domain.JwtData: datos del JWT.
//   - error: error si no se encuentra o falla la lectura.
//
// Errores:
//   - Retorna error si falla `redis.Client.Get` o `json.Unmarshal`.
func (r *tokenRepositoryImp) GetUserByJwt(
	ctx context.Context,
	jwt string,
) (*domain.JwtData, error) {

	val, err := redis.Client.Get(ctx, "auth:jwt:"+jwt).Result()
	if err != nil {
		logger.Log.Warn("JWT no encontrado en caché", zap.String("jwt_prefix", jwt[:10]+"..."))
		return nil, err
	}

	var data domain.JwtData
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		logger.Log.Error("Error al deserializar jwtData", zap.Error(err))
		return nil, err
	}

	return &data, nil
}

// GetUserByRefresh obtiene información desde auth:refresh:<refresh>.
//
// Parámetros:
//   - ctx: contexto de la operación.
//   - refresh: refresh token.
//
// Retorna:
//   - *domain.RefreshData: datos del refresh token.
//   - error: error si no se encuentra o falla la lectura.
//
// Errores:
//   - Retorna error si falla `redis.Client.Get` o `json.Unmarshal`.
func (r *tokenRepositoryImp) GetUserByRefresh(
	ctx context.Context,
	refresh string,
) (*domain.RefreshData, error) {

	val, err := redis.Client.Get(ctx, "auth:refresh:"+refresh).Result()
	if err != nil {
		logger.Log.Warn("Refresh Token no encontrado en caché", zap.String("refresh_prefix", refresh[:10]+"..."))
		return nil, err
	}

	var data domain.RefreshData
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		logger.Log.Error("Error al deserializar refreshData", zap.Error(err))
		return nil, err
	}

	return &data, nil
}
