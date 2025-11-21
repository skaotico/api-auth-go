// ============================================================
// @file: user.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: Define la entidad User.
// ============================================================

// Package user define las entidades y reglas de negocio del dominio de usuario.
package user

import "time"

// User representa un usuario del sistema.
type User struct {
	ID           int64      `json:"id"`
	Username     string     `json:"username"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"-"` // No se expone en JSON
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	Phone        *string    `json:"phone,omitempty"`
	BirthDate    *time.Time `json:"birth_date,omitempty"`
	IsActive     bool       `json:"is_active"`

	CountryID   int32   `json:"country_id"`
	AddressLine *string `json:"address_line,omitempty"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
