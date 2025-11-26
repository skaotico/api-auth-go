// ============================================================
// @file: auth_service.go
// @author: Yosemar Andrade
// @date: 2025-11-19
// @lastModified: 2025-11-19
// @description: Implementa el servicio de autenticación con login y generación de JWT.
// ============================================================

package impl

import (
	"api-auth/internal/domain/auth"
	domain "api-auth/internal/domain/user"
	mapper "api-auth/internal/mapper/user"
	repo "api-auth/internal/repository/auth"
	loginServiceDto "api-auth/internal/service/auth/dto"
	"api-auth/internal/service/auth/dto/config"
	userRespServDto "api-auth/internal/service/auth/dto/response"
	cacheService "api-auth/internal/service/cache"
	userService "api-auth/internal/service/user"
	utils "api-auth/pkg/util"
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// AuthService representa el servicio de autenticación responsable de
// manejar login de usuarios y generación de tokens JWT.
type AuthService struct {
	repo         repo.AuthRepository
	usService    userService.UserService
	jwtConfig    config.JWTConfig
	cacheService cacheService.CacheService

	logger *zap.Logger
}

// NewAuthService crea una nueva instancia de AuthService.
//
// Parámetros:
//
//	r: repositorio de autenticación
//	us: servicio de usuario para obtener datos de usuarios
//	jwtConfig: configuración de JWT (clave secreta, expiración, etc.)
//
// Retorna:
//
//	*AuthService: puntero a la nueva instancia de AuthService
func NewAuthService(r repo.AuthRepository, us userService.UserService, jwtConfig config.JWTConfig, cache cacheService.CacheService, logger *zap.Logger) *AuthService {
	return &AuthService{
		repo:         r,
		usService:    us,
		jwtConfig:    jwtConfig,
		cacheService: cache,
		logger:       logger.With(zap.String("service", "AuthService")),
	}
}

// Login realiza el proceso de autenticación de un usuario.
//
// Parámetros:
//
//	loginDto: DTO con el correo y contraseña del usuario
//
// Retorna:
//
//	*userRespServDto.UserServiceResponseDto: DTO con información del usuario y token JWT
//	error: error si el usuario no existe, la contraseña es inválida o ocurre un fallo en la generación del token
func (s *AuthService) Login(loginDto *loginServiceDto.LoginServiceDto) (*userRespServDto.UserServiceResponseDto, string, error) {

	s.logger.Info("Iniciando login", zap.String("email", loginDto.Email))

	// Buscar usuario
	userFind, err := s.usService.GetUserByEmail(loginDto.Email)
	if err != nil {
		s.logger.Warn("Usuario no encontrado", zap.String("email", loginDto.Email), zap.Error(err))
		return nil, "", domain.ErrUserNotFound
	}

	s.logger.Debug("Usuario encontrado",
		zap.Int("userId", userFind.ID),
		zap.String("email", userFind.Email),
	)

	// Validar contraseña
	if bcrypt.CompareHashAndPassword([]byte(userFind.PasswordHash), []byte(loginDto.Password)) != nil {
		s.logger.Warn("Contraseña incorrecta", zap.String("email", loginDto.Email))
		return nil, "", domain.ErrInvalidPassword
	}

	s.logger.Debug("Generando token JWT", zap.Int("userId", userFind.ID))

	jti, err := utils.NewRandomID()
	if err != nil {
		return nil, "", err
	}

	claims := jwt.MapClaims{
		"jti": jti,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(s.jwtConfig.Expiration).Unix(),
		"typ": "access",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(s.jwtConfig.Secret))
	if err != nil {
		s.logger.Error("Error firmando token JWT", zap.Error(err))
		return nil, "", err
	}

	// Contexto para Redis
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Datos para cache
	jwtData := auth.JwtData{
		TokenID:   jti,
		UserId:    strconv.Itoa(userFind.ID),
		Username:  userFind.Email,
		CreatedAt: time.Now().Unix(),
	}

	refreshData := auth.RefreshData{
		UserId:    strconv.Itoa(userFind.ID),
		CreatedAt: time.Now().Unix(),
	}

	// Generar Refresh token
	refreshToken, err := utils.NewRandomID()
	if err != nil {
		s.logger.Error("Error generando refresh token", zap.Error(err))
		return nil, "", err
	}

	// Guardar en Redis
	if err := s.cacheService.SaveTokens(ctx, signedToken, refreshToken, &jwtData, &refreshData, s.jwtConfig.Expiration, s.jwtConfig.RefreshTTL); err != nil {
		s.logger.Error("Error guardando tokens en Redis", zap.Error(err))
		return nil, "", err
	}

	s.logger.Info("Login exitoso",
		zap.Int("userId", userFind.ID),
		zap.String("email", userFind.Email),
	)

	//TODO: descomentar para verificar datos guardados en Redis llamando a los métodos del cache service, @Skaotico

	// Verificar datos guardados en Redis llamando a los métodos del cache service
	// savedJwtData, err := s.cacheService.GetJwtData(ctx, signedToken)
	// if err != nil {
	// 	s.logger.Error("Error recuperando JWT desde Redis para verificación", zap.Error(err))
	// } else {
	// 	s.logger.Info("JWT recuperado desde Redis",
	// 		zap.String("userId", savedJwtData.UserId),
	// 		zap.String("username", savedJwtData.Username),
	// 		zap.Int64("createdAt", savedJwtData.CreatedAt),
	// 	)
	// }

	// savedRefreshData, err := s.cacheService.GetRefreshData(ctx, refreshToken)
	// if err != nil {
	// 	s.logger.Error("Error recuperando Refresh desde Redis para verificación", zap.Error(err))
	// } else {
	// 	s.logger.Info("Refresh recuperado desde Redis",
	// 		zap.String("userId", savedRefreshData.UserId),
	// 		zap.Int64("createdAt", savedRefreshData.CreatedAt),
	// 	)
	// }

	// savedUserIndex, err := s.cacheService.GetUserIndex(ctx, jwtData.UserId)
	// if err != nil {
	// 	s.logger.Error("Error recuperando UserIndex desde Redis para verificación", zap.Error(err))
	// } else {
	// 	s.logger.Debug("UserIndex recuperado desde Redis",
	// 		zap.String("userId", jwtData.UserId),
	// 		zap.String("activeJwt", savedUserIndex.ActiveJwt[:20]+"..."),
	// 		zap.String("activeRefresh", savedUserIndex.ActiveRefresh[:20]+"..."),
	// 		zap.Int64("lastLogin", savedUserIndex.LastLogin),
	// 	)
	// }

	return mapper.MapUserToResponse(userFind, signedToken), refreshToken, nil
}

// RefreshToken renueva el access token y el refresh token.
//
// Parámetros:
//   - refreshToken: el token de refresco actual.
//
// Retorna:
//   - *UserServiceResponseDto: datos del usuario + nuevo token JWT.
//   - string: nuevo refresh token.
//   - error: si el token es inválido o ha expirado.
func (s *AuthService) RefreshToken(refreshToken string) (*userRespServDto.UserServiceResponseDto, string, error) {
	s.logger.Info("Iniciando refresh token")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// 1. Validar si el refresh token existe en Redis
	refreshData, err := s.cacheService.GetRefreshData(ctx, refreshToken)
	if err != nil {
		s.logger.Error("Refresh token inválido o expirado", zap.Error(err))
		return nil, "", errors.New("refresh token inválido o expirado")
	}

	// 2. Validar si el usuario existe
	userIdInt, err := strconv.Atoi(refreshData.UserId)
	if err != nil {
		s.logger.Error("Error convirtiendo userId a int", zap.Error(err))
		return nil, "", err
	}

	userFind, err := s.usService.GetUserByID(userIdInt)
	if err != nil {
		s.logger.Warn("Usuario asociado al token no encontrado", zap.String("userId", refreshData.UserId))
		return nil, "", domain.ErrUserNotFound
	}

	// 3. Validar reutilización de token (Token Rotation Check)
	userIndex, err := s.cacheService.GetUserIndex(ctx, refreshData.UserId)
	if err == nil {
		if userIndex.ActiveRefresh != refreshToken {
			s.logger.Warn("Detectado posible reuso de refresh token", zap.String("userId", refreshData.UserId))
			// Opcional: Invalidar todo
			// s.cacheService.DeleteAll(ctx, refreshData.UserId, userIndex.ActiveJwt, userIndex.ActiveRefresh)
			return nil, "", errors.New("token de refresco inválido")
		}
	}

	// 4. Generar nuevos tokens
	jti, err := utils.NewRandomID()
	if err != nil {
		return nil, "", err
	}

	claims := jwt.MapClaims{
		"jti": jti,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(s.jwtConfig.Expiration).Unix(),
		"typ": "access",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.jwtConfig.Secret))
	if err != nil {
		s.logger.Error("Error firmando token JWT", zap.Error(err))
		return nil, "", err
	}

	newRefreshToken, err := utils.NewRandomID()
	if err != nil {
		return nil, "", err
	}

	// 5. Guardar nuevos tokens
	newJwtData := auth.JwtData{
		TokenID:   jti,
		UserId:    refreshData.UserId,
		Username:  userFind.Email,
		CreatedAt: time.Now().Unix(),
	}
	newRefreshData := auth.RefreshData{
		UserId:    refreshData.UserId,
		CreatedAt: time.Now().Unix(),
	}

	if err := s.cacheService.SaveTokens(ctx, signedToken, newRefreshToken, &newJwtData, &newRefreshData, s.jwtConfig.Expiration, s.jwtConfig.RefreshTTL); err != nil {
		s.logger.Error("Error guardando nuevos tokens en Redis", zap.Error(err))
		return nil, "", err
	}

	s.logger.Info("Refresh token exitoso", zap.String("userId", refreshData.UserId))

	return mapper.MapUserToResponse(userFind, signedToken), newRefreshToken, nil
}
