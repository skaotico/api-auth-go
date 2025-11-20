// ============================================================
// @file: password.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: Reglas de validación para contraseñas.
// ============================================================

// Package rules contiene las reglas de validación de dominio.
package rules

import (
	"api-auth/internal/domain/user"
	"strings"
)

// ValidatePasswordNotEmpty valida que la contraseña no esté vacía.
//
// Parámetros:
//
//	pass: Contraseña a validar.
//
// Retorna:
//
//	error: Retorna user.ErrInvalidPassword si la contraseña está vacía, nil en caso contrario.
func ValidatePasswordNotEmpty(pass string) error {
	if strings.TrimSpace(pass) == "" {
		return user.ErrInvalidPassword
	}
	return nil
}
