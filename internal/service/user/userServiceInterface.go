// ============================================================
// @file: ServiceInterface.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: Interfaz del servicio de usuarios.
// ============================================================

// Package user define las interfaces del servicio de usuarios.
package user

import domain "api-auth/internal/domain/user"

// ServiceInterface define los métodos que debe implementar el servicio de usuarios.
type ServiceInterface interface {
	// GetAllUsers obtiene todos los usuarios registrados.
	GetAllUsers() ([]*domain.User, error)
	// GetUserByEmail obtiene un usuario según su email.
	GetUserByEmail(email string) (*domain.User, error)
	// Login valida las credenciales de un usuario.
	Login(email, password string) (*domain.User, error)
	// CreateUser crea un nuevo usuario.
	CreateUser(u *domain.User, plainPassword string) error
}
