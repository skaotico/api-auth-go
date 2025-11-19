package user

import "errors"

var (
	ErrInvalidEmail    = errors.New("email inválido")
	ErrInvalidPassword = errors.New("contraseña incorrecta")
	ErrUserNotFound    = errors.New("usuario no encontrado")
)
