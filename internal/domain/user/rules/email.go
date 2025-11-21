// ============================================================
// @file: email.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: Reglas de validación para correos electrónicos.
// ============================================================

// Package rules contiene las reglas de validación de dominio.
package rules

import (
	"api-auth/internal/domain/user"
	"strings"
)

// ValidateEmail valida si un correo electrónico tiene un formato correcto.
//
// Parámetros:
//
//	email: Correo electrónico a validar.
//
// Retorna:
//
//	error: Retorna user.ErrInvalidEmail si el correo no es válido, nil en caso contrario.
func ValidateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return user.ErrInvalidEmail
	}
	return nil
}
