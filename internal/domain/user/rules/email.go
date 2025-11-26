// ============================================================
// @file: email.go
// @author: Yosemar Andrade
// @date: 2025-11-26
// @lastModified: 2025-11-26
// @description: Define reglas de validación para el email.
// ============================================================

package rules

import (
	"api-auth/internal/domain/user"
	"strings"
)

// ValidateEmail verifica si el email tiene un formato válido.
//
// Parámetros:
//   - email: la dirección de correo a validar.
//
// Retorna:
//   - error: retorna error si el email no contiene '@'.
//
// Errores:
//   - Retorna `user.ErrInvalidEmail` si la validación falla.
func ValidateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return user.ErrInvalidEmail
	}
	return nil
}
