// ============================================================
// @file: handler.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: Handler encargado de exponer el estado de salud del servicio.
// ============================================================

// Package health contiene los handlers para el chequeo de salud.
package health

import (
	healthService "api-auth/internal/service/health"

	"github.com/gin-gonic/gin"
)

// HealtHandler representa el controlador encargado de gestionar las
// solicitudes relacionadas con el estado de salud del servicio.
type HealtHandler struct {
	service healthService.ServiceInterface
}

// NewHealthHandler crea una nueva instancia del handler de salud.
//
// Parámetros:
//   - s: Implementación del servicio de HealthServiceInterface.
//
// Retorna:
//   - *HealtHandler: Handler de salud listo para ser utilizado por Gin.
func NewHealthHandler(s healthService.ServiceInterface) *HealtHandler {
	return &HealtHandler{service: s}
}

// HealthCheck obtiene el estado actual del servicio y lo retorna como respuesta HTTP.
//
// @Summary Verificar el estado del servicio
// @Description Retorna información de estado general del sistema, base de datos, entorno y versión.
// @Tags Health
// @Produce json
// @Success 200 {object} response.HealthResponse
// @Router /health [get]
func (s *HealtHandler) HealthCheck(c *gin.Context) {
	c.JSON(200, s.service.HealthCheck())
}
