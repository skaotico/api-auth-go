// ============================================================
// @file: userMapper.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: Mapeador para convertir entidades de usuario a DTOs de respuesta.
// ============================================================

// Package mapper contiene las funciones de mapeo entre entidades y DTOs.
package mapper

import (
	domain "api-auth/internal/domain/user"
	resp "api-auth/internal/service/auth/dto/response"
)

// MapUserToResponse convierte una entidad User y un token en un UserServiceResponseDto.
//
// Parámetros:
//
//	u: Puntero a la entidad User.
//	token: Token de autenticación generado.
//
// Retorna:
//
//	*resp.UserServiceResponseDto: DTO de respuesta con los datos del usuario y el token.
func MapUserToResponse(u *domain.User, token string) *resp.UserServiceResponseDto {
	return &resp.UserServiceResponseDto{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Phone:     u.Phone,
		CountryID: u.CountryID,
		Address:   u.AddressLine,
		Token:     token,
	}
}
