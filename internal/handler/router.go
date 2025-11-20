// ============================================================
// @file: router.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: Configuración de rutas y servidor HTTP.
// ============================================================

// Package handler define la estructura de enrutamiento y configuración del servidor.
package handler

import (
	"github.com/gin-gonic/gin"
)

// Router representa el enrutador de la aplicación.
type Router struct {
	Engine *gin.Engine
}

// NewRouter crea una nueva instancia de Router y configura el motor de Gin.
//
// Retorna:
//
//	*Router: Instancia de Router inicializada.
func NewRouter() *Router {
	r := gin.Default()
	return &Router{
		Engine: r,
	}
}

// SetupRoutes configura todas las rutas de la API.
func (r *Router) SetupRoutes() {
	// Grupo de rutas para la versión 1 de la API
	v1 := r.Engine.Group("/v1")
	{
		// Ejemplo de una ruta de salud (Health Check)
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "UP", "service": "api-auth"})
		})

	}
}

// Run inicia el servidor HTTP en el puerto especificado.
//
// Parámetros:
//
//	addr: Dirección y puerto en el que escuchará el servidor.
//
// Retorna:
//
//	error: Error si falla el inicio del servidor.
func (r *Router) Run(addr string) error {
	return r.Engine.Run(addr)
}
