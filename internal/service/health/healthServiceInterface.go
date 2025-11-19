// ============================================================
// @file: health_service_interface.go
// @author: Yosemar Andrade
// @date: 2025-11-19
// @lastModified: 2025-11-19
// @description: Define la interfaz para el servicio de health check.
// ============================================================

package health

import "api-auth/internal/service/health/dto/response"

// HealthServiceInterface representa la interfaz que define los métodos
// que un servicio de health check debe implementar.
type HealthServiceInterface interface {
	// HealthCheck obtiene el estado de salud del servicio.
	//
	// Parámetros:
	//  Ninguno
	//
	// Retorna:
	//  config.HealthConfig: estructura que contiene información de estado, versión, entorno y nombre del servicio.
	//  error: error en caso de que no se pueda obtener el estado.
	HealthCheck() response.HealthResponse
}
