// ============================================================
// @file: error.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: Definición de errores específicos asociados a la lógica de usuarios.
// ============================================================

// Package user define las entidades y reglas de negocio del dominio de usuario.
package user

import "errors"

// ErrInvalidEmail representa un error cuando el correo ingresado no cumple
// con un formato válido o no es aceptado por el sistema.
var ErrInvalidEmail = errors.New("email inválido")

// ErrInvalidPassword indica que la contraseña proporcionada no coincide con
// la registrada para el usuario o no cumple con las validaciones necesarias.
var ErrInvalidPassword = errors.New("contraseña incorrecta")

// ErrUserNotFound se utiliza cuando el usuario solicitado no existe en el
// sistema o no puede ser recuperado desde la capa de persistencia.
var ErrUserNotFound = errors.New("usuario no encontrado")
