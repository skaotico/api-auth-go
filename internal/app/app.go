// ============================================================
// @file: app.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: Define la estructura y configuración principal de la aplicación.
// ============================================================

// Package app provee la configuración y arranque de la aplicación.
package app

import (
	authHandler "api-auth/internal/handler/auth"
	"api-auth/internal/middleware/logging"
	"api-auth/internal/middleware/response"
	authRepository "api-auth/internal/repository/auth"
	userRepository "api-auth/internal/repository/user"
	jwtConfig "api-auth/internal/service/auth/dto/config"
	authService "api-auth/internal/service/auth/impl"
	userService "api-auth/internal/service/user/impl"
	envPrimitivos "api-auth/pkg/config/env/dto/config"
	platform "api-auth/pkg/platform/bd"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// App representa la aplicación principal.
type App struct {
	Router *gin.Engine
	log    *zap.Logger
}

// NewApp inicializa la app con router, middlewares y dependencias.
//
// Parámetros:
//
//	logger: Instancia de zap.Logger para logging.
//	configEnv: Configuración de variables de entorno.
//
// Retorna:
//
//	*App: Puntero a la instancia de la aplicación inicializada.
func NewApp(logger *zap.Logger, configEnv *envPrimitivos.Config) *App {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(logging.GinZap(logger))
	router.Use(response.MiddlewareResponse())

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// -------------------------------
	// Inyección de dependencias
	// -------------------------------

	// USER
	repoUser := userRepository.NewRepository()
	serviceUser := userService.NewUserService(repoUser)
	// handlerUser := userHandler.NewUserHandler(serviceUser)

	// AUTH
	authRepo := authRepository.NewRepository()

	envJwtConfig := jwtConfig.JWTConfig{
		Secret:     configEnv.JWTSecret,
		Expiration: configEnv.JWTExpiration,
		RefreshTTL: configEnv.JWTRefreshTTL,
	}

	serviceAuth := authService.NewAuthService(authRepo, serviceUser, envJwtConfig)
	handlerAuth := authHandler.NewAuthHandler(serviceAuth)

	// -------------------------------
	// Setup de rutas
	// -------------------------------
	setupV1Routes(router, handlerAuth, configEnv.Version, configEnv.Environment)

	return &App{
		Router: router,
		log:    logger,
	}
}

// setupV1Routes registra todas las rutas de la versión 1.
//
// Parámetros:
//
//	router: Motor de Gin.
//	authHandler: Handler de autenticación.
//	version: Versión de la aplicación.
//	environment: Entorno de ejecución.
func setupV1Routes(router *gin.Engine, authHandler *authHandler.Handler, version string, environment string) {
	v1 := router.Group("/v1")
	{
		dbStatus := "OK"
		if err := platform.CheckDB(); err != nil {
			dbStatus = "DOWN"
		}

		v1.GET("/health", func(c *gin.Context) {
			c.Header("Cache-Control", "no-store")
			c.JSON(200, gin.H{
				"status":    "UP",
				"service":   "api-auth",
				"env":       environment,
				"version":   version,
				"time":      time.Now().Format(time.RFC3339),
				"status_db": dbStatus,
			})
		})

		// Users
		// v1.GET("/users", userHandler.GetUsers)
		// v1.POST("/users", userHandler.CreateUser)

		// Auth
		v1.POST("/auth/login", authHandler.Login)
	}
}

// Run inicia el servidor en el puerto especificado.
//
// Parámetros:
//
//	port: Puerto en el que escuchará el servidor (ej: ":8080").
//
// Retorna:
//
//	error: Error si falla el inicio del servidor.
func (a *App) Run(port string) error {
	return a.Router.Run(port)
}
