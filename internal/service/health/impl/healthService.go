// ============================================================
// @file: healthService.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: Implementación del servicio de estado y salud del sistema (Health Service).
// ============================================================

// Package impl implementa la lógica de negocio del servicio de salud.
package impl

import (
	"api-auth/internal/service/health/dto/config"
	"api-auth/internal/service/health/dto/response"
	db "api-auth/pkg/platform/bd"
	"time"
)

// HealthService representa el servicio encargado de obtener el estado de salud
// del sistema, incluyendo información de la base de datos, versión y entorno.
type HealthService struct {
	healthConfig config.HealthConfig
}

// NewHealthService crea una nueva instancia de HealthService.
//
// Parámetros:
//   - healthConfig: Configuración base del servicio de salud, que contiene
//     información como versión, estado, entorno y nombre del servicio.
//
// Retorna:
//   - *HealthService: Un puntero a la instancia creada del servicio.
func NewHealthService(healthConfig config.HealthConfig) *HealthService {
	return &HealthService{
		healthConfig: healthConfig,
	}
}

// HealthCheck obtiene el estado actual del sistema.
//
// La información retornada incluye:
// - Estado del servicio configurado.
// - Versión del proyecto.
// - Entorno de ejecución.
// - Nombre del servicio.
// - Estado de la base de datos.
// - Fecha y hora actual del servidor.
//
// Retorna:
//   - response.HealthResponse: Estructura con el estado general del sistema.
func (hs *HealthService) HealthCheck() response.HealthResponse {

	dbStatus := "CONNECTED"
	if err := db.CheckDB(); err != nil {
		dbStatus = "DISCONNECTED"
	}

	loc, _ := time.LoadLocation("America/Santiago")
	currentTime := time.Now().In(loc)

	return response.HealthResponse{
		Status:      hs.healthConfig.Status,
		Version:     hs.healthConfig.Version,
		Environment: hs.healthConfig.Environment,
		ServiceName: hs.healthConfig.ServiceName,
		DBStatus:    dbStatus,
		ServerTime:  currentTime,
	}
}
