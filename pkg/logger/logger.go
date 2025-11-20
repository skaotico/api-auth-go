// ============================================================
// @file: logger.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: Inicialización del logger central de la aplicación utilizando Uber Zap.
// ============================================================

// Package logger provee funcionalidades de logging.
package logger

import (
	"go.uber.org/zap"
)

// Log representa la instancia global del logger disponible para toda la aplicación.
// Se inicializa mediante la función Init().
var Log *zap.Logger

// Init configura e inicializa el logger de Uber Zap en modo producción.
// En caso de fallo durante la inicialización, se fuerza un panic, dado que la
// aplicación requiere el logger para registrar eventos críticos.
//
// Errores:
//   - Si ocurre un error al crear la instancia del logger, se dispara un panic.
func Init() {
	var err error

	Log, err = zap.NewProduction()
	if err != nil {
		panic("error al inicializar el logger de zap: " + err.Error())
	}
}
