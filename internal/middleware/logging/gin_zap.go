// ============================================================
// @file: gin_zap.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: Middleware de logging HTTP para Gin utilizando Uber Zap.
// ============================================================

// Package logging contiene los middlewares de logging.
package logging

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GinZap crea un middleware para Gin que registra información de cada
// solicitud HTTP utilizando el logger Zap proporcionado.
//
// El registro incluye:
//   - Método HTTP utilizado.
//   - Ruta solicitada.
//   - Código de estado HTTP.
//   - Tiempo de latencia.
//   - IP del cliente.
//   - User-Agent.
//
// Parámetros:
//   - logger: Instancia de zap.Logger utilizada para registrar los eventos.
//
// Retorna:
//   - gin.HandlerFunc: Middleware que puede ser agregado al router de Gin.
func GinZap(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		inicio := time.Now()
		c.Next()
		duracion := time.Since(inicio)

		logger.Info("Solicitud HTTP",
			zap.String("metodo", c.Request.Method),
			zap.String("ruta", c.Request.URL.Path),
			zap.Int("estado", c.Writer.Status()),
			zap.Duration("latencia", duracion),
			zap.String("ip_cliente", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
		)
	}
}
