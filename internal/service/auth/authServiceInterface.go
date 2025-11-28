package auth

import (
	loginServiceDto "api-auth/internal/service/auth/dto"
	userRespServDto "api-auth/internal/service/auth/dto/response"
)

// AuthServiceInterface define los métodos que debe implementar el servicio de autenticación.
type AuthServiceInterface interface {
	// Login realiza el proceso de autenticación de un usuario.
	//
	// Parámetros:
	//   - loginDto: DTO con email y contraseña.
	//
	// Retorna:
	//   - *UserServiceResponseDto: datos del usuario + token JWT.
	//   - string: refresh token.
	//   - error: si la autenticación falla o ocurre un error interno.
	Login(loginDto *loginServiceDto.LoginServiceDto) (*userRespServDto.UserServiceResponseDto, string, error)

	// RefreshToken renueva el access token y el refresh token.
	//
	// Parámetros:
	//   - refreshToken: el token de refresco actual.
	//
	// Retorna:
	//   - *UserServiceResponseDto: datos del usuario + nuevo token JWT.
	//   - string: nuevo refresh token.
	//   - error: si el token es inválido o ha expirado.
	RefreshToken(refreshToken string) (*userRespServDto.UserServiceResponseDto, string, error)
}
