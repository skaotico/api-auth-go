// ============================================================
// @file: loginRequestDto.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: DTO para la solicitud de inicio de sesión.
// ============================================================

// Package request contiene los DTOs de solicitud para el handler de autenticación.
package request

// LoginRequestDto representa los datos necesarios para iniciar sesión.
type LoginRequestDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
