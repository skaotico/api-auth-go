// ============================================================
// @file: db.go
// @author: Yosemar Andrade
// @date: 2025-11-26
// @lastModified: 2025-11-26
// @description: Paquete db maneja la conexión a la base de datos PostgreSQL.
// ============================================================

package db

import (
	"api-auth/pkg/logger"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

// DB es la instancia global de la conexión a la base de datos.
var DB *sql.DB

// ConnectDB establece la conexión con la base de datos PostgreSQL.
//
// Parámetros:
//   - No recibe parámetros.
//
// Retorna:
//   - error: retorna error si falla la apertura o el ping a la base de datos.
//
// Errores:
//   - Retorna error si `sql.Open` o `db.Ping` fallan.
func ConnectDB() error {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		"5432",
		"skaotico",
		"U2thb3RpY28yMDI1Lg==",
		"eco-registro",
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Log.Error("Error abriendo conexión a BD", zap.Error(err))
		return err
	}

	if err := db.Ping(); err != nil {
		logger.Log.Error("Error haciendo ping a BD", zap.Error(err))
		return err
	}

	DB = db
	logger.Log.Info("Conectado a PostgreSQL")
	return nil
}

// CheckDB valida el estado de la conexión (para health check).
func CheckDB() error {
	if DB == nil {
		return fmt.Errorf("database no inicializada")
	}
	return DB.Ping()
}
