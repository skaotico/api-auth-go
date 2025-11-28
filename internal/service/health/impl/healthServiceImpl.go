package impl

import (
	"api-auth/internal/service/health"
	"api-auth/internal/service/health/dto/config"
	"api-auth/internal/service/health/dto/response"
	db "api-auth/pkg/platform/bd"
	"time"
)

type HealthServiceImpl struct {
	healthConfig config.HealthConfig
}

func NewHealthService(healthConfig config.HealthConfig) health.HealthService {
	return &HealthServiceImpl{
		healthConfig: healthConfig,
	}
}

func (hs *HealthServiceImpl) HealthCheck() response.HealthResponse {

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
