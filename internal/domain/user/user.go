// ============================================================
// @file: user.go
// @author: Yosemar Andrade
// @date: 2025-11-26
// @lastModified: 2025-11-26
// @description: Define la entidad User y sus propiedades.
// ============================================================

package user

import "time"

// User representa la entidad de usuario en el sistema.
type User struct {
	ID           int        `json:"id"`
	Username     string     `json:"username"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"-"` // No se expone en JSON
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	Phone        *string    `json:"phone,omitempty"`
	BirthDate    *time.Time `json:"birth_date,omitempty"`
	IsActive     bool       `json:"is_active"`

	CountryID   int     `json:"country_id"`
	AddressLine *string `json:"address_line,omitempty"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
