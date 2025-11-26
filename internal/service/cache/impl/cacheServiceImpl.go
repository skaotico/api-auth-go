// ============================================================
// @file: cacheServiceImpl.go
// @author: Yosemar Andrade
// @date: 2025-11-26
// @lastModified: 2025-11-26
// @description: Implementación del servicio de caché con logging.
// ============================================================

package impl

import (
	"api-auth/internal/domain/auth"
	"api-auth/internal/domain/security"
	"api-auth/internal/service/cache"
	"api-auth/internal/service/cache/helper"
	"api-auth/pkg/platform/redis"
	"context"
	"encoding/json"
	"time"

	"go.uber.org/zap"
)

// CacheServiceImpl implementa operaciones de caching en Redis.
type CacheServiceImpl struct {
	log *zap.Logger
}

// NewCacheService crea una nueva instancia del servicio de caché.
func NewCacheService(logger *zap.Logger) cache.CacheService {
	logger.Info("Inicializando CacheService")
	return &CacheServiceImpl{log: logger}
}

// SaveTokens guarda JWT, Refresh y UserIndex en Redis.
//
// Parámetros:
//   - ctx: contexto para propagación y cancelación.
//   - jwt: token JWT.
//   - refresh: token Refresh.
//   - jwtData: información del JWT.
//   - refreshData: información del Refresh.
//   - jwtTTL: tiempo de expiración del JWT.
//   - refreshTTL: tiempo de expiración del Refresh.
//
// Retorna:
//   - Error si ocurre algún problema en la escritura en Redis.
func (s *CacheServiceImpl) SaveTokens(
	ctx context.Context,
	jwt string,
	refresh string,
	jwtData *auth.JwtData,
	refreshData *auth.RefreshData,
	jwtTTL time.Duration,
	refreshTTL time.Duration,
) error {

	s.log.Info("Guardando tokens en Redis",
		zap.String("userId", jwtData.UserId),
		zap.Duration("jwtTTL", jwtTTL),
		zap.Duration("refreshTTL", refreshTTL),
	)

	jBytes, err := json.Marshal(jwtData)
	if err != nil {
		s.log.Error("Error serializando JWT data", zap.Error(err))
		return err
	}

	rBytes, err := json.Marshal(refreshData)
	if err != nil {
		s.log.Error("Error serializando Refresh data", zap.Error(err))
		return err
	}

	// Guardar JWT
	if err := redis.Client.Set(ctx, helper.GetJwtKey(jwt), jBytes, jwtTTL).Err(); err != nil {
		s.log.Error("Error guardando JWT en Redis", zap.Error(err), zap.String("userId", jwtData.UserId))
		return err
	}

	// Guardar Refresh
	if err := redis.Client.Set(ctx, helper.GetRefreshKey(refresh), rBytes, refreshTTL).Err(); err != nil {
		s.log.Error("Error guardando Refresh en Redis", zap.Error(err), zap.String("userId", jwtData.UserId))
		return err
	}

	// Crear índice del usuario
	userIndex := auth.UserIndex{
		ActiveJwt:     jwt,
		ActiveRefresh: refresh,
		LastLogin:     time.Now().Unix(),
	}
	uBytes, err := json.Marshal(userIndex)
	if err != nil {
		s.log.Error("Error serializando UserIndex", zap.Error(err))
		return err
	}

	if err := redis.Client.Set(ctx, helper.GetUserKey(jwtData.UserId), uBytes, refreshTTL).Err(); err != nil {
		s.log.Error("Error guardando índice de usuario", zap.Error(err), zap.String("userId", jwtData.UserId))
		return err
	}

	s.log.Info("Tokens guardados correctamente", zap.Any("indice usuario", &userIndex))
	return nil
}

// GetJwtData obtiene datos del JWT desde Redis.
func (s *CacheServiceImpl) GetJwtData(ctx context.Context, jwt string) (*auth.JwtData, error) {
	s.log.Info("Obteniendo JWT desde Redis", zap.String("jwt", jwt))

	val, err := redis.Client.Get(ctx, helper.GetJwtKey(jwt)).Result()
	if err != nil {
		s.log.Error("Error obteniendo JWT desde Redis", zap.Error(err), zap.String("jwt", jwt))
		return nil, err
	}

	var data auth.JwtData
	json.Unmarshal([]byte(val), &data)

	s.log.Debug("JWT obtenido correctamente", zap.String("jwt", jwt))
	return &data, nil
}

