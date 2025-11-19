package logging

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

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
