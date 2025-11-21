// ============================================================
// @file: authService.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: Implementa el servicio de autenticación con login y generación de JWT.
// ============================================================

// Package impl implementa la lógica de negocio del servicio de autenticación.
package impl

import (
	domain "api-auth/internal/domain/user"
	mapper "api-auth/internal/mapper/user"
	authRepo "api-auth/internal/repository/auth"
	loginServiceDto "api-auth/internal/service/auth/dto"
	"api-auth/internal/service/auth/dto/config"
	userRespServDto "api-auth/internal/service/auth/dto/response"
	userService "api-auth/internal/service/user"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// AuthService representa el servicio de autenticación responsable de
// manejar login de usuarios y generación de tokens JWT.
type AuthService struct {
	repo      authRepo.Repository
	usService userService.ServiceInterface
	jwtConfig config.JWTConfig
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
func NewAuthService(r authRepo.Repository, us userService.ServiceInterface, jwtConfig config.JWTConfig) *AuthService {
	return &AuthService{
		repo:      r,
		usService: us,
		jwtConfig: jwtConfig,
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

	claims := jwt.MapClaims{
		"user_id": userFind.ID,
		"email":   userFind.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar token

	signedToken, err := token.SignedString([]byte(s.jwtConfig.Secret))
	if err != nil {
		log.Printf("Error al firmar el token JWT: %v", err)
		return nil, err
	}

	log.Printf("Login exitoso para el usuario %s. Token generado correctamente.", userFind.Email)

	return mapper.MapUserToResponse(userFind, signedToken), nil
}
