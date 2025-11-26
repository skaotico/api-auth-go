package impl

import (
	"api-auth/internal/domain/auth"
	"api-auth/internal/service/cache"
	"api-auth/pkg/platform/redis"
	"context"
	"encoding/json"
	"time"
)

type cacheServiceImpl struct{}

func NewCacheService() cache.CacheService {
	return &cacheServiceImpl{}
}

// SaveTokens guarda JWT, Refresh y UserIndex en Redis.
//
// Se utiliza un TTL independiente para JWT y Refresh, mientras que
// el índice del usuario hereda el TTL del refresh para mantener coherencia.
func (s *cacheServiceImpl) SaveTokens(
	ctx context.Context,
	jwt string,
	refresh string,
	jwtData *auth.JwtData,
	refreshData *auth.RefreshData,
	jwtTTL time.Duration,
	refreshTTL time.Duration,
) error {

	jBytes, _ := json.Marshal(jwtData)
	rBytes, _ := json.Marshal(refreshData)

	// Guardar JWT
	if err := redis.Client.Set(ctx, "auth:jwt:"+jwt, jBytes, jwtTTL).Err(); err != nil {
		return err
	}

	// Guardar Refresh
	if err := redis.Client.Set(ctx, "auth:refresh:"+refresh, rBytes, refreshTTL).Err(); err != nil {
		return err
	}

	// Crear índice del usuario
	userIndex := auth.UserIndex{
		ActiveJwt:     jwt,
		ActiveRefresh: refresh,
		LastLogin:     time.Now().Unix(),
	}
	uBytes, _ := json.Marshal(userIndex)

	return redis.Client.Set(ctx, "auth:user:"+jwtData.UserId, uBytes, refreshTTL).Err()
}

// GetJwtData obtiene datos del JWT desde Redis.
func (s *cacheServiceImpl) GetJwtData(ctx context.Context, jwt string) (*auth.JwtData, error) {
	val, err := redis.Client.Get(ctx, "auth:jwt:"+jwt).Result()
	if err != nil {
		return nil, err
	}

	var data auth.JwtData
	json.Unmarshal([]byte(val), &data)

	return &data, nil
}

// GetRefreshData obtiene datos del Refresh Token desde Redis.
func (s *cacheServiceImpl) GetRefreshData(ctx context.Context, refresh string) (*auth.RefreshData, error) {
	val, err := redis.Client.Get(ctx, "auth:refresh:"+refresh).Result()
	if err != nil {
		return nil, err
	}

	var data auth.RefreshData
	json.Unmarshal([]byte(val), &data)

	return &data, nil
}

// GetUserIndex obtiene el índice del usuario desde Redis.
func (s *cacheServiceImpl) GetUserIndex(ctx context.Context, userId string) (*auth.UserIndex, error) {
	val, err := redis.Client.Get(ctx, "auth:user:"+userId).Result()
	if err != nil {
		return nil, err
	}

	var data auth.UserIndex
	json.Unmarshal([]byte(val), &data)

	return &data, nil
}

// DeleteAll elimina JWT, Refresh y UserIndex para el usuario.
//
// Es ideal para logout o revocación de tokens.
func (s *cacheServiceImpl) DeleteAll(ctx context.Context, userId string, jwt string, refresh string) error {

	if err := redis.Client.Del(ctx, "auth:jwt:"+jwt).Err(); err != nil {
		return err
	}

	if err := redis.Client.Del(ctx, "auth:refresh:"+refresh).Err(); err != nil {
		return err
	}

	return redis.Client.Del(ctx, "auth:user:"+userId).Err()
}
