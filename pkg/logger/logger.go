// ============================================================
// @file: logger.go
// @author: Yosemar Andrade
// @date: 2025-11-26
// @lastModified: 2025-11-26
// @description: Paquete logger proporciona una instancia global de zap.Logger
// para el registro de eventos en la aplicación.
// ============================================================

package logger

import (
	"go.uber.org/zap"
)

// Log es la instancia global del logger.
var Log *zap.Logger

// Init inicializa el logger global utilizando la configuración de producción de zap.
//
// Parámetros:
//   - No recibe parámetros.
//
// Retorna:
//   - No retorna valores.
//
// Errores:
//   - Lanza un panic si ocurre un error al inicializar el logger.
func Init() {
	var err error

	Log, err = zap.NewProduction()
	if err != nil {
		panic("error al inicializar el logger de zap: " + err.Error())
	}
}
