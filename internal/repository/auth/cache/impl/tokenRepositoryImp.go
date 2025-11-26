package impl

import (
	"context"
	"encoding/json"
	"time"

	domain "api-auth/internal/domain/auth"
	"api-auth/internal/repository/auth/cache"
	"api-auth/pkg/platform/redis"
)

// Implementación del repositorio
type tokenRepositoryImp struct{}

// Constructor público
func NewTokenRepository() cache.TokenRepository {
	return &tokenRepositoryImp{}
}

// SaveTokens guarda:
// 1. auth:jwt:<jwt>
// 2. auth:refresh:<refresh>
// 3. auth:user:<userId>
func (r *tokenRepositoryImp) SaveTokens(
	ctx context.Context,
	jwt string,
	refresh string,
	jwtData domain.JwtData,
	refreshData domain.RefreshData,
	jwtTTL time.Duration,
	refreshTTL time.Duration,
) error {

	// Convertir structs a JSON
	jBytes, err := json.Marshal(jwtData)
	if err != nil {
		return err
	}

	rBytes, err := json.Marshal(refreshData)
	if err != nil {
		return err
	}

	// ====== Guardar JWT ======
	if err := redis.Client.Set(
		ctx,
		"auth:jwt:"+jwt,
		jBytes,
		jwtTTL,
	).Err(); err != nil {
		return err
	}

	// ====== Guardar Refresh ======
	if err := redis.Client.Set(
		ctx,
		"auth:refresh:"+refresh,
		rBytes,
		refreshTTL,
	).Err(); err != nil {
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
		return err
	}

	return redis.Client.Set(
		ctx,
		"auth:user:"+jwtData.UserId,
		uBytes,
		refreshTTL,
	).Err()
}

// GetUserByJwt obtiene auth:jwt:<jwt>
func (r *tokenRepositoryImp) GetUserByJwt(
	ctx context.Context,
	jwt string,
) (*domain.JwtData, error) {

	val, err := redis.Client.Get(ctx, "auth:jwt:"+jwt).Result()
	if err != nil {
		return nil, err
	}

	var data domain.JwtData
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		return nil, err
	}

	return &data, nil
}

// GetUserByRefresh obtiene auth:refresh:<refresh>
func (r *tokenRepositoryImp) GetUserByRefresh(
	ctx context.Context,
	refresh string,
) (*domain.RefreshData, error) {

	val, err := redis.Client.Get(ctx, "auth:refresh:"+refresh).Result()
	if err != nil {
		return nil, err
	}

	var data domain.RefreshData
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		return nil, err
	}

	return &data, nil
}
