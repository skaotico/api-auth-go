// ============================================================
// @file: error.go
// @author: Yosemar Andrade
// @date: 2025-11-26
// @lastModified: 2025-11-26
// @description: Define los errores de dominio para el módulo de usuario.
// ============================================================

package user

import "errors"

var (
	// ErrInvalidEmail indica que el formato del email es inválido.
	ErrInvalidEmail = errors.New("email inválido")
	// ErrInvalidPassword indica que la contraseña no cumple los requisitos o es incorrecta.
	ErrInvalidPassword = errors.New("contraseña incorrecta")
	// ErrUserNotFound indica que el usuario no fue encontrado en el sistema.
	ErrUserNotFound = errors.New("usuario no encontrado")
)
