package health

import (
	service "api-auth/internal/service/health"

	"github.com/gin-gonic/gin"
)

type HealtHandler struct {
	service service.HealthServiceInterface
}

func NewHealthHandler(s service.HealthServiceInterface) *HealtHandler {
	return &HealtHandler{service: s}
}

func (s *HealtHandler) HealthCheck(c *gin.Context) {

	c.JSON(200, s.service.HealthCheck())
}
