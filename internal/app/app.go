// ============================================================
// @file: app.go
// @author: Yosemar Andrade
// @date: 2025-11-18
// @lastModified: 2025-11-26
// @description: Configuración principal de la aplicación, inyección de dependencias y rutas.
// ============================================================

package app

import (
	authHandler "api-auth/internal/handler/auth"
	userHandler "api-auth/internal/handler/user"
	"api-auth/internal/middleware/logging"
	"api-auth/internal/middleware/response"
	middleware "api-auth/internal/middleware/security"
	authRepository "api-auth/internal/repository/auth"
	userRepository "api-auth/internal/repository/user"
	jwtConfig "api-auth/internal/service/auth/dto/config"
	authService "api-auth/internal/service/auth/impl"
	"api-auth/internal/service/cache"
	cacheImpl "api-auth/internal/service/cache/impl"
	healthService "api-auth/internal/service/health"
	healthConfig "api-auth/internal/service/health/dto/config"
	healthServiceImpl "api-auth/internal/service/health/impl"
	userService "api-auth/internal/service/user/impl"
	envPrimitivos "api-auth/pkg/config/env/dto/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// App representa la aplicación principal
type App struct {
	Router *gin.Engine
	log    *zap.Logger
}

// NewApp inicializa la app con router, middlewares y dependencias
//
// Parámetros:
//   - logger: instancia de zap.Logger.
//   - configEnv: configuración cargada desde variables de entorno.
//
// Retorna:
//   - *App: instancia de la aplicación.
func NewApp(logger *zap.Logger, configEnv *envPrimitivos.Config) *App {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(logging.GinZap(logger))
	router.Use(response.ResponseMiddleware())

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// -------------------------------
	// Inyección de dependencias
	// -------------------------------

	// USER
	repoUser := userRepository.NewUserRepository()
	serviceUser := userService.NewUserService(repoUser, logger)
	handlerUser := userHandler.NewUserHandler(serviceUser)

	// AUTH
	authRepo := authRepository.NewAuthRepository()

	envJwtConfig := jwtConfig.JWTConfig{
		Secret:     configEnv.JWTSecret,
		Expiration: configEnv.JWTExpiration,
		RefreshTTL: configEnv.JWTRefreshTTL,
	}

	cacheService := cacheImpl.NewCacheService(logger)

	serviceAuth := authService.NewAuthService(authRepo, serviceUser, envJwtConfig, cacheService, logger)
	handlerAuth := authHandler.NewAuthHandler(serviceAuth)

	// HEALTH
	envHealthConfig := healthConfig.HealthConfig{
		Status:      "UP",
		Version:     configEnv.Version,
		Environment: configEnv.Environment,
		ServiceName: "api-auth",
	}
	serviceHealth := healthServiceImpl.NewHealthService(envHealthConfig)

	// -------------------------------
	// Setup de rutas
	// -------------------------------
	setupV1Routes(router, handlerUser, handlerAuth, serviceHealth, cacheService)

	return &App{
		Router: router,
		log:    logger,
	}
}

// setupV1Routes registra todas las rutas de la versión 1
func setupV1Routes(router *gin.Engine, userHandler *userHandler.UserHandler, authHandler *authHandler.AuthHandler, healthService healthService.HealthService, cacheService cache.CacheService) {
	v1 := router.Group("/v1")
	{
		// Health Check
		v1.GET("/health", func(c *gin.Context) {
			c.Header("Cache-Control", "no-store")
			resp := healthService.HealthCheck()
			c.JSON(200, resp)
		})

		// Users
		v1.GET("/users", userHandler.GetUsers)
		v1.POST("/users", userHandler.CreateUser)

		// Auth
		v1.POST("/auth/login", middleware.RateLimitLogin(cacheService), authHandler.Login)
		v1.POST("/auth/refresh", authHandler.RefreshToken)
	}
}

// Run inicia el servidor en el puerto especificado
//
// Parámetros:
//   - port: puerto en el que escuchará el servidor (ej. ":8080").
//
// Retorna:
//   - error: si falla el inicio del servidor.
func (a *App) Run(port string) error {
	return a.Router.Run(port)
}
