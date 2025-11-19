// ============================================================
// @file: health_config.go
// @author: Yosemar Andrade
// @date: 2025-11-19
// @lastModified: 2025-11-19
// @description: Define la configuración de estado de salud del servicio.
// ============================================================

package config

// HealthConfig representa la información de estado de salud de un servicio.
//
// Incluye el estado general, versión, entorno, estado de la base de datos
// y el nombre del servicio.
type HealthConfig struct {
	// Status indica el estado general del servicio (ej. "UP", "DOWN").
	Status string

	// Version indica la versión actual del servicio.
	Version string

	// Environment indica el entorno en el que corre el servicio (ej. "dev", "prod").
	Environment string

	// ServiceName indica el nombre del servicio.
	ServiceName string
}
