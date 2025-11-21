// ============================================================
// @file: authServiceInterface.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: Interfaz del servicio de autenticación.
// ============================================================

// Package auth define las interfaces y DTOs del servicio de autenticación.
package auth

import (
	loginServiceDto "api-auth/internal/service/auth/dto"
	userRespServDto "api-auth/internal/service/auth/dto/response"
)

// ServiceInterface define los métodos que debe implementar el servicio de autenticación.
type ServiceInterface interface {
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
