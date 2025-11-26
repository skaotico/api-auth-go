// ============================================================
// @file: main.go
// @author: Yosemar Andrade
// @date: 2025-11-19
// @lastModified: 2025-11-19
// @description: Punto de entrada del servicio de autenticación. Se encarga de
// inicializar el logger, cargar la configuración, establecer conexión a la base
// de datos y levantar el servidor HTTP.
// ============================================================

package main

import (
	_ "api-auth/docs"
	"api-auth/internal/app"
	"api-auth/pkg/config/env"
	"api-auth/pkg/logger"
	config "api-auth/pkg/platform/bd"
	"api-auth/pkg/platform/redis"

	domain "api-auth/internal/domain/user"

	reqAut "api-auth/internal/handler/auth/dto/request"
	userHandler "api-auth/internal/handler/user"
	request "api-auth/internal/handler/user/dto/request"
	userRespServDto "api-auth/internal/service/auth/dto/response"

	"go.uber.org/zap"
)

var _ = request.CreateUserRequest{}
var _ = userHandler.NewUserHandler
var _ = userRespServDto.UserServiceResponseDto{}
var _ = domain.User{}
var _ = reqAut.LoginRequestDto{}

// @title API Auth Service
// @version 1.0
// @description Servicio de autenticación y gestión de usuarios.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /
// @schemes http

// main inicializa el servicio principal del API, configurando el logger, las
// variables de entorno, la base de datos y levantando el servidor HTTP.
//
// Parámetros:
//   - No recibe parámetros.
//
// Retorna:
//   - No retorna valores.
//
// Errores:
//   - Finaliza la ejecución con `logger.Log.Fatal` si ocurre un error crítico
//     al conectar con la base de datos o iniciar el servidor HTTP.
func main() {
	// Inicializar logger
	logger.Init()
	defer logger.Log.Sync()

	// Cargar configuración desde variables de entorno
	appConfig := env.Load()

	// Conectar a la base de datos
	if err := config.ConnectDB(); err != nil {
		logger.Log.Fatal("Error conectando a la base de datos", zap.Error(err))
	}

	logger.Log.Info("Conexión a la base de datos establecida")

	// Conectar a Redis
	if err := redis.ConnectRedis(); err != nil {
		logger.Log.Fatal("Error conectando a Redis", zap.Error(err))
	}
	logger.Log.Info("Conexión a Redis establecida")

	// Crear la instancia principal (inyectando logger)
	application := app.NewApp(logger.Log, appConfig)

	// Loguear puerto configurado
	logger.Log.Info("Servidor escuchando", zap.String("port", appConfig.AppPort))

	// Ejecutar servidor con puerto configurado
	if err := application.Run(appConfig.AppPort); err != nil {
		logger.Log.Fatal("Error al iniciar el servidor", zap.Error(err))
	}
}
