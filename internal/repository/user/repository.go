// ============================================================
// @file: repository.go
// @author: Yosemar Andrade
// @date: 2025-11-26
// @lastModified: 2025-11-26
// @description: Define la interfaz del repositorio de usuarios.
// ============================================================

package user

import (
	domain "api-auth/internal/domain/user"
)

// UserRepository define los métodos para el repositorio de usuarios.
type UserRepository interface {
	// FindByEmail busca un usuario por su correo electrónico.
	//
	// Parámetros:
	//   - email: correo electrónico del usuario.
	//
	// Retorna:
	//   - *domain.User: el usuario encontrado.
	//   - error: error si no se encuentra o hay fallo en BD.
	FindByEmail(email string) (*domain.User, error)
	FindByID(id int) (*domain.User, error)

	// FindAll lista todos los usuarios.
	//
	// Parámetros:
	//   - No recibe parámetros.
	//
	// Retorna:
	//   - []*domain.User: lista de usuarios.
	//   - error: error si falla la consulta.
	FindAll() ([]*domain.User, error)

	// Save guarda un nuevo usuario en la base de datos.
	//
	// Parámetros:
	//   - user: puntero al usuario a guardar.
	//
	// Retorna:
	//   - error: error si falla la inserción.
	Save(user *domain.User) error
}
