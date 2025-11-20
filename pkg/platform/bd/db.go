// ============================================================
// @file: db.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: Módulo de conexión y validación de estado de la base de datos PostgreSQL.
// ============================================================

// Package db gestiona la conexión a la base de datos.
package db

import (
	"database/sql"
	"fmt"
)

// DB representa la instancia global de conexión activa hacia PostgreSQL.
var DB *sql.DB

// ConnectDB establece la conexión con la base de datos PostgreSQL utilizando
// las credenciales configuradas localmente. La conexión es almacenada en la
// variable global DB.
//
// Retorna:
//   - error: Error si ocurre un problema al abrir o validar la conexión con la base de datos.
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
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	DB = db
	fmt.Println("Conectado a PostgreSQL")
	return nil
}

// CheckDB valida el estado de la conexión activa hacia la base de datos PostgreSQL.
// Esta función es utilizada principalmente para el servicio de Health Check.
//
// Retorna:
//   - error: Error si la conexión no ha sido inicializada o si el Ping a la base de datos falla.
func CheckDB() error {
	if DB == nil {
		return fmt.Errorf("database no inicializada")
	}
	return DB.Ping()
}
