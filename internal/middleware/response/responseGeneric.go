package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ApiResponseGeneric define la estructura estándar de respuesta
type ApiResponseGeneric[T any] struct {
	Success   bool   `json:"success"`
	Data      T      `json:"data,omitempty"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Path      string `json:"path"`
	ErrorCode string `json:"error_code,omitempty"`
}

// ResponseMiddleware devuelve un middleware que envuelve la respuesta
func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		path := c.Request.URL.Path
		timestamp := time.Now().Format(time.RFC3339)

		// Si hay error
		if respErr, exists := c.Get("response_error"); exists {
			errMap := respErr.(map[string]interface{})
			c.JSON(errMap["httpCode"].(int), ApiResponseGeneric[any]{
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
			c.JSON(http.StatusOK, ApiResponseGeneric[any]{
				Success:   true,
				Data:      resp,
				Message:   "Operación exitosa",
				Path:      path,
				Timestamp: timestamp,
			})
		}
	}
}
