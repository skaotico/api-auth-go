package impl

import (
	"api-auth/internal/service/health/dto/config"
	"api-auth/internal/service/health/dto/response"
	db "api-auth/pkg/platform/bd"
	"time"
)

type HealthService struct {
	healthConfig config.HealthConfig
}

func NewHealthService(healthConfig config.HealthConfig) *HealthService {
	return &HealthService{
		healthConfig: healthConfig,
	}
}

func (hs *HealthService) HealthCheck() response.HealthResponse {

	dbStatus := "CONNECTED"
	if err := db.CheckDB(); err != nil {
		dbStatus = "DISCONNECTED"
	}

	loc, _ := time.LoadLocation("America/Santiago")
	currentTime := time.Now().In(loc)
	return response.HealthResponse{
		Status:      hs.healthConfig.Status,
		Version:     hs.healthConfig.Version,
		Environment: hs.healthConfig.Environment,
		ServiceName: hs.healthConfig.ServiceName,
		DBStatus:    dbStatus,
		ServerTime:  currentTime,
	}
}
