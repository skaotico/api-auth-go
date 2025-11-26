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
	"context"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// AuthService representa el servicio de autenticación responsable de
// manejar login de usuarios y generación de tokens JWT.
type AuthService struct {
	repo         repo.AuthRepository
	usService    userService.UserServiceInterface
	jwtConfig    config.JWTConfig
	cacheService cacheService.CacheService
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
func NewAuthService(r repo.AuthRepository, us userService.UserServiceInterface, jwtConfig config.JWTConfig, cache cacheService.CacheService) *AuthService {
	return &AuthService{
		repo:         r,
		usService:    us,
		jwtConfig:    jwtConfig,
		cacheService: cache,
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
func (s *AuthService) Login(loginDto *loginServiceDto.LoginServiceDto) (*userRespServDto.UserServiceResponseDto, error) {

	log.Printf("Iniciando proceso de login para el correo: %s", loginDto.Email)

	// Buscar usuario
	userFind, err := s.usService.GetUserByEmail(loginDto.Email)
	if err != nil {
		log.Printf("Usuario no encontrado con el email: %s", loginDto.Email)
		return nil, domain.ErrUserNotFound
	}

	log.Printf("Usuario encontrado: ID=%d, Email=%s", userFind.ID, userFind.Email)

	// Validar contraseña
	if bcrypt.CompareHashAndPassword([]byte(userFind.PasswordHash), []byte(loginDto.Password)) != nil {
		log.Printf("Contraseña incorrecta para el usuario con email: %s", loginDto.Email)
		return nil, domain.ErrInvalidPassword
	}

	log.Printf("Generando token JWT para el usuario ID=%d", userFind.ID)

	//TODO: los claims no estan implementados, ya que se implenentara redis para el manejo de esto,
	// por ahora se deja como comentario, si mas adelante es necesario generar claims, este es un ejemplo de uso
	// Generar token JWT
	// claims := jwt.MapClaims{
	// 	"user_id": userFind.ID,
	// 	"email":   userFind.Email,
	// 	"exp":     time.Now().Add(time.Hour * 24).Unix(),
	// 	"iat":     time.Now().Unix(),
	// }

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token := jwt.New(jwt.SigningMethodHS256)
	// Firmar token

	signedToken, err := token.SignedString([]byte(s.jwtConfig.Secret))
	if err != nil {
		log.Printf("Error al firmar el token JWT: %v", err)
		return nil, err
	}

	//esto genera un contexto con un timeout de 3 segundos, para que redis pueda responder en un tiempo razonable
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Creao los datos del JWT, para su guardado en cache
	jwtData := auth.JwtData{
		UserId:    strconv.Itoa(userFind.ID),
		Username:  userFind.Email,
		CreatedAt: time.Now().Unix(),
	}

	// Creo los datos del Refresh Token, para su guardado en cache, falta agregar la ip y el user agent
	refreshData := auth.RefreshData{
		UserId: strconv.Itoa(userFind.ID),

		CreatedAt: time.Now().Unix(),
	}

	// Guardar token en redis
	err = s.cacheService.SaveTokens(ctx, signedToken, signedToken, &jwtData, &refreshData, 0, 0)

	if err != nil {
		log.Printf("Error al guardar el token JWT en Redis: %v", err)
		return nil, err
	}
	log.Printf("Login exitoso para el usuario %s. Token generado correctamente.", userFind.Email)

	return mapper.MapUserToResponse(userFind, signedToken), nil
}
