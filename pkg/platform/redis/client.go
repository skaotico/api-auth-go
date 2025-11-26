// ============================================================
// @file: client.go
// @author: Yosemar Andrade
// @date: 2025-11-26
// @lastModified: 2025-11-26
// @description: Paquete redis maneja la conexión y operaciones con Redis.
// ============================================================

package redis

import (
	"api-auth/pkg/logger"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// Client es la instancia global del cliente Redis.
var Client *redis.Client

// ConnectRedis inicializa la conexión a Redis.
//
// Parámetros:
//   - No recibe parámetros.
//
// Retorna:
//   - error: retorna error si falla el ping a Redis.
//
// Errores:
//   - Retorna error si no se puede conectar a Redis.
func ConnectRedis() error {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Username: "Skaotico",
		Password: "1q2w3e4r",
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := Client.Ping(ctx).Err(); err != nil {
		logger.Log.Error("No se pudo conectar a Redis", zap.Error(err))
		return fmt.Errorf("no se pudo conectar a Redis: %w", err)
	}

	logger.Log.Info("Conectado a Redis")
	return nil
}

// CheckRedis valida el estado de la conexión (para health check)
func CheckRedis() error {
	if Client == nil {
		return fmt.Errorf("cliente redis no inicializado")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return Client.Ping(ctx).Err()
}
