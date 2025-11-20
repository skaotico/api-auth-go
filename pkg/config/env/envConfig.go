// ============================================================
// @file: envConfig.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: Módulo responsable de cargar, validar y mapear variables de entorno.
// ============================================================

// Package env gestiona la carga de variables de entorno.
package env

import (
	envPrimitivos "api-auth/pkg/config/env/dto/config"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
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
//   - Finaliza la ejecución utilizando log.Fatalf si faltan variables
//     obligatorias o existe algún error crítico en el mapeo.
func Load() *envPrimitivos.Config {
	// Paso 1: Intentar cargar .env (solo para desarrollo/local).
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("Advertencia: No se encontró el archivo .env, usando variables de entorno del sistema.")
	}

	var cfg envPrimitivos.Config

	// Paso 2: Mapear y validar variables de entorno a la estructura Config.
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("Error crítico al cargar la configuración: %v", err)
	}

	// Loggear el entorno cargado para depuración.
	log.Printf("Configuración cargada. Entorno: %s", cfg.Environment)

	return &cfg
}
