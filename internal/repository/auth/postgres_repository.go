// ============================================================
// @file: postgres_repository.go
// @author: Yosemar Andrade
// @date: 2025-11-26
// @lastModified: 2025-11-26
// @description: Implementación del repositorio de autenticación para PostgreSQL.
// ============================================================

package auth

import (
	config "api-auth/pkg/platform/bd"
	"database/sql"
)

type postgresAuthRepository struct {
	db *sql.DB
}

// NewAuthRepository crea una nueva instancia del repositorio de autenticación.
//
// Parámetros:
//   - No recibe parámetros.
//
// Retorna:
//   - AuthRepository: interfaz del repositorio de autenticación.
//
// Errores:
//   - No retorna errores.
func NewAuthRepository() AuthRepository {
	return &postgresAuthRepository{
		db: config.DB,
	}
}
