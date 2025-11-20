// ============================================================
// @file: postgres_repository.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: Implementación de PostgreSQL para el repositorio de autenticación.
// ============================================================

// Package auth contiene la implementación del repositorio de autenticación.
package auth

import (
	config "api-auth/pkg/platform/bd"
	"database/sql"
)

// postgresAuthRepository implementa la interfaz Repository para PostgreSQL.
type postgresAuthRepository struct {
	db *sql.DB
}

// NewRepository crea una nueva instancia del repositorio de autenticación PostgreSQL.
//
// Retorna:
//   - Repository: Instancia del repositorio de autenticación.
func NewRepository() Repository {
	return &postgresAuthRepository{
		db: config.DB,
	}
}
