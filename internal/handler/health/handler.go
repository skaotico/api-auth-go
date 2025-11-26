package health

import (
	service "api-auth/internal/service/health"

	"github.com/gin-gonic/gin"
)

type HealtHandler struct {
	service service.HealthService
}

func NewHealthHandler(s service.HealthService) *HealtHandler {
	return &HealtHandler{service: s}
}

func (s *HealtHandler) HealthCheck(c *gin.Context) {

	c.JSON(200, s.service.HealthCheck())
}
