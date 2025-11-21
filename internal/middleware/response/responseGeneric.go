// ============================================================
// @file: responseGeneric.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: Middleware para estandarizar las respuestas de la API.
// ============================================================

// Package response contiene los middlewares y estructuras para la gestión de respuestas.
package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// APIResponseGeneric define la estructura estándar de respuesta.
type APIResponseGeneric[T any] struct {
	Success   bool   `json:"success"`
	Data      T      `json:"data,omitempty"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Path      string `json:"path"`
	ErrorCode string `json:"error_code,omitempty"`
}

// MiddlewareResponse devuelve un middleware que envuelve la respuesta.
//
// Retorna:
//
//	gin.HandlerFunc: Middleware que estandariza las respuestas.
func MiddlewareResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		path := c.Request.URL.Path
		timestamp := time.Now().Format(time.RFC3339)

		// Si hay error
		if respErr, exists := c.Get("response_error"); exists {
			errMap := respErr.(map[string]interface{})
			c.JSON(errMap["httpCode"].(int), APIResponseGeneric[any]{
				Success:   false,
				Message:   errMap["message"].(string),
				Path:      path,
				Timestamp: timestamp,
				ErrorCode: errMap["errorCode"].(string),
			})
			return
		}

		// Respuesta exitosa
		if resp, exists := c.Get("response"); exists {
			c.JSON(http.StatusOK, APIResponseGeneric[any]{
				Success:   true,
				Data:      resp,
				Message:   "Operación exitosa",
				Path:      path,
				Timestamp: timestamp,
			})
		}
	}
}
