// ============================================================
// @file: envConfig.go
// @author: Yosemar Andrade
// @date: 2025-11-26
// @lastModified: 2025-11-26
// @description: Módulo responsable de cargar, validar y mapear
// variables de entorno en la estructura de configuración del
// proyecto, priorizando variables del sistema y aplicando reglas
// de obligatoriedad mediante envconfig.
// ============================================================

package env

import (
	envPrimitivos "api-auth/pkg/config/env/dto/config"
	"api-auth/pkg/logger"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

// Load carga la configuración leyendo las variables de entorno.
//
// Primero intenta cargar el archivo `.env` (para desarrollo). Si no
// existe, continúa usando variables configuradas en el sistema. Luego
// mapea y valida los parámetros en Config utilizando envconfig.
//
// Retorna:
//   - *Config: estructura completamente cargada y validada.
//
// Errores:
//   - Finaliza la ejecución utilizando logger.Log.Fatal si faltan variables
//     obligatorias o existe algún error crítico en el mapeo.
func Load() *envPrimitivos.Config {
	// Paso 1: Intentar cargar .env (solo para desarrollo/local).
	if err := godotenv.Load("../../.env"); err != nil {
		logger.Log.Warn("Advertencia: No se encontró el archivo .env, usando variables de entorno del sistema.")
	}

	var cfg envPrimitivos.Config

	// Paso 2: Mapear y validar variables de entorno a la estructura Config.
	err := envconfig.Process("", &cfg)
	if err != nil {
		logger.Log.Fatal("Error crítico al cargar la configuración", zap.Error(err))
	}

	// Loggear el entorno cargado para depuración.
	logger.Log.Info("Configuración cargada", zap.String("environment", cfg.Environment))

	return &cfg
}
