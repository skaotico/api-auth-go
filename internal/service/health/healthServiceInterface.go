// ============================================================
// @file: healthServiceInterface.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: Define la interfaz para el servicio de health check.
// ============================================================

// Package health define las interfaces y DTOs del servicio de salud.
package health

import "api-auth/internal/service/health/dto/response"

// ServiceInterface representa la interfaz que define los métodos
// que un servicio de health check debe implementar.
type ServiceInterface interface {
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
