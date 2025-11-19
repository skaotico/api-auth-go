package auth

import (
	loginServiceDto "api-auth/internal/service/auth/dto"
	userRespServDto "api-auth/internal/service/auth/dto/response"
)

// AuthServiceInterface define los métodos que debe implementar el servicio de autenticación.
type AuthServiceInterface interface {
	// Login autentica a un usuario mediante email y contraseña.
	//
	// Parámetros:
	//   - loginDto: DTO con email y contraseña.
	//
	// Retorna:
	//   - *UserServiceResponseDto: datos del usuario + token JWT.
	//   - error: si la autenticación falla o ocurre un error interno.
	Login(loginDto *loginServiceDto.LoginServiceDto) (*userRespServDto.UserServiceResponseDto, error)
}
