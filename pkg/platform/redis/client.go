package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

// ConnectRedis inicializa la conexión a Redis
func ConnectRedis() error {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := Client.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("no se pudo conectar a Redis: %w", err)
	}

	fmt.Println("Conectado a Redis")
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
