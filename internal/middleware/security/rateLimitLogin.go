package middleware

import (
	"api-auth/internal/domain/security"
	"api-auth/internal/service/cache"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimitLogin middleware que limita a 3 intentos de login por IP en 1 minuto.
func RateLimitLogin(cacheService cache.CacheService) gin.HandlerFunc {
	return func(c *gin.Context) {

		ip := c.ClientIP()
		key := "rate_limit:login:ip:" + ip
		now := time.Now().Unix()

		// =========================================================
		// Obtener el registro desde cache
		// =========================================================
		data, _ := cacheService.GetRateLimit(c, key) // si falla, lo manejamos igual

		// =========================================================
		// Si no existe o expiró, reiniciamos
		// =========================================================
		if data == nil || data.ExpiresAt < now {
			data = &security.RateLimitData{
				Key:       key,
				Limit:     3,        // Máximo 3 intentos
				Attempts:  1,        // Primer intento
				ExpiresAt: now + 60, // Expira en 60 segundos
			}
		} else {
			// =========================================================
			// Incrementamos intentos si todavía está vigente
			// =========================================================
			data.Attempts++
		}

		// =========================================================
		// Guardar en redis (aunque falle, no afecta el flujo)
		// =========================================================
		_ = cacheService.SaveRateLimit(c, data)

		// =========================================================
		// Validación de límite
		// =========================================================
		if data.Attempts > data.Limit {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error":       "Has excedido el límite de intentos. Intenta de nuevo más tarde.",
				"retry_after": data.ExpiresAt - now, // tiempo restante
			})
			c.Abort()
			return
		}

		// =========================================================
		// Continuar si estamos bajo el límite
		// =========================================================
		c.Next()
	}
}
