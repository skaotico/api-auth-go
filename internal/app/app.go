package app

import (
	authHandler "api-auth/internal/handler/auth"
	userHandler "api-auth/internal/handler/user"
	"api-auth/internal/middleware/logging"
	"api-auth/internal/middleware/response"
	authRepository "api-auth/internal/repository/auth"
	userRepository "api-auth/internal/repository/user"
	jwtConfig "api-auth/internal/service/auth/dto/config"
	authService "api-auth/internal/service/auth/impl"
	cache "api-auth/internal/service/cache/impl"
	userService "api-auth/internal/service/user/impl"
	envPrimitivos "api-auth/pkg/config/env/dto/config"
	platform "api-auth/pkg/platform/bd"
	"time"

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
	serviceUser := userService.NewUserService(repoUser)
	// handlerUser := userHandler.NewUserHandler(serviceUser)

	// AUTH
	authRepo := authRepository.NewAuthRepository()

	envJwtConfig := jwtConfig.JWTConfig{
		Secret:     configEnv.JWTSecret,
		Expiration: configEnv.JWTExpiration,
		RefreshTTL: configEnv.JWTRefreshTTL,
	}

	cacheService := cache.NewCacheService()

	serviceAuth := authService.NewAuthService(authRepo, serviceUser, envJwtConfig, cacheService)
	handlerAuth := authHandler.NewAuthHandler(serviceAuth)

	// -------------------------------
	// Setup de rutas
	// -------------------------------
	setupV1Routes(router, nil, handlerAuth, configEnv.Version, configEnv.Environment)

	return &App{
		Router: router,
		log:    logger,
	}
}

// setupV1Routes registra todas las rutas de la versión 1
func setupV1Routes(router *gin.Engine, userHandler *userHandler.UserHandler, authHandler *authHandler.AuthHandler, version string, environment string) {
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

// Run inicia el servidor en el puerto especificado
func (a *App) Run(port string) error {
	return a.Router.Run(port)
}
