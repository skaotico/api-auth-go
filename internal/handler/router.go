package handler

import (
	"github.com/gin-gonic/gin"
)

// Router representa el enrutador de la aplicación.
type Router struct {
	Engine *gin.Engine
}

// NewRouter crea una nueva instancia de Router y configura el motor de Gin.
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
func (r *Router) Run(addr string) error {
	return r.Engine.Run(addr)
}
