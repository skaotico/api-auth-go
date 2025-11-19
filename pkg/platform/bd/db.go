package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

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

// CheckDB valida el estado de la conexión (para health check).
func CheckDB() error {
	if DB == nil {
		return fmt.Errorf("database no inicializada")
	}
	return DB.Ping()
}
