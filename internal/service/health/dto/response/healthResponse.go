// ============================================================
// @file: health_response.go
// @author: Yosemar Andrade
// @date: 2025-11-19
// @lastModified: 2025-11-19
// @description: Define la estructura de respuesta para el health check.
// ============================================================

package response

import "time"

// HealthResponse representa la respuesta del health check del servicio.
//
// Incluye los datos de configuración inicial, el estado de la base de datos
// y la hora del servidor en Chile.
type HealthResponse struct {
	// Status indica el estado general del servicio (ej. "UP", "DOWN").
	Status string `json:"status"`

	// Version indica la versión actual del servicio.
	Version string `json:"version"`

	// Environment indica el entorno en el que corre el servicio (ej. "dev", "prod").
	Environment string `json:"environment"`

	// ServiceName indica el nombre del servicio.
	ServiceName string `json:"service_name"`

	// DBStatus indica el estado de la base de datos asociada al servicio.
	DBStatus string `json:"db_status"`

	// ServerTime indica la hora del servidor en Chile.
	ServerTime time.Time `json:"server_time"`
}