// GetRefreshData obtiene datos del Refresh Token desde Redis.
func (s *CacheServiceImpl) GetRefreshData(ctx context.Context, refresh string) (*auth.RefreshData, error) {
	s.log.Info("Obteniendo Refresh desde Redis", zap.String("refresh", refresh))

	val, err := redis.Client.Get(ctx, helper.GetRefreshKey(refresh)).Result()
	if err != nil {
		s.log.Error("Error obteniendo Refresh en Redis", zap.Error(err), zap.String("refresh", refresh))
		return nil, err
	}

	var data auth.RefreshData
	json.Unmarshal([]byte(val), &data)

	s.log.Info("Refresh obtenido correctamente", zap.String("refresh", refresh))
	return &data, nil
}

// GetUserIndex obtiene el índice del usuario desde Redis.
func (s *CacheServiceImpl) GetUserIndex(ctx context.Context, userId string) (*auth.UserIndex, error) {
	s.log.Debug("Obteniendo índice del usuario desde Redis", zap.String("userId", userId))

	val, err := redis.Client.Get(ctx, helper.GetUserKey(userId)).Result()
	if err != nil {
		s.log.Error("Error obteniendo índice del usuario", zap.Error(err), zap.String("userId", userId))
		return nil, err
	}

	var data auth.UserIndex
	json.Unmarshal([]byte(val), &data)

	s.log.Debug("Índice del usuario obtenido correctamente", zap.String("userId", userId))
	return &data, nil
}

// DeleteAll elimina JWT, Refresh y UserIndex.
func (s *CacheServiceImpl) DeleteAll(ctx context.Context, userId string, jwt string, refresh string) error {
	s.log.Debug("Eliminando tokens y userIndex de Redis", zap.String("userId", userId))

	if err := redis.Client.Del(ctx, helper.GetJwtKey(jwt)).Err(); err != nil {
		s.log.Error("Error eliminando JWT", zap.Error(err), zap.String("userId", userId))
		return err
	}

	if err := redis.Client.Del(ctx, helper.GetRefreshKey(refresh)).Err(); err != nil {
		s.log.Error("Error eliminando Refresh", zap.Error(err), zap.String("userId", userId))
		return err
	}

	if err := redis.Client.Del(ctx, helper.GetUserKey(userId)).Err(); err != nil {
		s.log.Error("Error eliminando userIndex", zap.Error(err), zap.String("userId", userId))
		return err
	}

	s.log.Debug("Tokens y userIndex eliminados exitosamente", zap.String("userId", userId))
	return nil
}

// ============================================================
// Rate Limit Implementation
// ============================================================

// SaveRateLimit guarda la data de rate limit en Redis.
func (s *CacheServiceImpl) SaveRateLimit(ctx context.Context, data *security.RateLimitData) error {
	s.log.Debug("Guardando RateLimit",
		zap.String("key", data.Key),
		zap.Int64("limit", data.Limit),
		zap.Int64("expiresAt", data.ExpiresAt),
	)

	b, err := json.Marshal(data)
	if err != nil {
		s.log.Error("Error serializando RateLimit", zap.Error(err))
		return err
	}

	ttl := time.Until(time.Unix(data.ExpiresAt, 0))
	if err := redis.Client.Set(ctx, data.Key, b, ttl).Err(); err != nil {
		s.log.Error("Error guardando RateLimit en Redis", zap.Error(err), zap.String("key", data.Key))
		return err
	}

	s.log.Debug("RateLimit guardado correctamente", zap.String("key", data.Key))
	return nil
}

// GetRateLimit obtiene reglas de rate limiting desde Redis.
func (s *CacheServiceImpl) GetRateLimit(ctx context.Context, key string) (*security.RateLimitData, error) {
	s.log.Debug("Obteniendo RateLimit desde Redis", zap.String("key", key))

	val, err := redis.Client.Get(ctx, key).Result()
	if err != nil {
		s.log.Error("Error obteniendo RateLimit", zap.Error(err), zap.String("key", key))
		return nil, err
	}

	var data security.RateLimitData
	json.Unmarshal([]byte(val), &data)

	s.log.Debug("RateLimit obtenido correctamente", zap.String("key", key))
	return &data, nil
}
