// ============================================================
// @file: jwt.go
// @author: Yosemar Andrade
// @created: 2025-11-20
// @description: Módulo encargado de la generación y validación de tokens JWT.
// ============================================================

// Package platform contiene utilidades de plataforma como JWT.
package platform

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTManager representa el componente responsable de manejar la creación y
// validación de tokens JWT, utilizando una clave secreta y un tiempo de duración.
type JWTManager struct {
	secretKey     []byte
	tokenDuration time.Duration
}

// NewJWTManager crea una nueva instancia de JWTManager.
//
// Parámetros:
//   - secret: Clave secreta utilizada para firmar los tokens JWT.
//   - duration: Tiempo de expiración que será aplicado a cada token generado.
//
// Retorna:
//   - *JWTManager: Instancia configurada del manejador de JWT.
func NewJWTManager(secret string, duration time.Duration) *JWTManager {
	return &JWTManager{
		secretKey:     []byte(secret),
		tokenDuration: duration,
	}
}

// Generate crea un nuevo token JWT firmado utilizando la clave del JWTManager.
//
// Parámetros:
//   - userID: Identificador único del usuario al que pertenecen las credenciales.
//   - email: Correo electrónico asociado al usuario.
//
// Retorna:
//   - string: Token JWT generado.
//   - error: Error si ocurre un problema al firmar el token.
//
// Errores:
//   - Error en caso de que falle la firma del token JWT.
func (j *JWTManager) Generate(userID int, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(j.tokenDuration).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(j.secretKey)
}

// Validate verifica y decodifica un token JWT firmado.
//
// Parámetros:
//   - tokenString: Token JWT en formato string a validar.
//
// Retorna:
//   - *jwt.Token: Objeto token parseado con sus claims si es válido.
//   - error: Error si el token es inválido, expirado o se encuentra mal firmado.
//
// Errores:
//   - Si el método de firmado no coincide.
//   - Si el token es incorrecto o no puede ser parseado.
func (j *JWTManager) Validate(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return j.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
