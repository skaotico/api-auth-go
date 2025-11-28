// ============================================================
// @file: jwt.go
// @author: Yosemar Andrade
// @date: 2025-11-26
// @lastModified: 2025-11-26
// @description: Paquete platform proporciona utilidades para la gestión de JWT.
// ============================================================

package platform

import (
	"api-auth/pkg/logger"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

// JWTManager gestiona la creación y validación de tokens JWT.
type JWTManager struct {
	secretKey     []byte
	tokenDuration time.Duration
}

// NewJWTManager crea una nueva instancia de JWTManager.
//
// Parámetros:
//   - secret: clave secreta para firmar los tokens.
//   - duration: duración de validez del token.
//
// Retorna:
//   - *JWTManager: nueva instancia configurada.
//
// Errores:
//   - No retorna errores.
func NewJWTManager(secret string, duration time.Duration) *JWTManager {
	return &JWTManager{
		secretKey:     []byte(secret),
		tokenDuration: duration,
	}
}

// Generate crea un nuevo token JWT para un usuario.
//
// Parámetros:
//   - userID: ID del usuario.
//   - email: correo electrónico del usuario.
//
// Retorna:
//   - string: token JWT firmado.
//   - error: error si falla la firma del token.
//
// Errores:
//   - Retorna error si `token.SignedString` falla.
func (j *JWTManager) Generate(userID int, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(j.tokenDuration).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(j.secretKey)
	if err != nil {
		logger.Log.Error("Error firmando token JWT", zap.Error(err))
		return "", err
	}
	return signedToken, nil
}

// Validate verifica la validez de un token JWT.
//
// Parámetros:
//   - tokenString: el token en formato string.
//
// Retorna:
//   - *jwt.Token: el token parseado y validado.
//   - error: error si el token es inválido o el método de firma es incorrecto.
//
// Errores:
//   - Retorna error si el token no puede ser parseado o la firma no coincide.
func (j *JWTManager) Validate(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return j.secretKey, nil
	})

	if err != nil {
		// No logueamos error aquí como Error level porque puede ser un token expirado o inválido (comportamiento esperado)
		// Pero podemos loguearlo como Debug
		logger.Log.Debug("Token inválido o error al parsear", zap.Error(err))
		return nil, err
	}

	return token, nil
}
