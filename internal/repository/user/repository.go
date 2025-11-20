// ============================================================
// @file: repository.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: Interfaz para el repositorio de usuarios.
// ============================================================

// Package user contiene la implementación del repositorio de usuarios.
package user

import (
	domain "api-auth/internal/domain/user"
)

// Repository define las operaciones de persistencia para usuarios.
type Repository interface {
	// FindByEmail busca un usuario por su correo electrónico.
	FindByEmail(email string) (*domain.User, error)
	// FindAll lista todos los usuarios.
	FindAll() ([]*domain.User, error)
	// Save guarda un nuevo usuario.
	Save(user *domain.User) error
}
