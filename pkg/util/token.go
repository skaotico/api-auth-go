// ============================================================
// @file: token.go
// @author: Yosemar Andrade
// @date: 2025-11-26
// @lastModified: 2025-11-26
// @description: Paquete utils proporciona utilidades generales para la aplicación,
// incluyendo la generación de tokens seguros.
// ============================================================

package utils

import (
	"api-auth/pkg/logger"
	"crypto/rand"
	"encoding/base64"

	"go.uber.org/zap"
)

// NewRefreshToken genera un refresh token seguro, aleatorio y opaco.
//
// Parámetros:
//   - No recibe parámetros.
//
// Retorna:
//   - string: el token generado en formato base64 URL.
//   - error: error si falla la generación de bytes aleatorios.
//
// Errores:
//   - Retorna error si `rand.Read` falla.
func NewRandomID() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		logger.Log.Error("Error generando refresh token", zap.Error(err))
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
